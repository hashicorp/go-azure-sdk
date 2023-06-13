package backupcrrjobs

import "github.com/Azure/go-autorest/autorest"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type BackupCrrJobsClient struct {
	Client  autorest.Client
	baseUri string
}

func NewBackupCrrJobsClientWithBaseURI(endpoint string) BackupCrrJobsClient {
	return BackupCrrJobsClient{
		Client:  autorest.NewClientWithUserAgent(userAgent()),
		baseUri: endpoint,
	}
}
