package virtualendpointfrontlineserviceplan

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import "github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"

type CloudPCFrontLineServicePlanOperationPredicate struct {
}

func (p CloudPCFrontLineServicePlanOperationPredicate) Matches(input beta.CloudPCFrontLineServicePlan) bool {

	return true
}
