package meprofileposition

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MeProfilePositionClient struct {
	Client *msgraph.Client
}

func NewMeProfilePositionClientWithBaseURI(api sdkEnv.Api) (*MeProfilePositionClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "meprofileposition", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating MeProfilePositionClient: %+v", err)
	}

	return &MeProfilePositionClient{
		Client: client,
	}, nil
}
