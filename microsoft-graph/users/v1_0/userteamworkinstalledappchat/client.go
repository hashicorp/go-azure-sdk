package userteamworkinstalledappchat

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserTeamworkInstalledAppChatClient struct {
	Client *msgraph.Client
}

func NewUserTeamworkInstalledAppChatClientWithBaseURI(api sdkEnv.Api) (*UserTeamworkInstalledAppChatClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "userteamworkinstalledappchat", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating UserTeamworkInstalledAppChatClient: %+v", err)
	}

	return &UserTeamworkInstalledAppChatClient{
		Client: client,
	}, nil
}
