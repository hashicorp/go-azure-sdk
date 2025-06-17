package teamprimarychannelplannerplantask

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import "github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"

type PlannerTaskOperationPredicate struct {
}

func (p PlannerTaskOperationPredicate) Matches(input beta.PlannerTask) bool {

	return true
}
