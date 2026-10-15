package updateruns

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UpdateGroupStatus struct {
	AfterGates     *[]UpdateRunGateStatus `json:"afterGates,omitempty"`
	BeforeGates    *[]UpdateRunGateStatus `json:"beforeGates,omitempty"`
	MaxConcurrency *int64                 `json:"maxConcurrency,omitempty"`
	Members        *[]MemberUpdateStatus  `json:"members,omitempty"`
	Name           *string                `json:"name,omitempty"`
	Status         *UpdateStatus          `json:"status,omitempty"`
}
