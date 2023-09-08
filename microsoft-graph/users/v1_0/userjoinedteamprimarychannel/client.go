package userjoinedteamprimarychannel

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserJoinedTeamPrimaryChannelClient struct {
	Client *msgraph.Client
}

func NewUserJoinedTeamPrimaryChannelClientWithBaseURI(api sdkEnv.Api) (*UserJoinedTeamPrimaryChannelClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "userjoinedteamprimarychannel", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating UserJoinedTeamPrimaryChannelClient: %+v", err)
	}

	return &UserJoinedTeamPrimaryChannelClient{
		Client: client,
	}, nil
}
