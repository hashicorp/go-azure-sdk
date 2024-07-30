package lifecycleworkflowworkflow

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TimeSlot struct {
	End       *DateTimeTimeZone `json:"end,omitempty"`
	ODataType *string           `json:"@odata.type,omitempty"`
	Start     *DateTimeTimeZone `json:"start,omitempty"`
}
