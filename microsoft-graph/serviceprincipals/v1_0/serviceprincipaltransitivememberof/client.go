package serviceprincipaltransitivememberof

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ServicePrincipalTransitiveMemberOfClient struct {
	Client *msgraph.Client
}

func NewServicePrincipalTransitiveMemberOfClientWithBaseURI(api sdkEnv.Api) (*ServicePrincipalTransitiveMemberOfClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "serviceprincipaltransitivememberof", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating ServicePrincipalTransitiveMemberOfClient: %+v", err)
	}

	return &ServicePrincipalTransitiveMemberOfClient{
		Client: client,
	}, nil
}
