package authenticationstrengthpolicy

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import "github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"

type AuthenticationStrengthPolicyOperationPredicate struct {
}

func (p AuthenticationStrengthPolicyOperationPredicate) Matches(input beta.AuthenticationStrengthPolicy) bool {

	return true
}
