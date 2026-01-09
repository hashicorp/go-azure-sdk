package provisioning

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import "github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"

type ProvisioningObjectSummaryOperationPredicate struct {
}

func (p ProvisioningObjectSummaryOperationPredicate) Matches(input beta.ProvisioningObjectSummary) bool {

	return true
}
