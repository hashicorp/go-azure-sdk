package joinedteamprimarychannelsharedwithteamteam

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type JoinedTeamPrimaryChannelSharedWithTeamTeamClient struct {
	Client *msgraph.Client
}

func NewJoinedTeamPrimaryChannelSharedWithTeamTeamClientWithBaseURI(sdkApi sdkEnv.Api) (*JoinedTeamPrimaryChannelSharedWithTeamTeamClient, error) {
	client, err := msgraph.NewClient(sdkApi, "joinedteamprimarychannelsharedwithteamteam", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating JoinedTeamPrimaryChannelSharedWithTeamTeamClient: %+v", err)
	}

	return &JoinedTeamPrimaryChannelSharedWithTeamTeamClient{
		Client: client,
	}, nil
}
