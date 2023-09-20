/*
Copyright 2021 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/upbound/upjet/pkg/controller"

	apikey "github.com/crossplane-contrib/provider-confluent/internal/controller/confluent/apikey"
	cluster "github.com/crossplane-contrib/provider-confluent/internal/controller/confluent/cluster"
	clusterconfig "github.com/crossplane-contrib/provider-confluent/internal/controller/confluent/clusterconfig"
	environment "github.com/crossplane-contrib/provider-confluent/internal/controller/confluent/environment"
	kafkaacl "github.com/crossplane-contrib/provider-confluent/internal/controller/confluent/kafkaacl"
	rolebinding "github.com/crossplane-contrib/provider-confluent/internal/controller/confluent/rolebinding"
	serviceaccount "github.com/crossplane-contrib/provider-confluent/internal/controller/confluent/serviceaccount"
	providerconfig "github.com/crossplane-contrib/provider-confluent/internal/controller/providerconfig"
)

// Setup creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		apikey.Setup,
		cluster.Setup,
		clusterconfig.Setup,
		environment.Setup,
		kafkaacl.Setup,
		rolebinding.Setup,
		serviceaccount.Setup,
		providerconfig.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
