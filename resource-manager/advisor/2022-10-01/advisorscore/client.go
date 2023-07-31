package advisorscore

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	"github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AdvisorScoreClient struct {
	Client *resourcemanager.Client
}

func NewAdvisorScoreClientWithBaseURI(api environments.Api) (*AdvisorScoreClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(api, "advisorscore", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating AdvisorScoreClient: %+v", err)
	}

	return &AdvisorScoreClient{
		Client: client,
	}, nil
}
