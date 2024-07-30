package privilegedaccesgroupassignmentapproval

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import "github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/stable"

type ApprovalOperationPredicate struct {
}

func (p ApprovalOperationPredicate) Matches(input stable.Approval) bool {

	return true
}
