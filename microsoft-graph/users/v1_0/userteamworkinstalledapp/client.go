package userteamworkinstalledapp

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserTeamworkInstalledAppClient struct {
	Client *msgraph.Client
}

func NewUserTeamworkInstalledAppClientWithBaseURI(api sdkEnv.Api) (*UserTeamworkInstalledAppClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "userteamworkinstalledapp", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating UserTeamworkInstalledAppClient: %+v", err)
	}

	return &UserTeamworkInstalledAppClient{
		Client: client,
	}, nil
}
