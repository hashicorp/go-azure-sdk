package joinedteamprimarychannelmessagehostedcontent

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type JoinedTeamPrimaryChannelMessageHostedContentClient struct {
	Client *msgraph.Client
}

func NewJoinedTeamPrimaryChannelMessageHostedContentClientWithBaseURI(sdkApi sdkEnv.Api) (*JoinedTeamPrimaryChannelMessageHostedContentClient, error) {
	client, err := msgraph.NewClient(sdkApi, "joinedteamprimarychannelmessagehostedcontent", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating JoinedTeamPrimaryChannelMessageHostedContentClient: %+v", err)
	}

	return &JoinedTeamPrimaryChannelMessageHostedContentClient{
		Client: client,
	}, nil
}
