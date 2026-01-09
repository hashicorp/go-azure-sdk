package partnerbillingreconciliationbilled

import (
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/stable"
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CreatePartnerBillingReconciliationBilledExportRequest struct {
	AttributeSet *stable.PartnersBillingAttributeSet `json:"attributeSet,omitempty"`
	InvoiceId    nullable.Type[string]               `json:"invoiceId,omitempty"`
}
