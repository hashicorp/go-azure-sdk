package endpoint

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EndpointClient struct {
	Client *msgraph.Client
}

func NewEndpointClientWithBaseURI(sdkApi sdkEnv.Api) (*EndpointClient, error) {
	client, err := msgraph.NewClient(sdkApi, "endpoint", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating EndpointClient: %+v", err)
	}

	return &EndpointClient{
		Client: client,
	}, nil
}
