package entities

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TimelineError struct {
	ErrorMessage string             `json:"errorMessage"`
	Kind         EntityTimelineKind `json:"kind"`
	QueryId      *string            `json:"queryId,omitempty"`
}
