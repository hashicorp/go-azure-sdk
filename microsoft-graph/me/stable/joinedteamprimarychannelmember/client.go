package joinedteamprimarychannelmember

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type JoinedTeamPrimaryChannelMemberClient struct {
	Client *msgraph.Client
}

func NewJoinedTeamPrimaryChannelMemberClientWithBaseURI(sdkApi sdkEnv.Api) (*JoinedTeamPrimaryChannelMemberClient, error) {
	client, err := msgraph.NewClient(sdkApi, "joinedteamprimarychannelmember", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating JoinedTeamPrimaryChannelMemberClient: %+v", err)
	}

	return &JoinedTeamPrimaryChannelMemberClient{
		Client: client,
	}, nil
}
