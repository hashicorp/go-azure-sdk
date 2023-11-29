package dppjob

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CrossRegionRestoreJobsRequest struct {
	SourceBackupVaultId string `json:"sourceBackupVaultId"`
	SourceRegion        string `json:"sourceRegion"`
}
