package mejoinedteamprimarychannelsharedwithteamteam

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MeJoinedTeamPrimaryChannelSharedWithTeamTeamClient struct {
	Client *msgraph.Client
}

func NewMeJoinedTeamPrimaryChannelSharedWithTeamTeamClientWithBaseURI(api sdkEnv.Api) (*MeJoinedTeamPrimaryChannelSharedWithTeamTeamClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "mejoinedteamprimarychannelsharedwithteamteam", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating MeJoinedTeamPrimaryChannelSharedWithTeamTeamClient: %+v", err)
	}

	return &MeJoinedTeamPrimaryChannelSharedWithTeamTeamClient{
		Client: client,
	}, nil
}
