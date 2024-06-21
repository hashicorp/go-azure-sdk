package webapps

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MSDeployLogEntry struct {
	Message *string               `json:"message,omitempty"`
	Time    *string               `json:"time,omitempty"`
	Type    *MSDeployLogEntryType `json:"type,omitempty"`
}
