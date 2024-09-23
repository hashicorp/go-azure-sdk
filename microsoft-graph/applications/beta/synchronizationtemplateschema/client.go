package synchronizationtemplateschema

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SynchronizationTemplateSchemaClient struct {
	Client *msgraph.Client
}

func NewSynchronizationTemplateSchemaClientWithBaseURI(sdkApi sdkEnv.Api) (*SynchronizationTemplateSchemaClient, error) {
	client, err := msgraph.NewClient(sdkApi, "synchronizationtemplateschema", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating SynchronizationTemplateSchemaClient: %+v", err)
	}

	return &SynchronizationTemplateSchemaClient{
		Client: client,
	}, nil
}
