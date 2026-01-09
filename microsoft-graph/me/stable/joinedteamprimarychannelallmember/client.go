package joinedteamprimarychannelallmember

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type JoinedTeamPrimaryChannelAllMemberClient struct {
	Client *msgraph.Client
}

func NewJoinedTeamPrimaryChannelAllMemberClientWithBaseURI(sdkApi sdkEnv.Api) (*JoinedTeamPrimaryChannelAllMemberClient, error) {
	client, err := msgraph.NewClient(sdkApi, "joinedteamprimarychannelallmember", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating JoinedTeamPrimaryChannelAllMemberClient: %+v", err)
	}

	return &JoinedTeamPrimaryChannelAllMemberClient{
		Client: client,
	}, nil
}
