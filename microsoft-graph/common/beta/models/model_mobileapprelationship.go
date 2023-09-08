package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MobileAppRelationship struct {
	Id                   *string                          `json:"id,omitempty"`
	ODataType            *string                          `json:"@odata.type,omitempty"`
	TargetDisplayName    *string                          `json:"targetDisplayName,omitempty"`
	TargetDisplayVersion *string                          `json:"targetDisplayVersion,omitempty"`
	TargetId             *string                          `json:"targetId,omitempty"`
	TargetPublisher      *string                          `json:"targetPublisher,omitempty"`
	TargetType           *MobileAppRelationshipTargetType `json:"targetType,omitempty"`
}
