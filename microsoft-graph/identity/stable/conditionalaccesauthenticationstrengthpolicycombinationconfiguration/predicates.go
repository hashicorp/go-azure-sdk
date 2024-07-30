package conditionalaccesauthenticationstrengthpolicycombinationconfiguration

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import "github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/stable"

type AuthenticationCombinationConfigurationOperationPredicate struct {
}

func (p AuthenticationCombinationConfigurationOperationPredicate) Matches(input stable.AuthenticationCombinationConfiguration) bool {

	return true
}
