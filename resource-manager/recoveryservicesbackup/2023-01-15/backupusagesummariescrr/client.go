package backupusagesummariescrr

import "github.com/Azure/go-autorest/autorest"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type BackupUsageSummariesCRRClient struct {
	Client  autorest.Client
	baseUri string
}

func NewBackupUsageSummariesCRRClientWithBaseURI(endpoint string) BackupUsageSummariesCRRClient {
	return BackupUsageSummariesCRRClient{
		Client:  autorest.NewClientWithUserAgent(userAgent()),
		baseUri: endpoint,
	}
}
