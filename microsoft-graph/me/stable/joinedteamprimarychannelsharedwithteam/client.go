package joinedteamprimarychannelsharedwithteam

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type JoinedTeamPrimaryChannelSharedWithTeamClient struct {
	Client *msgraph.Client
}

func NewJoinedTeamPrimaryChannelSharedWithTeamClientWithBaseURI(sdkApi sdkEnv.Api) (*JoinedTeamPrimaryChannelSharedWithTeamClient, error) {
	client, err := msgraph.NewClient(sdkApi, "joinedteamprimarychannelsharedwithteam", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating JoinedTeamPrimaryChannelSharedWithTeamClient: %+v", err)
	}

	return &JoinedTeamPrimaryChannelSharedWithTeamClient{
		Client: client,
	}, nil
}
