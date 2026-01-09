package managedapplogcollectionrequest

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ManagedAppLogCollectionRequestClient struct {
	Client *msgraph.Client
}

func NewManagedAppLogCollectionRequestClientWithBaseURI(sdkApi sdkEnv.Api) (*ManagedAppLogCollectionRequestClient, error) {
	client, err := msgraph.NewClient(sdkApi, "managedapplogcollectionrequest", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating ManagedAppLogCollectionRequestClient: %+v", err)
	}

	return &ManagedAppLogCollectionRequestClient{
		Client: client,
	}, nil
}
