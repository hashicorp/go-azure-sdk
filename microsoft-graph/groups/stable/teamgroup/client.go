package teamgroup

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TeamGroupClient struct {
	Client *msgraph.Client
}

func NewTeamGroupClientWithBaseURI(sdkApi sdkEnv.Api) (*TeamGroupClient, error) {
	client, err := msgraph.NewClient(sdkApi, "teamgroup", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating TeamGroupClient: %+v", err)
	}

	return &TeamGroupClient{
		Client: client,
	}, nil
}
