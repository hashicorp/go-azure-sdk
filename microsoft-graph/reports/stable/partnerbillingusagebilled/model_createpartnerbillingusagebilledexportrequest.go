package partnerbillingusagebilled

import (
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/stable"
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CreatePartnerBillingUsageBilledExportRequest struct {
	AttributeSet *stable.PartnersBillingAttributeSet `json:"attributeSet,omitempty"`
	InvoiceId    nullable.Type[string]               `json:"invoiceId,omitempty"`
}
