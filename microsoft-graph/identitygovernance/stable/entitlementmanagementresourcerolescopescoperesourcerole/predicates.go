package entitlementmanagementresourcerolescopescoperesourcerole

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import "github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/stable"

type AccessPackageResourceRoleOperationPredicate struct {
}

func (p AccessPackageResourceRoleOperationPredicate) Matches(input stable.AccessPackageResourceRole) bool {

	return true
}
