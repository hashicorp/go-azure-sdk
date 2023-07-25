package certificateordersdiagnostics

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Status struct {
	Message  *string        `json:"message,omitempty"`
	StatusId *InsightStatus `json:"statusId,omitempty"`
}
