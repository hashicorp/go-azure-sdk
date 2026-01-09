package onpremisessynchronization

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type OnPremisesSynchronizationClient struct {
	Client *msgraph.Client
}

func NewOnPremisesSynchronizationClientWithBaseURI(sdkApi sdkEnv.Api) (*OnPremisesSynchronizationClient, error) {
	client, err := msgraph.NewClient(sdkApi, "onpremisessynchronization", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating OnPremisesSynchronizationClient: %+v", err)
	}

	return &OnPremisesSynchronizationClient{
		Client: client,
	}, nil
}
