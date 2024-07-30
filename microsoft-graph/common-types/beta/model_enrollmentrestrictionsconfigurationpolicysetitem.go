package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EnrollmentRestrictionsConfigurationPolicySetItem struct {
	CreatedDateTime      *string                                                    `json:"createdDateTime,omitempty"`
	DisplayName          *string                                                    `json:"displayName,omitempty"`
	ErrorCode            *EnrollmentRestrictionsConfigurationPolicySetItemErrorCode `json:"errorCode,omitempty"`
	GuidedDeploymentTags *[]string                                                  `json:"guidedDeploymentTags,omitempty"`
	Id                   *string                                                    `json:"id,omitempty"`
	ItemType             *string                                                    `json:"itemType,omitempty"`
	LastModifiedDateTime *string                                                    `json:"lastModifiedDateTime,omitempty"`
	Limit                *int64                                                     `json:"limit,omitempty"`
	ODataType            *string                                                    `json:"@odata.type,omitempty"`
	PayloadId            *string                                                    `json:"payloadId,omitempty"`
	Priority             *int64                                                     `json:"priority,omitempty"`
	Status               *EnrollmentRestrictionsConfigurationPolicySetItemStatus    `json:"status,omitempty"`
}
