package profileposition

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ProfilePositionClient struct {
	Client *msgraph.Client
}

func NewProfilePositionClientWithBaseURI(sdkApi sdkEnv.Api) (*ProfilePositionClient, error) {
	client, err := msgraph.NewClient(sdkApi, "profileposition", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating ProfilePositionClient: %+v", err)
	}

	return &ProfilePositionClient{
		Client: client,
	}, nil
}
