package b2xuserflowidentityprovider

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import "github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/stable"

type IdentityProviderOperationPredicate struct {
}

func (p IdentityProviderOperationPredicate) Matches(input stable.IdentityProvider) bool {

	return true
}
