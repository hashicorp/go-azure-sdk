package incidentteam

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TeamProperties struct {
	GroupIds        *[]string `json:"groupIds,omitempty"`
	MemberIds       *[]string `json:"memberIds,omitempty"`
	TeamDescription *string   `json:"teamDescription,omitempty"`
	TeamName        string    `json:"teamName"`
}
