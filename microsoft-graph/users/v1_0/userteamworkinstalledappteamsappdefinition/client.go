package userteamworkinstalledappteamsappdefinition

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserTeamworkInstalledAppTeamsAppDefinitionClient struct {
	Client *msgraph.Client
}

func NewUserTeamworkInstalledAppTeamsAppDefinitionClientWithBaseURI(api sdkEnv.Api) (*UserTeamworkInstalledAppTeamsAppDefinitionClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "userteamworkinstalledappteamsappdefinition", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating UserTeamworkInstalledAppTeamsAppDefinitionClient: %+v", err)
	}

	return &UserTeamworkInstalledAppTeamsAppDefinitionClient{
		Client: client,
	}, nil
}
