package teaminstalledapp

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TeamInstalledAppClient struct {
	Client *msgraph.Client
}

func NewTeamInstalledAppClientWithBaseURI(sdkApi sdkEnv.Api) (*TeamInstalledAppClient, error) {
	client, err := msgraph.NewClient(sdkApi, "teaminstalledapp", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating TeamInstalledAppClient: %+v", err)
	}

	return &TeamInstalledAppClient{
		Client: client,
	}, nil
}
