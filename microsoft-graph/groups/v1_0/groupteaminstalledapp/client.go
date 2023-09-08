package groupteaminstalledapp

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GroupTeamInstalledAppClient struct {
	Client *msgraph.Client
}

func NewGroupTeamInstalledAppClientWithBaseURI(api sdkEnv.Api) (*GroupTeamInstalledAppClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "groupteaminstalledapp", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating GroupTeamInstalledAppClient: %+v", err)
	}

	return &GroupTeamInstalledAppClient{
		Client: client,
	}, nil
}
