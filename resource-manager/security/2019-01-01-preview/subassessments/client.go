package subassessments

import "github.com/Azure/go-autorest/autorest"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SubAssessmentsClient struct {
	Client  autorest.Client
	baseUri string
}

func NewSubAssessmentsClientWithBaseURI(endpoint string) SubAssessmentsClient {
	return SubAssessmentsClient{
		Client:  autorest.NewClientWithUserAgent(userAgent()),
		baseUri: endpoint,
	}
}
