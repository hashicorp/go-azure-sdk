package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MobileAppSupersedence struct {
	Id                   *string                                `json:"id,omitempty"`
	ODataType            *string                                `json:"@odata.type,omitempty"`
	SupersededAppCount   *int64                                 `json:"supersededAppCount,omitempty"`
	SupersedenceType     *MobileAppSupersedenceSupersedenceType `json:"supersedenceType,omitempty"`
	SupersedingAppCount  *int64                                 `json:"supersedingAppCount,omitempty"`
	TargetDisplayName    *string                                `json:"targetDisplayName,omitempty"`
	TargetDisplayVersion *string                                `json:"targetDisplayVersion,omitempty"`
	TargetId             *string                                `json:"targetId,omitempty"`
	TargetPublisher      *string                                `json:"targetPublisher,omitempty"`
	TargetType           *MobileAppSupersedenceTargetType       `json:"targetType,omitempty"`
}
