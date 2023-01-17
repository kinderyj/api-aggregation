package animal

import (
	"code/github.com/kinderyj/api-agg/pkg/apis/animal/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

var AddToScheme1 = func(scheme *runtime.Scheme) error {
	metav1.AddToGroupVersion(scheme, schema.GroupVersion{
		Group:   "animal.kinderyj.io",
		Version: "v1alpha1",
	})
	// +kubebuilder:scaffold:install

	scheme.AddKnownTypes(schema.GroupVersion{
		Group:   "animal.kinderyj.io",
		Version: "v1alpha1",
	}, &v1alpha1.Cat{}, &v1alpha1.CatList{})
	return nil
}
