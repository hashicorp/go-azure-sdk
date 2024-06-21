package networkwatchers

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TroubleshootingResult struct {
	Code      *string                   `json:"code,omitempty"`
	EndTime   *string                   `json:"endTime,omitempty"`
	Results   *[]TroubleshootingDetails `json:"results,omitempty"`
	StartTime *string                   `json:"startTime,omitempty"`
}
