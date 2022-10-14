package locations

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Trial struct {
	AvailableHosts *int64       `json:"availableHosts,omitempty"`
	Status         *TrialStatus `json:"status,omitempty"`
}
