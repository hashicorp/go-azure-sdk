package authenticationfido2method

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import "github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"

type Fido2AuthenticationMethodOperationPredicate struct {
}

func (p Fido2AuthenticationMethodOperationPredicate) Matches(input beta.Fido2AuthenticationMethod) bool {

	return true
}
