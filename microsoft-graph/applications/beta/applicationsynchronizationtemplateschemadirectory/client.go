package applicationsynchronizationtemplateschemadirectory

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ApplicationSynchronizationTemplateSchemaDirectoryClient struct {
	Client *msgraph.Client
}

func NewApplicationSynchronizationTemplateSchemaDirectoryClientWithBaseURI(api sdkEnv.Api) (*ApplicationSynchronizationTemplateSchemaDirectoryClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "applicationsynchronizationtemplateschemadirectory", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating ApplicationSynchronizationTemplateSchemaDirectoryClient: %+v", err)
	}

	return &ApplicationSynchronizationTemplateSchemaDirectoryClient{
		Client: client,
	}, nil
}
