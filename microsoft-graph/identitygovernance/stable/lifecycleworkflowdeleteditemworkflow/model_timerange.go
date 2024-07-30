package lifecycleworkflowdeleteditemworkflow

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TimeRange struct {
	EndTime   *string `json:"endTime,omitempty"`
	ODataType *string `json:"@odata.type,omitempty"`
	StartTime *string `json:"startTime,omitempty"`
}
