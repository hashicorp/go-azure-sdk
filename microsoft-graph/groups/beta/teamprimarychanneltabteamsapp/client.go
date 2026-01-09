package teamprimarychanneltabteamsapp

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TeamPrimaryChannelTabTeamsAppClient struct {
	Client *msgraph.Client
}

func NewTeamPrimaryChannelTabTeamsAppClientWithBaseURI(sdkApi sdkEnv.Api) (*TeamPrimaryChannelTabTeamsAppClient, error) {
	client, err := msgraph.NewClient(sdkApi, "teamprimarychanneltabteamsapp", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating TeamPrimaryChannelTabTeamsAppClient: %+v", err)
	}

	return &TeamPrimaryChannelTabTeamsAppClient{
		Client: client,
	}, nil
}
