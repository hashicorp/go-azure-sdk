package applicationconnectorgroup

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ApplicationConnectorGroupClient struct {
	Client *msgraph.Client
}

func NewApplicationConnectorGroupClientWithBaseURI(api sdkEnv.Api) (*ApplicationConnectorGroupClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "applicationconnectorgroup", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating ApplicationConnectorGroupClient: %+v", err)
	}

	return &ApplicationConnectorGroupClient{
		Client: client,
	}, nil
}
