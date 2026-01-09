package onlinemeeting

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type OnlineMeetingClient struct {
	Client *msgraph.Client
}

func NewOnlineMeetingClientWithBaseURI(sdkApi sdkEnv.Api) (*OnlineMeetingClient, error) {
	client, err := msgraph.NewClient(sdkApi, "onlinemeeting", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating OnlineMeetingClient: %+v", err)
	}

	return &OnlineMeetingClient{
		Client: client,
	}, nil
}
