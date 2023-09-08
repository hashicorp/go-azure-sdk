package serviceprincipalsynchronizationsecret

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ServicePrincipalSynchronizationSecretClient struct {
	Client *msgraph.Client
}

func NewServicePrincipalSynchronizationSecretClientWithBaseURI(api sdkEnv.Api) (*ServicePrincipalSynchronizationSecretClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "serviceprincipalsynchronizationsecret", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating ServicePrincipalSynchronizationSecretClient: %+v", err)
	}

	return &ServicePrincipalSynchronizationSecretClient{
		Client: client,
	}, nil
}
