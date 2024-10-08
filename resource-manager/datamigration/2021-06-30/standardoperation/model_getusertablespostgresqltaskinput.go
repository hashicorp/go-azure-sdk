package standardoperation

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GetUserTablesPostgreSqlTaskInput struct {
	ConnectionInfo    PostgreSqlConnectionInfo `json:"connectionInfo"`
	SelectedDatabases []string                 `json:"selectedDatabases"`
}
