package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RetentionSetting struct {
	Interval  *string `json:"interval,omitempty"`
	ODataType *string `json:"@odata.type,omitempty"`
	Period    *string `json:"period,omitempty"`
}
