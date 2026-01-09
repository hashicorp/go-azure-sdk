package teamwork

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TeamworkClient struct {
	Client *msgraph.Client
}

func NewTeamworkClientWithBaseURI(sdkApi sdkEnv.Api) (*TeamworkClient, error) {
	client, err := msgraph.NewClient(sdkApi, "teamwork", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating TeamworkClient: %+v", err)
	}

	return &TeamworkClient{
		Client: client,
	}, nil
}
