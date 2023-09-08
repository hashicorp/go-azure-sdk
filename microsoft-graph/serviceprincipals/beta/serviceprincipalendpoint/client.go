package serviceprincipalendpoint

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ServicePrincipalEndpointClient struct {
	Client *msgraph.Client
}

func NewServicePrincipalEndpointClientWithBaseURI(api sdkEnv.Api) (*ServicePrincipalEndpointClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "serviceprincipalendpoint", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating ServicePrincipalEndpointClient: %+v", err)
	}

	return &ServicePrincipalEndpointClient{
		Client: client,
	}, nil
}
