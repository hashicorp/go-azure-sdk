package meteamworkinstalledapp

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MeTeamworkInstalledAppClient struct {
	Client *msgraph.Client
}

func NewMeTeamworkInstalledAppClientWithBaseURI(api sdkEnv.Api) (*MeTeamworkInstalledAppClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "meteamworkinstalledapp", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating MeTeamworkInstalledAppClient: %+v", err)
	}

	return &MeTeamworkInstalledAppClient{
		Client: client,
	}, nil
}
