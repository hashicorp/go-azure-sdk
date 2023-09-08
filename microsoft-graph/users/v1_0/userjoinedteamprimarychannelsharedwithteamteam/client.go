package userjoinedteamprimarychannelsharedwithteamteam

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserJoinedTeamPrimaryChannelSharedWithTeamTeamClient struct {
	Client *msgraph.Client
}

func NewUserJoinedTeamPrimaryChannelSharedWithTeamTeamClientWithBaseURI(api sdkEnv.Api) (*UserJoinedTeamPrimaryChannelSharedWithTeamTeamClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "userjoinedteamprimarychannelsharedwithteamteam", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating UserJoinedTeamPrimaryChannelSharedWithTeamTeamClient: %+v", err)
	}

	return &UserJoinedTeamPrimaryChannelSharedWithTeamTeamClient{
		Client: client,
	}, nil
}
