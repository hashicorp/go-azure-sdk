package teamchannelmember

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TeamChannelMemberClient struct {
	Client *msgraph.Client
}

func NewTeamChannelMemberClientWithBaseURI(sdkApi sdkEnv.Api) (*TeamChannelMemberClient, error) {
	client, err := msgraph.NewClient(sdkApi, "teamchannelmember", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating TeamChannelMemberClient: %+v", err)
	}

	return &TeamChannelMemberClient{
		Client: client,
	}, nil
}
