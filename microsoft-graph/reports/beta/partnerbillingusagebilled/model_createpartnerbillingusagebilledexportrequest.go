package partnerbillingusagebilled

import (
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CreatePartnerBillingUsageBilledExportRequest struct {
	AttributeSet *beta.PartnersBillingAttributeSet `json:"attributeSet,omitempty"`
	InvoiceId    nullable.Type[string]             `json:"invoiceId,omitempty"`
}
