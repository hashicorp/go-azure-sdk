package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IdentityGovernanceTimeBasedAttributeTrigger struct {
	ODataType          *string                                                        `json:"@odata.type,omitempty"`
	OffsetInDays       *int64                                                         `json:"offsetInDays,omitempty"`
	TimeBasedAttribute *IdentityGovernanceTimeBasedAttributeTriggerTimeBasedAttribute `json:"timeBasedAttribute,omitempty"`
}
