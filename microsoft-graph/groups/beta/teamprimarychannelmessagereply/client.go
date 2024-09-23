package teamprimarychannelmessagereply

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TeamPrimaryChannelMessageReplyClient struct {
	Client *msgraph.Client
}

func NewTeamPrimaryChannelMessageReplyClientWithBaseURI(sdkApi sdkEnv.Api) (*TeamPrimaryChannelMessageReplyClient, error) {
	client, err := msgraph.NewClient(sdkApi, "teamprimarychannelmessagereply", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating TeamPrimaryChannelMessageReplyClient: %+v", err)
	}

	return &TeamPrimaryChannelMessageReplyClient{
		Client: client,
	}, nil
}
