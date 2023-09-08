package meeventinstanceextension

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MeEventInstanceExtensionClient struct {
	Client *msgraph.Client
}

func NewMeEventInstanceExtensionClientWithBaseURI(api sdkEnv.Api) (*MeEventInstanceExtensionClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "meeventinstanceextension", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating MeEventInstanceExtensionClient: %+v", err)
	}

	return &MeEventInstanceExtensionClient{
		Client: client,
	}, nil
}
