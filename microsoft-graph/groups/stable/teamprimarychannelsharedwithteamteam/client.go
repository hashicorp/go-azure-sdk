package teamprimarychannelsharedwithteamteam

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TeamPrimaryChannelSharedWithTeamTeamClient struct {
	Client *msgraph.Client
}

func NewTeamPrimaryChannelSharedWithTeamTeamClientWithBaseURI(sdkApi sdkEnv.Api) (*TeamPrimaryChannelSharedWithTeamTeamClient, error) {
	client, err := msgraph.NewClient(sdkApi, "teamprimarychannelsharedwithteamteam", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating TeamPrimaryChannelSharedWithTeamTeamClient: %+v", err)
	}

	return &TeamPrimaryChannelSharedWithTeamTeamClient{
		Client: client,
	}, nil
}
