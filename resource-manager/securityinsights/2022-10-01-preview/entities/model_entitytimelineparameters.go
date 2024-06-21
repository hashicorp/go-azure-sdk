package entities

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EntityTimelineParameters struct {
	EndTime        string                `json:"endTime"`
	Kinds          *[]EntityTimelineKind `json:"kinds,omitempty"`
	NumberOfBucket *int64                `json:"numberOfBucket,omitempty"`
	StartTime      string                `json:"startTime"`
}
