package serviceprincipalsynchronization

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ServicePrincipalSynchronizationClient struct {
	Client *msgraph.Client
}

func NewServicePrincipalSynchronizationClientWithBaseURI(api sdkEnv.Api) (*ServicePrincipalSynchronizationClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "serviceprincipalsynchronization", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating ServicePrincipalSynchronizationClient: %+v", err)
	}

	return &ServicePrincipalSynchronizationClient{
		Client: client,
	}, nil
}
