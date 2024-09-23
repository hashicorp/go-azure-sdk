package mailfolderuserconfiguration

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import "github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"

type UserConfigurationOperationPredicate struct {
}

func (p UserConfigurationOperationPredicate) Matches(input beta.UserConfiguration) bool {

	return true
}
