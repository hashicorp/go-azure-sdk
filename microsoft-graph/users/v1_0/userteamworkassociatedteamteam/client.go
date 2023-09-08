package userteamworkassociatedteamteam

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserTeamworkAssociatedTeamTeamClient struct {
	Client *msgraph.Client
}

func NewUserTeamworkAssociatedTeamTeamClientWithBaseURI(api sdkEnv.Api) (*UserTeamworkAssociatedTeamTeamClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "userteamworkassociatedteamteam", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating UserTeamworkAssociatedTeamTeamClient: %+v", err)
	}

	return &UserTeamworkAssociatedTeamTeamClient{
		Client: client,
	}, nil
}
