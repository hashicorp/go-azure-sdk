package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidDeviceOwnerKioskModeApp struct {
	ClassName *string `json:"className,omitempty"`
	ODataType *string `json:"@odata.type,omitempty"`
	Package   *string `json:"package,omitempty"`
}
