package userjoinedteamoperation

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserJoinedTeamOperationClient struct {
	Client *msgraph.Client
}

func NewUserJoinedTeamOperationClientWithBaseURI(api sdkEnv.Api) (*UserJoinedTeamOperationClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "userjoinedteamoperation", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating UserJoinedTeamOperationClient: %+v", err)
	}

	return &UserJoinedTeamOperationClient{
		Client: client,
	}, nil
}
