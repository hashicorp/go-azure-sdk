package checkprincipalaccess

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SubjectInfo struct {
	GroupIds    *[]string `json:"groupIds,omitempty"`
	PrincipalId string    `json:"principalId"`
}
