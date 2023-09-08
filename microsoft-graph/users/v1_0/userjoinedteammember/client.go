package userjoinedteammember

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserJoinedTeamMemberClient struct {
	Client *msgraph.Client
}

func NewUserJoinedTeamMemberClientWithBaseURI(api sdkEnv.Api) (*UserJoinedTeamMemberClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "userjoinedteammember", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating UserJoinedTeamMemberClient: %+v", err)
	}

	return &UserJoinedTeamMemberClient{
		Client: client,
	}, nil
}
