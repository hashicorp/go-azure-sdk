package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ConditionalAccessPolicyDetail struct {
	Conditions      *ConditionalAccessConditionSet    `json:"conditions,omitempty"`
	GrantControls   *ConditionalAccessGrantControls   `json:"grantControls,omitempty"`
	ODataType       *string                           `json:"@odata.type,omitempty"`
	SessionControls *ConditionalAccessSessionControls `json:"sessionControls,omitempty"`
}
