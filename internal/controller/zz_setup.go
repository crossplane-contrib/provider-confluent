// SPDX-FileCopyrightText: 2023 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"

	apikey "github.com/crossplane-contrib/provider-confluent/internal/controller/confluent/apikey"
	cluster "github.com/crossplane-contrib/provider-confluent/internal/controller/confluent/cluster"
	clusterconfig "github.com/crossplane-contrib/provider-confluent/internal/controller/confluent/clusterconfig"
	environment "github.com/crossplane-contrib/provider-confluent/internal/controller/confluent/environment"
	kafkaacl "github.com/crossplane-contrib/provider-confluent/internal/controller/confluent/kafkaacl"
	kafkatopic "github.com/crossplane-contrib/provider-confluent/internal/controller/confluent/kafkatopic"
	rolebinding "github.com/crossplane-contrib/provider-confluent/internal/controller/confluent/rolebinding"
	schemaregistrycluster "github.com/crossplane-contrib/provider-confluent/internal/controller/confluent/schemaregistrycluster"
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
		kafkatopic.Setup,
		rolebinding.Setup,
		schemaregistrycluster.Setup,
		serviceaccount.Setup,
		providerconfig.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
