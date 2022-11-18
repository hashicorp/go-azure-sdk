package benefitrecommendations

import "github.com/Azure/go-autorest/autorest"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type BenefitRecommendationsClient struct {
	Client  autorest.Client
	baseUri string
}

func NewBenefitRecommendationsClientWithBaseURI(endpoint string) BenefitRecommendationsClient {
	return BenefitRecommendationsClient{
		Client:  autorest.NewClientWithUserAgent(userAgent()),
		baseUri: endpoint,
	}
}
