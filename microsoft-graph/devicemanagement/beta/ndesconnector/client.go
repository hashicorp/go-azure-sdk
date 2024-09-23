package ndesconnector

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type NdesConnectorClient struct {
	Client *msgraph.Client
}

func NewNdesConnectorClientWithBaseURI(sdkApi sdkEnv.Api) (*NdesConnectorClient, error) {
	client, err := msgraph.NewClient(sdkApi, "ndesconnector", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating NdesConnectorClient: %+v", err)
	}

	return &NdesConnectorClient{
		Client: client,
	}, nil
}
