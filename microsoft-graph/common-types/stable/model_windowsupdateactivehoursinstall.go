package stable

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsUpdateActiveHoursInstall struct {
	ActiveHoursEnd   *string `json:"activeHoursEnd,omitempty"`
	ActiveHoursStart *string `json:"activeHoursStart,omitempty"`
	ODataType        *string `json:"@odata.type,omitempty"`
}
