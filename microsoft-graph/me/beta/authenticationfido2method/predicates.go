package authenticationfido2method

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import "github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"

type Fido2AuthenticationMethodOperationPredicate struct {
}

func (p Fido2AuthenticationMethodOperationPredicate) Matches(input beta.Fido2AuthenticationMethod) bool {

	return true
}
