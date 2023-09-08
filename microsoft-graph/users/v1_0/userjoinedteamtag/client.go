package userjoinedteamtag

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserJoinedTeamTagClient struct {
	Client *msgraph.Client
}

func NewUserJoinedTeamTagClientWithBaseURI(api sdkEnv.Api) (*UserJoinedTeamTagClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "userjoinedteamtag", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating UserJoinedTeamTagClient: %+v", err)
	}

	return &UserJoinedTeamTagClient{
		Client: client,
	}, nil
}
