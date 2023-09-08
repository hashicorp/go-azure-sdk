package meteamworkinstalledappteamsappdefinition

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MeTeamworkInstalledAppTeamsAppDefinitionClient struct {
	Client *msgraph.Client
}

func NewMeTeamworkInstalledAppTeamsAppDefinitionClientWithBaseURI(api sdkEnv.Api) (*MeTeamworkInstalledAppTeamsAppDefinitionClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "meteamworkinstalledappteamsappdefinition", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating MeTeamworkInstalledAppTeamsAppDefinitionClient: %+v", err)
	}

	return &MeTeamworkInstalledAppTeamsAppDefinitionClient{
		Client: client,
	}, nil
}
