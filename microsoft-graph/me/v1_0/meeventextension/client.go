package meeventextension

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MeEventExtensionClient struct {
	Client *msgraph.Client
}

func NewMeEventExtensionClientWithBaseURI(api sdkEnv.Api) (*MeEventExtensionClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "meeventextension", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating MeEventExtensionClient: %+v", err)
	}

	return &MeEventExtensionClient{
		Client: client,
	}, nil
}
