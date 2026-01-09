package teamchannelsharedwithteam

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TeamChannelSharedWithTeamClient struct {
	Client *msgraph.Client
}

func NewTeamChannelSharedWithTeamClientWithBaseURI(sdkApi sdkEnv.Api) (*TeamChannelSharedWithTeamClient, error) {
	client, err := msgraph.NewClient(sdkApi, "teamchannelsharedwithteam", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating TeamChannelSharedWithTeamClient: %+v", err)
	}

	return &TeamChannelSharedWithTeamClient{
		Client: client,
	}, nil
}
