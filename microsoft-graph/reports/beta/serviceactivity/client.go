package serviceactivity

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ServiceActivityClient struct {
	Client *msgraph.Client
}

func NewServiceActivityClientWithBaseURI(sdkApi sdkEnv.Api) (*ServiceActivityClient, error) {
	client, err := msgraph.NewClient(sdkApi, "serviceactivity", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating ServiceActivityClient: %+v", err)
	}

	return &ServiceActivityClient{
		Client: client,
	}, nil
}
