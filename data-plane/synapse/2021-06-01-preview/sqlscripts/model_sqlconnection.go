package sqlscripts

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SqlConnection struct {
	DatabaseName *string            `json:"databaseName,omitempty"`
	Name         *string            `json:"name,omitempty"`
	PoolName     *string            `json:"poolName,omitempty"`
	Type         *SqlConnectionType `json:"type,omitempty"`
}
