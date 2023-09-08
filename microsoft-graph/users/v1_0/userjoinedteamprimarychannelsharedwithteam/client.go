package userjoinedteamprimarychannelsharedwithteam

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserJoinedTeamPrimaryChannelSharedWithTeamClient struct {
	Client *msgraph.Client
}

func NewUserJoinedTeamPrimaryChannelSharedWithTeamClientWithBaseURI(api sdkEnv.Api) (*UserJoinedTeamPrimaryChannelSharedWithTeamClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "userjoinedteamprimarychannelsharedwithteam", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating UserJoinedTeamPrimaryChannelSharedWithTeamClient: %+v", err)
	}

	return &UserJoinedTeamPrimaryChannelSharedWithTeamClient{
		Client: client,
	}, nil
}
