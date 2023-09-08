package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityAction struct {
	ActionReason       *string                    `json:"actionReason,omitempty"`
	AppId              *string                    `json:"appId,omitempty"`
	AzureTenantId      *string                    `json:"azureTenantId,omitempty"`
	ClientContext      *string                    `json:"clientContext,omitempty"`
	CompletedDateTime  *string                    `json:"completedDateTime,omitempty"`
	CreatedDateTime    *string                    `json:"createdDateTime,omitempty"`
	ErrorInfo          *ResultInfo                `json:"errorInfo,omitempty"`
	Id                 *string                    `json:"id,omitempty"`
	LastActionDateTime *string                    `json:"lastActionDateTime,omitempty"`
	Name               *string                    `json:"name,omitempty"`
	ODataType          *string                    `json:"@odata.type,omitempty"`
	Parameters         *[]KeyValuePair            `json:"parameters,omitempty"`
	States             *[]SecurityActionState     `json:"states,omitempty"`
	Status             *SecurityActionStatus      `json:"status,omitempty"`
	User               *string                    `json:"user,omitempty"`
	VendorInformation  *SecurityVendorInformation `json:"vendorInformation,omitempty"`
}
