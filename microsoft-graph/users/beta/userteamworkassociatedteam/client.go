package userteamworkassociatedteam

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserTeamworkAssociatedTeamClient struct {
	Client *msgraph.Client
}

func NewUserTeamworkAssociatedTeamClientWithBaseURI(api sdkEnv.Api) (*UserTeamworkAssociatedTeamClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "userteamworkassociatedteam", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating UserTeamworkAssociatedTeamClient: %+v", err)
	}

	return &UserTeamworkAssociatedTeamClient{
		Client: client,
	}, nil
}
