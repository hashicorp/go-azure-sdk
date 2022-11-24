package migrationrecoverypoints

import "github.com/Azure/go-autorest/autorest"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MigrationRecoveryPointsClient struct {
	Client  autorest.Client
	baseUri string
}

func NewMigrationRecoveryPointsClientWithBaseURI(endpoint string) MigrationRecoveryPointsClient {
	return MigrationRecoveryPointsClient{
		Client:  autorest.NewClientWithUserAgent(userAgent()),
		baseUri: endpoint,
	}
}
