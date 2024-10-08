package standardoperation

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GetUserTablesOracleTaskInput struct {
	ConnectionInfo  OracleConnectionInfo `json:"connectionInfo"`
	SelectedSchemas []string             `json:"selectedSchemas"`
}
