package partnerbillingoperation

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import "github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/stable"

type PartnersBillingOperationOperationPredicate struct {
}

func (p PartnersBillingOperationOperationPredicate) Matches(input stable.PartnersBillingOperation) bool {

	return true
}
