package event

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EventClient struct {
	Client *msgraph.Client
}

func NewEventClientWithBaseURI(sdkApi sdkEnv.Api) (*EventClient, error) {
	client, err := msgraph.NewClient(sdkApi, "event", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating EventClient: %+v", err)
	}

	return &EventClient{
		Client: client,
	}, nil
}
