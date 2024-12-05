package workspacemanagergroups

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WorkspaceManagerGroupProperties struct {
	Description         *string  `json:"description,omitempty"`
	DisplayName         string   `json:"displayName"`
	MemberResourceNames []string `json:"memberResourceNames"`
}
