package serviceprincipalmemberof

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ServicePrincipalMemberOfClient struct {
	Client *msgraph.Client
}

func NewServicePrincipalMemberOfClientWithBaseURI(api sdkEnv.Api) (*ServicePrincipalMemberOfClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "serviceprincipalmemberof", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating ServicePrincipalMemberOfClient: %+v", err)
	}

	return &ServicePrincipalMemberOfClient{
		Client: client,
	}, nil
}
