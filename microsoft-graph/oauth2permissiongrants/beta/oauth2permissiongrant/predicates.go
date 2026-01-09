package oauth2permissiongrant

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import "github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"

type OAuth2PermissionGrantOperationPredicate struct {
}

func (p OAuth2PermissionGrantOperationPredicate) Matches(input beta.OAuth2PermissionGrant) bool {

	return true
}
