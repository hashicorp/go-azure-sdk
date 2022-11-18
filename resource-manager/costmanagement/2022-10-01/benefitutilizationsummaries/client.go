package benefitutilizationsummaries

import "github.com/Azure/go-autorest/autorest"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type BenefitUtilizationSummariesClient struct {
	Client  autorest.Client
	baseUri string
}

func NewBenefitUtilizationSummariesClientWithBaseURI(endpoint string) BenefitUtilizationSummariesClient {
	return BenefitUtilizationSummariesClient{
		Client:  autorest.NewClientWithUserAgent(userAgent()),
		baseUri: endpoint,
	}
}
