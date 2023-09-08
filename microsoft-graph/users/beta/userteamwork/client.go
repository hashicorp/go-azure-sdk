package userteamwork

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserTeamworkClient struct {
	Client *msgraph.Client
}

func NewUserTeamworkClientWithBaseURI(api sdkEnv.Api) (*UserTeamworkClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "userteamwork", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating UserTeamworkClient: %+v", err)
	}

	return &UserTeamworkClient{
		Client: client,
	}, nil
}
