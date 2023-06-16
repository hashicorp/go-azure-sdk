package recoverypoints

import "github.com/Azure/go-autorest/autorest"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RecoveryPointsClient struct {
	Client  autorest.Client
	baseUri string
}

func NewRecoveryPointsClientWithBaseURI(endpoint string) RecoveryPointsClient {
	return RecoveryPointsClient{
		Client:  autorest.NewClientWithUserAgent(userAgent()),
		baseUri: endpoint,
	}
}
