package integrationaccounts

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TrackingEvent struct {
	Error      *TrackingEventErrorInfo `json:"error,omitempty"`
	EventLevel EventLevel              `json:"eventLevel"`
	EventTime  string                  `json:"eventTime"`
	Record     *interface{}            `json:"record,omitempty"`
	RecordType TrackingRecordType      `json:"recordType"`
}
