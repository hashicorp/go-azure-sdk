package teamprimarychannelmessagereplyhostedcontent

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TeamPrimaryChannelMessageReplyHostedContentClient struct {
	Client *msgraph.Client
}

func NewTeamPrimaryChannelMessageReplyHostedContentClientWithBaseURI(sdkApi sdkEnv.Api) (*TeamPrimaryChannelMessageReplyHostedContentClient, error) {
	client, err := msgraph.NewClient(sdkApi, "teamprimarychannelmessagereplyhostedcontent", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating TeamPrimaryChannelMessageReplyHostedContentClient: %+v", err)
	}

	return &TeamPrimaryChannelMessageReplyHostedContentClient{
		Client: client,
	}, nil
}
