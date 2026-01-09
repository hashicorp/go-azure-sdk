package joinedteamtag

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type JoinedTeamTagClient struct {
	Client *msgraph.Client
}

func NewJoinedTeamTagClientWithBaseURI(sdkApi sdkEnv.Api) (*JoinedTeamTagClient, error) {
	client, err := msgraph.NewClient(sdkApi, "joinedteamtag", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating JoinedTeamTagClient: %+v", err)
	}

	return &JoinedTeamTagClient{
		Client: client,
	}, nil
}
