package sqlmigrationservices

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SqlDbMigrationStatusDetails struct {
	ListOfCopyProgressDetails *[]CopyProgressDetails `json:"listOfCopyProgressDetails,omitempty"`
	MigrationState            *string                `json:"migrationState,omitempty"`
	SqlDataCopyErrors         *[]string              `json:"sqlDataCopyErrors,omitempty"`
}
