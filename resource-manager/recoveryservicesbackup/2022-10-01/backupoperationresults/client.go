package backupoperationresults

import "github.com/Azure/go-autorest/autorest"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type BackupOperationResultsClient struct {
	Client  autorest.Client
	baseUri string
}

func NewBackupOperationResultsClientWithBaseURI(endpoint string) BackupOperationResultsClient {
	return BackupOperationResultsClient{
		Client:  autorest.NewClientWithUserAgent(userAgent()),
		baseUri: endpoint,
	}
}
