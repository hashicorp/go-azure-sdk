package permissionsmanagementpermissionsrequestchange

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import "github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"

type PermissionsRequestChangeOperationPredicate struct {
}

func (p PermissionsRequestChangeOperationPredicate) Matches(input beta.PermissionsRequestChange) bool {

	return true
}
