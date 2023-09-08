package meteamworkinstalledappteamsapp

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MeTeamworkInstalledAppTeamsAppClient struct {
	Client *msgraph.Client
}

func NewMeTeamworkInstalledAppTeamsAppClientWithBaseURI(api sdkEnv.Api) (*MeTeamworkInstalledAppTeamsAppClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "meteamworkinstalledappteamsapp", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating MeTeamworkInstalledAppTeamsAppClient: %+v", err)
	}

	return &MeTeamworkInstalledAppTeamsAppClient{
		Client: client,
	}, nil
}
