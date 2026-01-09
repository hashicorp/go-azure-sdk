package chatinstalledapp

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ChatInstalledAppClient struct {
	Client *msgraph.Client
}

func NewChatInstalledAppClientWithBaseURI(sdkApi sdkEnv.Api) (*ChatInstalledAppClient, error) {
	client, err := msgraph.NewClient(sdkApi, "chatinstalledapp", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating ChatInstalledAppClient: %+v", err)
	}

	return &ChatInstalledAppClient{
		Client: client,
	}, nil
}
