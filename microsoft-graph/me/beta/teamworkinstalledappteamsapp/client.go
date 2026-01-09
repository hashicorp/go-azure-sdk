package teamworkinstalledappteamsapp

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TeamworkInstalledAppTeamsAppClient struct {
	Client *msgraph.Client
}

func NewTeamworkInstalledAppTeamsAppClientWithBaseURI(sdkApi sdkEnv.Api) (*TeamworkInstalledAppTeamsAppClient, error) {
	client, err := msgraph.NewClient(sdkApi, "teamworkinstalledappteamsapp", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating TeamworkInstalledAppTeamsAppClient: %+v", err)
	}

	return &TeamworkInstalledAppTeamsAppClient{
		Client: client,
	}, nil
}
