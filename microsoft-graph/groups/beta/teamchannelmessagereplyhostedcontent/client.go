package teamchannelmessagereplyhostedcontent

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TeamChannelMessageReplyHostedContentClient struct {
	Client *msgraph.Client
}

func NewTeamChannelMessageReplyHostedContentClientWithBaseURI(sdkApi sdkEnv.Api) (*TeamChannelMessageReplyHostedContentClient, error) {
	client, err := msgraph.NewClient(sdkApi, "teamchannelmessagereplyhostedcontent", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating TeamChannelMessageReplyHostedContentClient: %+v", err)
	}

	return &TeamChannelMessageReplyHostedContentClient{
		Client: client,
	}, nil
}
