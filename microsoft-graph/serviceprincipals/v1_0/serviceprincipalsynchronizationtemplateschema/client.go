package serviceprincipalsynchronizationtemplateschema

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ServicePrincipalSynchronizationTemplateSchemaClient struct {
	Client *msgraph.Client
}

func NewServicePrincipalSynchronizationTemplateSchemaClientWithBaseURI(api sdkEnv.Api) (*ServicePrincipalSynchronizationTemplateSchemaClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "serviceprincipalsynchronizationtemplateschema", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating ServicePrincipalSynchronizationTemplateSchemaClient: %+v", err)
	}

	return &ServicePrincipalSynchronizationTemplateSchemaClient{
		Client: client,
	}, nil
}
