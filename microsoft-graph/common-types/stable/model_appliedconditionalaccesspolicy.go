package stable

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AppliedConditionalAccessPolicy struct {
	DisplayName             *string                               `json:"displayName,omitempty"`
	EnforcedGrantControls   *[]string                             `json:"enforcedGrantControls,omitempty"`
	EnforcedSessionControls *[]string                             `json:"enforcedSessionControls,omitempty"`
	Id                      *string                               `json:"id,omitempty"`
	ODataType               *string                               `json:"@odata.type,omitempty"`
	Result                  *AppliedConditionalAccessPolicyResult `json:"result,omitempty"`
}
