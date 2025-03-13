package fetchsecondaryrecoverypoints

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type FetchSecondaryRPsRequestParameters struct {
	SourceBackupInstanceId *string `json:"sourceBackupInstanceId,omitempty"`
	SourceRegion           *string `json:"sourceRegion,omitempty"`
}
