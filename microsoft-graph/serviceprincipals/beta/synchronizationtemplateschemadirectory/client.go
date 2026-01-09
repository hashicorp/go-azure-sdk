package synchronizationtemplateschemadirectory

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SynchronizationTemplateSchemaDirectoryClient struct {
	Client *msgraph.Client
}

func NewSynchronizationTemplateSchemaDirectoryClientWithBaseURI(sdkApi sdkEnv.Api) (*SynchronizationTemplateSchemaDirectoryClient, error) {
	client, err := msgraph.NewClient(sdkApi, "synchronizationtemplateschemadirectory", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating SynchronizationTemplateSchemaDirectoryClient: %+v", err)
	}

	return &SynchronizationTemplateSchemaDirectoryClient{
		Client: client,
	}, nil
}
