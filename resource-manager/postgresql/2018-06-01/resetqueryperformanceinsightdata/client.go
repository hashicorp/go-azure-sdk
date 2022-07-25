package resetqueryperformanceinsightdata

import "github.com/Azure/go-autorest/autorest"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ResetQueryPerformanceInsightDataClient struct {
	Client  autorest.Client
	baseUri string
}

func NewResetQueryPerformanceInsightDataClientWithBaseURI(endpoint string) ResetQueryPerformanceInsightDataClient {
	return ResetQueryPerformanceInsightDataClient{
		Client:  autorest.NewClientWithUserAgent(userAgent()),
		baseUri: endpoint,
	}
}
