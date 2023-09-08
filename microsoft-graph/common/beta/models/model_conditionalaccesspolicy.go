package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ConditionalAccessPolicy struct {
	Conditions       *ConditionalAccessConditionSet    `json:"conditions,omitempty"`
	CreatedDateTime  *string                           `json:"createdDateTime,omitempty"`
	Description      *string                           `json:"description,omitempty"`
	DisplayName      *string                           `json:"displayName,omitempty"`
	GrantControls    *ConditionalAccessGrantControls   `json:"grantControls,omitempty"`
	Id               *string                           `json:"id,omitempty"`
	ModifiedDateTime *string                           `json:"modifiedDateTime,omitempty"`
	ODataType        *string                           `json:"@odata.type,omitempty"`
	SessionControls  *ConditionalAccessSessionControls `json:"sessionControls,omitempty"`
	State            *ConditionalAccessPolicyState     `json:"state,omitempty"`
}
