/*
Copyright 2021 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/upbound/upjet/pkg/controller"

	key "github.com/crossplane-contrib/provider-confluent/internal/controller/api/key"
	environment "github.com/crossplane-contrib/provider-confluent/internal/controller/confluent/environment"
	acl "github.com/crossplane-contrib/provider-confluent/internal/controller/kafka/acl"
	cluster "github.com/crossplane-contrib/provider-confluent/internal/controller/kafka/cluster"
	clusterconfig "github.com/crossplane-contrib/provider-confluent/internal/controller/kafka/clusterconfig"
	providerconfig "github.com/crossplane-contrib/provider-confluent/internal/controller/providerconfig"
	binding "github.com/crossplane-contrib/provider-confluent/internal/controller/role/binding"
	account "github.com/crossplane-contrib/provider-confluent/internal/controller/service/account"
)

// Setup creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		key.Setup,
		environment.Setup,
		acl.Setup,
		cluster.Setup,
		clusterconfig.Setup,
		providerconfig.Setup,
		binding.Setup,
		account.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
