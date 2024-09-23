package windowsqualityupdatepolicyassignment

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import "github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"

type WindowsQualityUpdatePolicyAssignmentOperationPredicate struct {
}

func (p WindowsQualityUpdatePolicyAssignmentOperationPredicate) Matches(input beta.WindowsQualityUpdatePolicyAssignment) bool {

	return true
}
