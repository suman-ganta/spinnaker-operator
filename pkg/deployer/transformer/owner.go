package transformer

import (
	spinnakerv1alpha1 "github.com/armory/spinnaker-operator/pkg/apis/spinnaker/v1alpha1"
	"github.com/armory/spinnaker-operator/pkg/generated"
	"github.com/armory/spinnaker-operator/pkg/halconfig"
	"github.com/go-logr/logr"
	"k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/controller/controllerutil"
)

type ownerTransformer struct {
	*DefaultTransformer
	svc spinnakerv1alpha1.SpinnakerServiceInterface
	log logr.Logger
}

type ownerTransformerGenerator struct{}

func (g *ownerTransformerGenerator) NewTransformer(svc spinnakerv1alpha1.SpinnakerServiceInterface,
	hc *halconfig.SpinnakerConfig, client client.Client, log logr.Logger) (Transformer, error) {
	base := &DefaultTransformer{}
	tr := ownerTransformer{svc: svc, log: log, DefaultTransformer: base}
	base.ChildTransformer = &tr
	return &tr, nil
}

func (g *ownerTransformerGenerator) GetName() string {
	return "SetOwner"
}

// transform adjusts settings to the configuration
func (t *ownerTransformer) TransformManifests(scheme *runtime.Scheme, gen *generated.SpinnakerGeneratedConfig) error {
	// Set SpinnakerService instance as the owner and controller
	for k := range gen.Config {
		s := gen.Config[k]
		if s.Deployment != nil {
			if err := controllerutil.SetControllerReference(t.svc, s.Deployment, scheme); err != nil {
				return err
			}
		}
		if s.Service != nil {
			if err := controllerutil.SetControllerReference(t.svc, s.Service, scheme); err != nil {
				return err
			}
		}
		// Don't own the resources, they'll be owned by the Deployment
	}
	return nil
}