package joinedteaminstalledappteamsappdefinition

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type JoinedTeamInstalledAppTeamsAppDefinitionClient struct {
	Client *msgraph.Client
}

func NewJoinedTeamInstalledAppTeamsAppDefinitionClientWithBaseURI(sdkApi sdkEnv.Api) (*JoinedTeamInstalledAppTeamsAppDefinitionClient, error) {
	client, err := msgraph.NewClient(sdkApi, "joinedteaminstalledappteamsappdefinition", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating JoinedTeamInstalledAppTeamsAppDefinitionClient: %+v", err)
	}

	return &JoinedTeamInstalledAppTeamsAppDefinitionClient{
		Client: client,
	}, nil
}
