package userjoinedteamprimarychanneltab

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserJoinedTeamPrimaryChannelTabClient struct {
	Client *msgraph.Client
}

func NewUserJoinedTeamPrimaryChannelTabClientWithBaseURI(api sdkEnv.Api) (*UserJoinedTeamPrimaryChannelTabClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "userjoinedteamprimarychanneltab", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating UserJoinedTeamPrimaryChannelTabClient: %+v", err)
	}

	return &UserJoinedTeamPrimaryChannelTabClient{
		Client: client,
	}, nil
}
