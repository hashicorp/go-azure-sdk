package cognitiveservicescommitmentplans

import "github.com/Azure/go-autorest/autorest"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CognitiveServicesCommitmentPlansClient struct {
	Client  autorest.Client
	baseUri string
}

func NewCognitiveServicesCommitmentPlansClientWithBaseURI(endpoint string) CognitiveServicesCommitmentPlansClient {
	return CognitiveServicesCommitmentPlansClient{
		Client:  autorest.NewClientWithUserAgent(userAgent()),
		baseUri: endpoint,
	}
}
