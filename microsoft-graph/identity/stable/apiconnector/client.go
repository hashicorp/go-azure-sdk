package apiconnector

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ApiConnectorClient struct {
	Client *msgraph.Client
}

func NewApiConnectorClientWithBaseURI(sdkApi sdkEnv.Api) (*ApiConnectorClient, error) {
	client, err := msgraph.NewClient(sdkApi, "apiconnector", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating ApiConnectorClient: %+v", err)
	}

	return &ApiConnectorClient{
		Client: client,
	}, nil
}
