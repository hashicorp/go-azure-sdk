package userjoinedteaminstalledappteamsappdefinition

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserJoinedTeamInstalledAppTeamsAppDefinitionClient struct {
	Client *msgraph.Client
}

func NewUserJoinedTeamInstalledAppTeamsAppDefinitionClientWithBaseURI(api sdkEnv.Api) (*UserJoinedTeamInstalledAppTeamsAppDefinitionClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "userjoinedteaminstalledappteamsappdefinition", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating UserJoinedTeamInstalledAppTeamsAppDefinitionClient: %+v", err)
	}

	return &UserJoinedTeamInstalledAppTeamsAppDefinitionClient{
		Client: client,
	}, nil
}
