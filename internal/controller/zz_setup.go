// SPDX-FileCopyrightText: 2023 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"

	apikey "github.com/stakater/provider-confluent/internal/controller/confluent/apikey"
	cluster "github.com/stakater/provider-confluent/internal/controller/confluent/cluster"
	clusterconfig "github.com/stakater/provider-confluent/internal/controller/confluent/clusterconfig"
	environment "github.com/stakater/provider-confluent/internal/controller/confluent/environment"
	identityprovider "github.com/stakater/provider-confluent/internal/controller/confluent/identityprovider"
	kafkaacl "github.com/stakater/provider-confluent/internal/controller/confluent/kafkaacl"
	rolebinding "github.com/stakater/provider-confluent/internal/controller/confluent/rolebinding"
	serviceaccount "github.com/stakater/provider-confluent/internal/controller/confluent/serviceaccount"
	providerconfig "github.com/stakater/provider-confluent/internal/controller/providerconfig"
)

// Setup creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		apikey.Setup,
		cluster.Setup,
		clusterconfig.Setup,
		environment.Setup,
		identityprovider.Setup,
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
