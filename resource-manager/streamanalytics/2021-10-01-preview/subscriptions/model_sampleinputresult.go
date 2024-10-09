package subscriptions

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SampleInputResult struct {
	Diagnostics       *[]string                `json:"diagnostics,omitempty"`
	Error             *ErrorError              `json:"error,omitempty"`
	EventsDownloadURL *string                  `json:"eventsDownloadUrl,omitempty"`
	LastArrivalTime   *string                  `json:"lastArrivalTime,omitempty"`
	Status            *SampleInputResultStatus `json:"status,omitempty"`
}
