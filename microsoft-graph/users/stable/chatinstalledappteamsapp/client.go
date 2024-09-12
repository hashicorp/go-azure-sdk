package chatinstalledappteamsapp

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ChatInstalledAppTeamsAppClient struct {
	Client *msgraph.Client
}

func NewChatInstalledAppTeamsAppClientWithBaseURI(sdkApi sdkEnv.Api) (*ChatInstalledAppTeamsAppClient, error) {
	client, err := msgraph.NewClient(sdkApi, "chatinstalledappteamsapp", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating ChatInstalledAppTeamsAppClient: %+v", err)
	}

	return &ChatInstalledAppTeamsAppClient{
		Client: client,
	}, nil
}
