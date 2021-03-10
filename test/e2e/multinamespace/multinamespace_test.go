package multinamespace

import (
	"os"
	"strings"
	"sync"
	"testing"

	ispnv1 "github.com/infinispan/infinispan-operator/pkg/apis/infinispan/v1"
	tutils "github.com/infinispan/infinispan-operator/test/e2e/utils"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	logf "sigs.k8s.io/controller-runtime/pkg/runtime/log"
)

var testKube = tutils.NewTestKubernetes(os.Getenv("TESTING_CONTEXT"))
var serviceAccountKube = tutils.NewTestKubernetes("")

var log = logf.Log.WithName("multinamespace_test")

var MinimalSpec = ispnv1.Infinispan{
	TypeMeta: tutils.InfinispanTypeMeta,
	ObjectMeta: metav1.ObjectMeta{
		Name: tutils.DefaultClusterName,
	},
	Spec: ispnv1.InfinispanSpec{
		Replicas: 1,
	},
}

func TestMain(m *testing.M) {
	nsAsString := strings.ToLower(tutils.MultiNamespace)
	if "TRUE" == tutils.RunLocalOperator {
		stopCh := testKube.InstallAndRunOperator(nsAsString, "../../../deploy/crds/", true)
		code := m.Run()
		close(stopCh)
		os.Exit(code)
	} else {
		code := m.Run()
		os.Exit(code)
	}
}

// Test if single node working correctly
func TestMultinamespaceNodeStartup(t *testing.T) {
	// Create a resource without passing any config
	nsAsString := strings.ToLower(tutils.MultiNamespace)
	namespaces := strings.Split(nsAsString, ",")
	var wg sync.WaitGroup
	for _, namespace := range namespaces {
		spec := MinimalSpec.DeepCopy()
		spec.Namespace = namespace
		// Register it
		testKube.CreateInfinispan(spec, namespace)
		defer testKube.DeleteInfinispan(spec, tutils.SinglePodTimeout)
		wg.Add(1)
		go func() {
			testKube.WaitForInfinispanPods(1, tutils.SinglePodTimeout, spec.Name, spec.Namespace)
			testKube.WaitForInfinispanCondition(spec.Name, spec.Namespace, ispnv1.ConditionWellFormed)
			wg.Done()
		}()
	}
	wg.Wait()
}
