package lifecycleworkflowdeleteditemworkflowrun

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import "github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/stable"

type IdentityGovernanceRunOperationPredicate struct {
}

func (p IdentityGovernanceRunOperationPredicate) Matches(input stable.IdentityGovernanceRun) bool {

	return true
}
