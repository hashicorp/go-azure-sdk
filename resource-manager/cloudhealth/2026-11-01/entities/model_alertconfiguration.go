package entities

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AlertConfiguration struct {
	ActionGroupIds *[]string     `json:"actionGroupIds,omitempty"`
	Description    *string       `json:"description,omitempty"`
	Severity       AlertSeverity `json:"severity"`
}
