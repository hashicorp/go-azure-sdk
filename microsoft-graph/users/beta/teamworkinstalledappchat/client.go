package teamworkinstalledappchat

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TeamworkInstalledAppChatClient struct {
	Client *msgraph.Client
}

func NewTeamworkInstalledAppChatClientWithBaseURI(sdkApi sdkEnv.Api) (*TeamworkInstalledAppChatClient, error) {
	client, err := msgraph.NewClient(sdkApi, "teamworkinstalledappchat", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating TeamworkInstalledAppChatClient: %+v", err)
	}

	return &TeamworkInstalledAppChatClient{
		Client: client,
	}, nil
}
