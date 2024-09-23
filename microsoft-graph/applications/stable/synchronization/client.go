package synchronization

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SynchronizationClient struct {
	Client *msgraph.Client
}

func NewSynchronizationClientWithBaseURI(sdkApi sdkEnv.Api) (*SynchronizationClient, error) {
	client, err := msgraph.NewClient(sdkApi, "synchronization", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating SynchronizationClient: %+v", err)
	}

	return &SynchronizationClient{
		Client: client,
	}, nil
}
