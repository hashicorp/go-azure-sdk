package synchronizationjobschema

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SynchronizationJobSchemaClient struct {
	Client *msgraph.Client
}

func NewSynchronizationJobSchemaClientWithBaseURI(sdkApi sdkEnv.Api) (*SynchronizationJobSchemaClient, error) {
	client, err := msgraph.NewClient(sdkApi, "synchronizationjobschema", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating SynchronizationJobSchemaClient: %+v", err)
	}

	return &SynchronizationJobSchemaClient{
		Client: client,
	}, nil
}
