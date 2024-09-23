package teampermissiongrant

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import "github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/stable"

type ResourceSpecificPermissionGrantOperationPredicate struct {
}

func (p ResourceSpecificPermissionGrantOperationPredicate) Matches(input stable.ResourceSpecificPermissionGrant) bool {

	return true
}
