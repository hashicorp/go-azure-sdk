package serviceprincipalsynchronizationjobschemadirectory

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ServicePrincipalSynchronizationJobSchemaDirectoryClient struct {
	Client *msgraph.Client
}

func NewServicePrincipalSynchronizationJobSchemaDirectoryClientWithBaseURI(api sdkEnv.Api) (*ServicePrincipalSynchronizationJobSchemaDirectoryClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "serviceprincipalsynchronizationjobschemadirectory", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating ServicePrincipalSynchronizationJobSchemaDirectoryClient: %+v", err)
	}

	return &ServicePrincipalSynchronizationJobSchemaDirectoryClient{
		Client: client,
	}, nil
}
