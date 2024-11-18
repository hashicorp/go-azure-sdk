package replicationeligibilityresults

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ReplicationEligibilityResultsClient struct {
	Client *resourcemanager.Client
}

func NewReplicationEligibilityResultsClientWithBaseURI(sdkApi sdkEnv.Api) (*ReplicationEligibilityResultsClient, error) {
	client, err := resourcemanager.NewClient(sdkApi, "replicationeligibilityresults", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating ReplicationEligibilityResultsClient: %+v", err)
	}

	return &ReplicationEligibilityResultsClient{
		Client: client,
	}, nil
}
