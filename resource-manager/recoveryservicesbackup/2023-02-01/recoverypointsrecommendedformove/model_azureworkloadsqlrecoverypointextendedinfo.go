package recoverypointsrecommendedformove

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AzureWorkloadSQLRecoveryPointExtendedInfo struct {
	DataDirectoryPaths     *[]SQLDataDirectory `json:"dataDirectoryPaths,omitempty"`
	DataDirectoryTimeInUTC *string             `json:"dataDirectoryTimeInUTC,omitempty"`
}
