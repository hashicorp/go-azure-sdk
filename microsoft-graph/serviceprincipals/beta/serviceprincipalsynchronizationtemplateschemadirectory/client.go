package serviceprincipalsynchronizationtemplateschemadirectory

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ServicePrincipalSynchronizationTemplateSchemaDirectoryClient struct {
	Client *msgraph.Client
}

func NewServicePrincipalSynchronizationTemplateSchemaDirectoryClientWithBaseURI(api sdkEnv.Api) (*ServicePrincipalSynchronizationTemplateSchemaDirectoryClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "serviceprincipalsynchronizationtemplateschemadirectory", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating ServicePrincipalSynchronizationTemplateSchemaDirectoryClient: %+v", err)
	}

	return &ServicePrincipalSynchronizationTemplateSchemaDirectoryClient{
		Client: client,
	}, nil
}
