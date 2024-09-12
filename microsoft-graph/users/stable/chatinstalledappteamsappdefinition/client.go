package chatinstalledappteamsappdefinition

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ChatInstalledAppTeamsAppDefinitionClient struct {
	Client *msgraph.Client
}

func NewChatInstalledAppTeamsAppDefinitionClientWithBaseURI(sdkApi sdkEnv.Api) (*ChatInstalledAppTeamsAppDefinitionClient, error) {
	client, err := msgraph.NewClient(sdkApi, "chatinstalledappteamsappdefinition", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating ChatInstalledAppTeamsAppDefinitionClient: %+v", err)
	}

	return &ChatInstalledAppTeamsAppDefinitionClient{
		Client: client,
	}, nil
}
