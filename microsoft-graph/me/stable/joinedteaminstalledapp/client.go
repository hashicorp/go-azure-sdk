package joinedteaminstalledapp

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type JoinedTeamInstalledAppClient struct {
	Client *msgraph.Client
}

func NewJoinedTeamInstalledAppClientWithBaseURI(sdkApi sdkEnv.Api) (*JoinedTeamInstalledAppClient, error) {
	client, err := msgraph.NewClient(sdkApi, "joinedteaminstalledapp", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating JoinedTeamInstalledAppClient: %+v", err)
	}

	return &JoinedTeamInstalledAppClient{
		Client: client,
	}, nil
}
