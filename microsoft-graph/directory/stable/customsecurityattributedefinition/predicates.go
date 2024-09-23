package customsecurityattributedefinition

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import "github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/stable"

type CustomSecurityAttributeDefinitionOperationPredicate struct {
}

func (p CustomSecurityAttributeDefinitionOperationPredicate) Matches(input stable.CustomSecurityAttributeDefinition) bool {

	return true
}
