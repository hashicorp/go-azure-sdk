package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MembershipRuleProcessingStatus struct {
	ErrorMessage          *string                               `json:"errorMessage,omitempty"`
	LastMembershipUpdated *string                               `json:"lastMembershipUpdated,omitempty"`
	ODataType             *string                               `json:"@odata.type,omitempty"`
	Status                *MembershipRuleProcessingStatusStatus `json:"status,omitempty"`
}
