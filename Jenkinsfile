#!/usr/bin/env groovy

pipeline {
    agent {
        label 'slave-group-k8s'
    }

    environment {
        GO111MODULE = 'on'
        KIND_VERSION = 'v0.11.0'
        METALLB_VERSION = 'v0.9.6'
        KUBECONFIG = "$WORKSPACE/kind-kube-config.yaml"
        TESTING_NAMESPACE = 'namespace-for-testing'
        WATCH_NAMESPACE = 'namespace-for-testing'
        PATH="/opt/go/bin:$PATH"
        RUN_SA_OPERATOR = 'true'
        MAKE_DATADIR_WRITABLE = 'true'
    }

    options {
        timeout(time: 30, unit: 'MINUTES')
        timestamps()
        buildDiscarder(logRotator(numToKeepStr: '100', daysToKeepStr: '61'))
    }

    stages {
        stage('Build') {
            steps {
                // TODO Ignore doc PRS etc
                sh 'go mod vendor'
                sh 'make lint'
            }
        }

        stage('Unit') {
            steps {
                sh 'make unit-test'
            }
        }

        stage('Create k8s Cluster') {
            steps {
                sh 'kind create cluster --config kind-config.yaml'
                writeFile file: "${KUBECONFIG}", text: sh(script: 'kind get kubeconfig', , returnStdout: true)

                script {
                    env.TESTING_CONTEXT = sh(script: 'kubectl --insecure-skip-tls-verify config current-context', , returnStdout: true).trim()
                }

                sh 'build/travis/prepare-sa-test.sh'
            }
        }

        // stage('Core') {
        //     steps {
        //         sh "make test PARALLEL_COUNT=2"
        //     }
        // }

        stage('Batch') {
            steps {
                sh 'make batch-test PARALLEL_COUNT=2'
            }
        }

        stage('Multinamespace') {
            steps {
                sh "kubectl config use-context $TESTING_CONTEXT"
                sh 'make multinamespace-test'
            }
        }

        stage('Backup/Restore') {
            steps {
                catchError(buildResult: 'SUCCESS', stageResult: 'FAILURE') {
                    sh 'make backuprestore-test'
                }
            }
        }

        stage('Xsite') {
            steps {
                sh 'build/travis/configure-travis-ci-xsite.sh'
                sh 'go test -v ./test/e2e/xsite/ -timeout 30m'
            }
        }
    }

    post {
        failure {
            sh "kubectl config use-context $TESTING_CONTEXT"
            sh 'kubectl get events --all-namespaces'
            sh 'kubectl cluster-info'
            sh 'kubectl logs daemonset/speaker -n metallb-system --context kind-xsite1'
            sh 'kubectl logs deployment/controller -n metallb-system --context kind-xsite1'
            sh 'kubectl logs daemonset/speaker -n metallb-system --context kind-xsite2'
            sh 'kubectl logs deployment/controller -n metallb-system --context kind-xsite2'
            sh 'df -h'
            sh 'docker ps -a'
            sh 'for log in $(docker ps -qa | xargs); do docker logs --tail 500 $log; done'
        }

        cleanup {
            sh 'kind delete clusters --all'
            sh 'docker container prune -f'
            sh 'docker rmi $(docker images -f "dangling=true" -q) || true'
        }
    }
}
