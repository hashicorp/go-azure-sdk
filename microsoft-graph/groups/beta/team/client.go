package team

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TeamClient struct {
	Client *msgraph.Client
}

func NewTeamClientWithBaseURI(sdkApi sdkEnv.Api) (*TeamClient, error) {
	client, err := msgraph.NewClient(sdkApi, "team", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating TeamClient: %+v", err)
	}

	return &TeamClient{
		Client: client,
	}, nil
}
