package userjoinedteam

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserJoinedTeamClient struct {
	Client *msgraph.Client
}

func NewUserJoinedTeamClientWithBaseURI(api sdkEnv.Api) (*UserJoinedTeamClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "userjoinedteam", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating UserJoinedTeamClient: %+v", err)
	}

	return &UserJoinedTeamClient{
		Client: client,
	}, nil
}
