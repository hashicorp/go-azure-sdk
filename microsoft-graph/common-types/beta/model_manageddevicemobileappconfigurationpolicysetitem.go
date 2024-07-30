package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ManagedDeviceMobileAppConfigurationPolicySetItem struct {
	CreatedDateTime      *string                                                    `json:"createdDateTime,omitempty"`
	DisplayName          *string                                                    `json:"displayName,omitempty"`
	ErrorCode            *ManagedDeviceMobileAppConfigurationPolicySetItemErrorCode `json:"errorCode,omitempty"`
	GuidedDeploymentTags *[]string                                                  `json:"guidedDeploymentTags,omitempty"`
	Id                   *string                                                    `json:"id,omitempty"`
	ItemType             *string                                                    `json:"itemType,omitempty"`
	LastModifiedDateTime *string                                                    `json:"lastModifiedDateTime,omitempty"`
	ODataType            *string                                                    `json:"@odata.type,omitempty"`
	PayloadId            *string                                                    `json:"payloadId,omitempty"`
	Status               *ManagedDeviceMobileAppConfigurationPolicySetItemStatus    `json:"status,omitempty"`
}
