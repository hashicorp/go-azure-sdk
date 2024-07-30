package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type VirtualEventRegistrationPredefinedQuestion struct {
	DisplayName *string                                          `json:"displayName,omitempty"`
	Id          *string                                          `json:"id,omitempty"`
	IsRequired  *bool                                            `json:"isRequired,omitempty"`
	Label       *VirtualEventRegistrationPredefinedQuestionLabel `json:"label,omitempty"`
	ODataType   *string                                          `json:"@odata.type,omitempty"`
}
