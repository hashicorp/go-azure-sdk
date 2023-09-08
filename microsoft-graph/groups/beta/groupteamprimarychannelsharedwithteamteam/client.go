package groupteamprimarychannelsharedwithteamteam

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GroupTeamPrimaryChannelSharedWithTeamTeamClient struct {
	Client *msgraph.Client
}

func NewGroupTeamPrimaryChannelSharedWithTeamTeamClientWithBaseURI(api sdkEnv.Api) (*GroupTeamPrimaryChannelSharedWithTeamTeamClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "groupteamprimarychannelsharedwithteamteam", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating GroupTeamPrimaryChannelSharedWithTeamTeamClient: %+v", err)
	}

	return &GroupTeamPrimaryChannelSharedWithTeamTeamClient{
		Client: client,
	}, nil
}
