/*
Copyright 2022 Upbound Inc.
*/
// Code generated by angryjet. DO NOT EDIT.

package v1alpha1

import (
	"context"
	confluentkafkaacl "github.com/crossplane-contrib/provider-confluent/config/confluent_kafka_acl"
	reference "github.com/crossplane/crossplane-runtime/pkg/reference"
	errors "github.com/pkg/errors"
	client "sigs.k8s.io/controller-runtime/pkg/client"
)

// ResolveReferences of this APIKey.
func (mg *APIKey) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	for i3 := 0; i3 < len(mg.Spec.ForProvider.ManagedResource); i3++ {
		for i4 := 0; i4 < len(mg.Spec.ForProvider.ManagedResource[i3].Environment); i4++ {
			rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
				CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ManagedResource[i3].Environment[i4].ID),
				Extract:      reference.ExternalName(),
				Reference:    mg.Spec.ForProvider.ManagedResource[i3].Environment[i4].IDRef,
				Selector:     mg.Spec.ForProvider.ManagedResource[i3].Environment[i4].IDSelector,
				To: reference.To{
					List:    &EnvironmentList{},
					Managed: &Environment{},
				},
			})
			if err != nil {
				return errors.Wrap(err, "mg.Spec.ForProvider.ManagedResource[i3].Environment[i4].ID")
			}
			mg.Spec.ForProvider.ManagedResource[i3].Environment[i4].ID = reference.ToPtrValue(rsp.ResolvedValue)
			mg.Spec.ForProvider.ManagedResource[i3].Environment[i4].IDRef = rsp.ResolvedReference

		}
	}
	for i3 := 0; i3 < len(mg.Spec.ForProvider.ManagedResource); i3++ {
		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ManagedResource[i3].ID),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.ForProvider.ManagedResource[i3].IDRef,
			Selector:     mg.Spec.ForProvider.ManagedResource[i3].IDSelector,
			To: reference.To{
				List:    &ClusterList{},
				Managed: &Cluster{},
			},
		})
		if err != nil {
			return errors.Wrap(err, "mg.Spec.ForProvider.ManagedResource[i3].ID")
		}
		mg.Spec.ForProvider.ManagedResource[i3].ID = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.ForProvider.ManagedResource[i3].IDRef = rsp.ResolvedReference

	}
	for i3 := 0; i3 < len(mg.Spec.ForProvider.Owner); i3++ {
		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.Owner[i3].ID),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.ForProvider.Owner[i3].IDRef,
			Selector:     mg.Spec.ForProvider.Owner[i3].IDSelector,
			To: reference.To{
				List:    &ServiceAccountList{},
				Managed: &ServiceAccount{},
			},
		})
		if err != nil {
			return errors.Wrap(err, "mg.Spec.ForProvider.Owner[i3].ID")
		}
		mg.Spec.ForProvider.Owner[i3].ID = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.ForProvider.Owner[i3].IDRef = rsp.ResolvedReference

	}

	return nil
}

// ResolveReferences of this Cluster.
func (mg *Cluster) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	for i3 := 0; i3 < len(mg.Spec.ForProvider.Environment); i3++ {
		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.Environment[i3].ID),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.ForProvider.Environment[i3].IDRef,
			Selector:     mg.Spec.ForProvider.Environment[i3].IDSelector,
			To: reference.To{
				List:    &EnvironmentList{},
				Managed: &Environment{},
			},
		})
		if err != nil {
			return errors.Wrap(err, "mg.Spec.ForProvider.Environment[i3].ID")
		}
		mg.Spec.ForProvider.Environment[i3].ID = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.ForProvider.Environment[i3].IDRef = rsp.ResolvedReference

	}

	return nil
}

// ResolveReferences of this KafkaACL.
func (mg *KafkaACL) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	for i3 := 0; i3 < len(mg.Spec.ForProvider.KafkaCluster); i3++ {
		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.KafkaCluster[i3].ID),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.ForProvider.KafkaCluster[i3].IDRef,
			Selector:     mg.Spec.ForProvider.KafkaCluster[i3].IDSelector,
			To: reference.To{
				List:    &ClusterList{},
				Managed: &Cluster{},
			},
		})
		if err != nil {
			return errors.Wrap(err, "mg.Spec.ForProvider.KafkaCluster[i3].ID")
		}
		mg.Spec.ForProvider.KafkaCluster[i3].ID = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.ForProvider.KafkaCluster[i3].IDRef = rsp.ResolvedReference

	}
	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.Principal),
		Extract:      confluentkafkaacl.ExtractResourceID(),
		Reference:    mg.Spec.ForProvider.PrincipalRef,
		Selector:     mg.Spec.ForProvider.PrincipalSelector,
		To: reference.To{
			List:    &ServiceAccountList{},
			Managed: &ServiceAccount{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.Principal")
	}
	mg.Spec.ForProvider.Principal = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.PrincipalRef = rsp.ResolvedReference

	return nil
}
