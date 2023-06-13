package recoverypointscrr

import "github.com/Azure/go-autorest/autorest"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RecoveryPointsCrrClient struct {
	Client  autorest.Client
	baseUri string
}

func NewRecoveryPointsCrrClientWithBaseURI(endpoint string) RecoveryPointsCrrClient {
	return RecoveryPointsCrrClient{
		Client:  autorest.NewClientWithUserAgent(userAgent()),
		baseUri: endpoint,
	}
}
