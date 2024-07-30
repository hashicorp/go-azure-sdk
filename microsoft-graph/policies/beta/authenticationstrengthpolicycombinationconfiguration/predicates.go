package authenticationstrengthpolicycombinationconfiguration

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import "github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"

type AuthenticationCombinationConfigurationOperationPredicate struct {
}

func (p AuthenticationCombinationConfigurationOperationPredicate) Matches(input beta.AuthenticationCombinationConfiguration) bool {

	return true
}
