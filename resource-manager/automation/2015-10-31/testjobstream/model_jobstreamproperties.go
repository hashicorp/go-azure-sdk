package testjobstream

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type JobStreamProperties struct {
	JobStreamId *string                 `json:"jobStreamId,omitempty"`
	StreamText  *string                 `json:"streamText,omitempty"`
	StreamType  *JobStreamType          `json:"streamType,omitempty"`
	Summary     *string                 `json:"summary,omitempty"`
	Time        *string                 `json:"time,omitempty"`
	Value       *map[string]interface{} `json:"value,omitempty"`
}
