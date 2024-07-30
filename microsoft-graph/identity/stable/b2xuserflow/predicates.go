package b2xuserflow

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import "github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/stable"

type B2xIdentityUserFlowOperationPredicate struct {
}

func (p B2xIdentityUserFlowOperationPredicate) Matches(input stable.B2xIdentityUserFlow) bool {

	return true
}
