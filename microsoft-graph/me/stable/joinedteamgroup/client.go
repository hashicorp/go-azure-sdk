package joinedteamgroup

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type JoinedTeamGroupClient struct {
	Client *msgraph.Client
}

func NewJoinedTeamGroupClientWithBaseURI(sdkApi sdkEnv.Api) (*JoinedTeamGroupClient, error) {
	client, err := msgraph.NewClient(sdkApi, "joinedteamgroup", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating JoinedTeamGroupClient: %+v", err)
	}

	return &JoinedTeamGroupClient{
		Client: client,
	}, nil
}
