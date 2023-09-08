package userjoinedteamprimarychannelmessagereplyhostedcontent

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserJoinedTeamPrimaryChannelMessageReplyHostedContentClient struct {
	Client *msgraph.Client
}

func NewUserJoinedTeamPrimaryChannelMessageReplyHostedContentClientWithBaseURI(api sdkEnv.Api) (*UserJoinedTeamPrimaryChannelMessageReplyHostedContentClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "userjoinedteamprimarychannelmessagereplyhostedcontent", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating UserJoinedTeamPrimaryChannelMessageReplyHostedContentClient: %+v", err)
	}

	return &UserJoinedTeamPrimaryChannelMessageReplyHostedContentClient{
		Client: client,
	}, nil
}
