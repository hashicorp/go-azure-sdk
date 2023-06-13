package recoverypointsgetaccesstoken

import "github.com/Azure/go-autorest/autorest"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RecoveryPointsGetAccessTokenClient struct {
	Client  autorest.Client
	baseUri string
}

func NewRecoveryPointsGetAccessTokenClientWithBaseURI(endpoint string) RecoveryPointsGetAccessTokenClient {
	return RecoveryPointsGetAccessTokenClient{
		Client:  autorest.NewClientWithUserAgent(userAgent()),
		baseUri: endpoint,
	}
}
