package updateruns

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UpdateStatus struct {
	CompletedTime *string      `json:"completedTime,omitempty"`
	Error         *ErrorDetail `json:"error,omitempty"`
	StartTime     *string      `json:"startTime,omitempty"`
	State         *UpdateState `json:"state,omitempty"`
}
