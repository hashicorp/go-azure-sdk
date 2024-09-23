package teaminstalledappteamsappdefinition

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TeamInstalledAppTeamsAppDefinitionClient struct {
	Client *msgraph.Client
}

func NewTeamInstalledAppTeamsAppDefinitionClientWithBaseURI(sdkApi sdkEnv.Api) (*TeamInstalledAppTeamsAppDefinitionClient, error) {
	client, err := msgraph.NewClient(sdkApi, "teaminstalledappteamsappdefinition", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating TeamInstalledAppTeamsAppDefinitionClient: %+v", err)
	}

	return &TeamInstalledAppTeamsAppDefinitionClient{
		Client: client,
	}, nil
}
