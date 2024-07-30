package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecureScoreControlProfile struct {
	ActionType            *string                          `json:"actionType,omitempty"`
	ActionUrl             *string                          `json:"actionUrl,omitempty"`
	AzureTenantId         *string                          `json:"azureTenantId,omitempty"`
	ComplianceInformation *[]ComplianceInformation         `json:"complianceInformation,omitempty"`
	ControlCategory       *string                          `json:"controlCategory,omitempty"`
	ControlStateUpdates   *[]SecureScoreControlStateUpdate `json:"controlStateUpdates,omitempty"`
	Deprecated            *bool                            `json:"deprecated,omitempty"`
	Id                    *string                          `json:"id,omitempty"`
	ImplementationCost    *string                          `json:"implementationCost,omitempty"`
	LastModifiedDateTime  *string                          `json:"lastModifiedDateTime,omitempty"`
	MaxScore              *float64                         `json:"maxScore,omitempty"`
	ODataType             *string                          `json:"@odata.type,omitempty"`
	Rank                  *int64                           `json:"rank,omitempty"`
	Remediation           *string                          `json:"remediation,omitempty"`
	RemediationImpact     *string                          `json:"remediationImpact,omitempty"`
	Service               *string                          `json:"service,omitempty"`
	Threats               *[]string                        `json:"threats,omitempty"`
	Tier                  *string                          `json:"tier,omitempty"`
	Title                 *string                          `json:"title,omitempty"`
	UserImpact            *string                          `json:"userImpact,omitempty"`
	VendorInformation     *SecurityVendorInformation       `json:"vendorInformation,omitempty"`
}
