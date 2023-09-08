package serviceprincipalcreatedobject

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ServicePrincipalCreatedObjectClient struct {
	Client *msgraph.Client
}

func NewServicePrincipalCreatedObjectClientWithBaseURI(api sdkEnv.Api) (*ServicePrincipalCreatedObjectClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "serviceprincipalcreatedobject", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating ServicePrincipalCreatedObjectClient: %+v", err)
	}

	return &ServicePrincipalCreatedObjectClient{
		Client: client,
	}, nil
}
