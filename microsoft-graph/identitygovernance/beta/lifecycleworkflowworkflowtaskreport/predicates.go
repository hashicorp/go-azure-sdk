package lifecycleworkflowworkflowtaskreport

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import "github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"

type IdentityGovernanceTaskReportOperationPredicate struct {
}

func (p IdentityGovernanceTaskReportOperationPredicate) Matches(input beta.IdentityGovernanceTaskReport) bool {

	return true
}
