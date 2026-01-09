package synchronizationjobschemadirectory

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SynchronizationJobSchemaDirectoryClient struct {
	Client *msgraph.Client
}

func NewSynchronizationJobSchemaDirectoryClientWithBaseURI(sdkApi sdkEnv.Api) (*SynchronizationJobSchemaDirectoryClient, error) {
	client, err := msgraph.NewClient(sdkApi, "synchronizationjobschemadirectory", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating SynchronizationJobSchemaDirectoryClient: %+v", err)
	}

	return &SynchronizationJobSchemaDirectoryClient{
		Client: client,
	}, nil
}
