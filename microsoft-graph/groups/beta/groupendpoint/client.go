package groupendpoint

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GroupEndpointClient struct {
	Client *msgraph.Client
}

func NewGroupEndpointClientWithBaseURI(api sdkEnv.Api) (*GroupEndpointClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "groupendpoint", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating GroupEndpointClient: %+v", err)
	}

	return &GroupEndpointClient{
		Client: client,
	}, nil
}
