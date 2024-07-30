package stable

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type StaffAvailabilityItem struct {
	AvailabilityItems *[]AvailabilityItem `json:"availabilityItems,omitempty"`
	ODataType         *string             `json:"@odata.type,omitempty"`
	StaffId           *string             `json:"staffId,omitempty"`
}
