package org.infinispan.crd;

import cz.xtf.core.openshift.crd.CustomResourceDefinitionContextProvider;
import io.fabric8.kubernetes.client.dsl.base.CustomResourceDefinitionContext;

public class GrafanaDataSourceContextProvider implements CustomResourceDefinitionContextProvider {
   private CustomResourceDefinitionContext grafanaContext;

   public CustomResourceDefinitionContext getContext() {
      if (grafanaContext == null) {
         grafanaContext = new CustomResourceDefinitionContext.Builder()
               .withGroup("integreatly.org")
               .withPlural("grafanadatasources")
               .withScope("Namespaced")
               .withVersion("v1alpha1")
               .build();
      }

      return grafanaContext;
   }
}
