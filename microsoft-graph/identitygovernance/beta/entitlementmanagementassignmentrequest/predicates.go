package entitlementmanagementassignmentrequest

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import "github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"

type AccessPackageAssignmentRequestOperationPredicate struct {
}

func (p AccessPackageAssignmentRequestOperationPredicate) Matches(input beta.AccessPackageAssignmentRequest) bool {

	return true
}
