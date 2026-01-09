package thread

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ThreadClient struct {
	Client *msgraph.Client
}

func NewThreadClientWithBaseURI(sdkApi sdkEnv.Api) (*ThreadClient, error) {
	client, err := msgraph.NewClient(sdkApi, "thread", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating ThreadClient: %+v", err)
	}

	return &ThreadClient{
		Client: client,
	}, nil
}
