package workflowrunactions

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RequestHistoryProperties struct {
	EndTime   *string   `json:"endTime,omitempty"`
	Request   *Request  `json:"request,omitempty"`
	Response  *Response `json:"response,omitempty"`
	StartTime *string   `json:"startTime,omitempty"`
}
