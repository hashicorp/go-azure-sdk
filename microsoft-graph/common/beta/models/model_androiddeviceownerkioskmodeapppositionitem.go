package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidDeviceOwnerKioskModeAppPositionItem struct {
	Item      *AndroidDeviceOwnerKioskModeHomeScreenItem `json:"item,omitempty"`
	ODataType *string                                    `json:"@odata.type,omitempty"`
	Position  *int64                                     `json:"position,omitempty"`
}
