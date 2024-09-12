package partnerbillingusageunbilled

import (
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CreatePartnerBillingUsageUnbilledExportRequest struct {
	AttributeSet  *beta.PartnersBillingAttributeSet  `json:"attributeSet,omitempty"`
	BillingPeriod *beta.PartnersBillingBillingPeriod `json:"billingPeriod,omitempty"`
	CurrencyCode  nullable.Type[string]              `json:"currencyCode,omitempty"`
}
