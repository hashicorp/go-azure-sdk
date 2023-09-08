package userjoinedteamprimarychannelmessagehostedcontent

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserJoinedTeamPrimaryChannelMessageHostedContentClient struct {
	Client *msgraph.Client
}

func NewUserJoinedTeamPrimaryChannelMessageHostedContentClientWithBaseURI(api sdkEnv.Api) (*UserJoinedTeamPrimaryChannelMessageHostedContentClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "userjoinedteamprimarychannelmessagehostedcontent", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating UserJoinedTeamPrimaryChannelMessageHostedContentClient: %+v", err)
	}

	return &UserJoinedTeamPrimaryChannelMessageHostedContentClient{
		Client: client,
	}, nil
}
