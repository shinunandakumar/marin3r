package generators

import (
	"fmt"
	"time"

	"github.com/3scale-ops/marin3r/pkg/envoy"
	corev1 "k8s.io/api/core/v1"
)

type GeneratorOptions struct {
	InstanceName              string
	Namespace                 string
	EnvoyAPIVersion           envoy.APIVersion
	EnvoyNodeID               string
	EnvoyClusterID            string
	ClientCertificateDuration time.Duration
	AdminBindAddress          string
	DeploymentImage           string
	DeploymentResources       corev1.ResourceRequirements
}

func (cfg *GeneratorOptions) labels() map[string]string {
	return map[string]string{
		"app.kubernetes.io/name":       "marin3r",
		"app.kubernetes.io/managed-by": "marin3r-operator",
		"app.kubernetes.io/component":  "envoy-deployment",
		"app.kubernetes.io/instance":   cfg.InstanceName,
	}
}

func (cfg *GeneratorOptions) resourceName() string {
	return fmt.Sprintf("%s-%s", "marin3r-envoy-deployment", cfg.InstanceName)
}
