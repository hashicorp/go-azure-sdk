package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidDeviceOwnerSystemUpdateFreezePeriod struct {
	EndDay     *int64  `json:"endDay,omitempty"`
	EndMonth   *int64  `json:"endMonth,omitempty"`
	ODataType  *string `json:"@odata.type,omitempty"`
	StartDay   *int64  `json:"startDay,omitempty"`
	StartMonth *int64  `json:"startMonth,omitempty"`
}
