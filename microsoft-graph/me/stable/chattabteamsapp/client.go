package chattabteamsapp

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ChatTabTeamsAppClient struct {
	Client *msgraph.Client
}

func NewChatTabTeamsAppClientWithBaseURI(sdkApi sdkEnv.Api) (*ChatTabTeamsAppClient, error) {
	client, err := msgraph.NewClient(sdkApi, "chattabteamsapp", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating ChatTabTeamsAppClient: %+v", err)
	}

	return &ChatTabTeamsAppClient{
		Client: client,
	}, nil
}
