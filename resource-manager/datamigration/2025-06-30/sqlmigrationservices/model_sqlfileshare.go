package sqlmigrationservices

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SqlFileShare struct {
	Password *string `json:"password,omitempty"`
	Path     *string `json:"path,omitempty"`
	Username *string `json:"username,omitempty"`
}
