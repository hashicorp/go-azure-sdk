package presence

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PresenceClient struct {
	Client *msgraph.Client
}

func NewPresenceClientWithBaseURI(sdkApi sdkEnv.Api) (*PresenceClient, error) {
	client, err := msgraph.NewClient(sdkApi, "presence", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating PresenceClient: %+v", err)
	}

	return &PresenceClient{
		Client: client,
	}, nil
}
