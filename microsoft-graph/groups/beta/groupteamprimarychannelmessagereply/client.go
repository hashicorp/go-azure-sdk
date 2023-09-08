package groupteamprimarychannelmessagereply

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GroupTeamPrimaryChannelMessageReplyClient struct {
	Client *msgraph.Client
}

func NewGroupTeamPrimaryChannelMessageReplyClientWithBaseURI(api sdkEnv.Api) (*GroupTeamPrimaryChannelMessageReplyClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "groupteamprimarychannelmessagereply", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating GroupTeamPrimaryChannelMessageReplyClient: %+v", err)
	}

	return &GroupTeamPrimaryChannelMessageReplyClient{
		Client: client,
	}, nil
}
