package userjoinedteamprimarychannelmessagereply

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserJoinedTeamPrimaryChannelMessageReplyClient struct {
	Client *msgraph.Client
}

func NewUserJoinedTeamPrimaryChannelMessageReplyClientWithBaseURI(api sdkEnv.Api) (*UserJoinedTeamPrimaryChannelMessageReplyClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "userjoinedteamprimarychannelmessagereply", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating UserJoinedTeamPrimaryChannelMessageReplyClient: %+v", err)
	}

	return &UserJoinedTeamPrimaryChannelMessageReplyClient{
		Client: client,
	}, nil
}
