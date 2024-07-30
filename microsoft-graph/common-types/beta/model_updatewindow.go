package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UpdateWindow struct {
	ODataType             *string `json:"@odata.type,omitempty"`
	UpdateWindowEndTime   *string `json:"updateWindowEndTime,omitempty"`
	UpdateWindowStartTime *string `json:"updateWindowStartTime,omitempty"`
}
