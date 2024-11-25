package common

import (
	"fmt"

	xpref "github.com/crossplane/crossplane-runtime/pkg/reference"
	xpresource "github.com/crossplane/crossplane-runtime/pkg/resource"
	"github.com/crossplane/upjet/pkg/resource"
)

const (
	SelfPackagePath            = "github.com/crossplane-contrib/provider-confluent/config/common"
	ExtractPrincipalIDFuncPath = SelfPackagePath + ".ExtractPrincipalID()"
)

// Extract the identifier of a user to use for
// principal references
func ExtractPrincipalID() xpref.ExtractValueFn {
	return func(mr xpresource.Managed) string {
		tr, ok := mr.(resource.Terraformed)
		if !ok {
			return ""
		}

		userID := tr.GetID()
		if len(userID) == 0 {
			return ""
		}

		return fmt.Sprintf("User:%v", userID) // append 'User:' infront of the service account ID when resolving 'principal' field.
	}
}
