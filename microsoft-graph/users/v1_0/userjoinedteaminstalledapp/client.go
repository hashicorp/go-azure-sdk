package userjoinedteaminstalledapp

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserJoinedTeamInstalledAppClient struct {
	Client *msgraph.Client
}

func NewUserJoinedTeamInstalledAppClientWithBaseURI(api sdkEnv.Api) (*UserJoinedTeamInstalledAppClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "userjoinedteaminstalledapp", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating UserJoinedTeamInstalledAppClient: %+v", err)
	}

	return &UserJoinedTeamInstalledAppClient{
		Client: client,
	}, nil
}
