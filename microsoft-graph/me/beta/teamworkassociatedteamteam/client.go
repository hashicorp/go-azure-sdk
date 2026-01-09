package teamworkassociatedteamteam

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TeamworkAssociatedTeamTeamClient struct {
	Client *msgraph.Client
}

func NewTeamworkAssociatedTeamTeamClientWithBaseURI(sdkApi sdkEnv.Api) (*TeamworkAssociatedTeamTeamClient, error) {
	client, err := msgraph.NewClient(sdkApi, "teamworkassociatedteamteam", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating TeamworkAssociatedTeamTeamClient: %+v", err)
	}

	return &TeamworkAssociatedTeamTeamClient{
		Client: client,
	}, nil
}
