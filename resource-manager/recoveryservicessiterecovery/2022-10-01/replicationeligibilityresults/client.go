package replicationeligibilityresults

import "github.com/Azure/go-autorest/autorest"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ReplicationEligibilityResultsClient struct {
	Client  autorest.Client
	baseUri string
}

func NewReplicationEligibilityResultsClientWithBaseURI(endpoint string) ReplicationEligibilityResultsClient {
	return ReplicationEligibilityResultsClient{
		Client:  autorest.NewClientWithUserAgent(userAgent()),
		baseUri: endpoint,
	}
}
