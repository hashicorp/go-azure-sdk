package backupusagesummaries

import "github.com/Azure/go-autorest/autorest"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type BackupUsageSummariesClient struct {
	Client  autorest.Client
	baseUri string
}

func NewBackupUsageSummariesClientWithBaseURI(endpoint string) BackupUsageSummariesClient {
	return BackupUsageSummariesClient{
		Client:  autorest.NewClientWithUserAgent(userAgent()),
		baseUri: endpoint,
	}
}
