package recoverypointsrecommendedformove

import "github.com/Azure/go-autorest/autorest"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RecoveryPointsRecommendedForMoveClient struct {
	Client  autorest.Client
	baseUri string
}

func NewRecoveryPointsRecommendedForMoveClientWithBaseURI(endpoint string) RecoveryPointsRecommendedForMoveClient {
	return RecoveryPointsRecommendedForMoveClient{
		Client:  autorest.NewClientWithUserAgent(userAgent()),
		baseUri: endpoint,
	}
}
