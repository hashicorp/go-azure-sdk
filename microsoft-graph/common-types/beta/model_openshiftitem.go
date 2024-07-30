package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type OpenShiftItem struct {
	Activities    *[]ShiftActivity    `json:"activities,omitempty"`
	DisplayName   *string             `json:"displayName,omitempty"`
	EndDateTime   *string             `json:"endDateTime,omitempty"`
	Notes         *string             `json:"notes,omitempty"`
	ODataType     *string             `json:"@odata.type,omitempty"`
	OpenSlotCount *int64              `json:"openSlotCount,omitempty"`
	StartDateTime *string             `json:"startDateTime,omitempty"`
	Theme         *OpenShiftItemTheme `json:"theme,omitempty"`
}
