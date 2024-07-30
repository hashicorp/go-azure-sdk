package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type BusinessFlow struct {
	CustomData       *string               `json:"customData,omitempty"`
	DeDuplicationId  *string               `json:"deDuplicationId,omitempty"`
	Description      *string               `json:"description,omitempty"`
	DisplayName      *string               `json:"displayName,omitempty"`
	Id               *string               `json:"id,omitempty"`
	ODataType        *string               `json:"@odata.type,omitempty"`
	Policy           *GovernancePolicy     `json:"policy,omitempty"`
	PolicyTemplateId *string               `json:"policyTemplateId,omitempty"`
	RecordVersion    *string               `json:"recordVersion,omitempty"`
	SchemaId         *string               `json:"schemaId,omitempty"`
	Settings         *BusinessFlowSettings `json:"settings,omitempty"`
}
