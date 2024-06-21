package hybridrunbookworkergroup

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type HybridRunbookWorkerLegacy struct {
	IP               *string `json:"ip,omitempty"`
	LastSeenDateTime *string `json:"lastSeenDateTime,omitempty"`
	Name             *string `json:"name,omitempty"`
	RegistrationTime *string `json:"registrationTime,omitempty"`
}
