package serviceprincipalowner

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ServicePrincipalOwnerClient struct {
	Client *msgraph.Client
}

func NewServicePrincipalOwnerClientWithBaseURI(api sdkEnv.Api) (*ServicePrincipalOwnerClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "serviceprincipalowner", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating ServicePrincipalOwnerClient: %+v", err)
	}

	return &ServicePrincipalOwnerClient{
		Client: client,
	}, nil
}
