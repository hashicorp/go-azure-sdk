package memessageextension

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MeMessageExtensionClient struct {
	Client *msgraph.Client
}

func NewMeMessageExtensionClientWithBaseURI(api sdkEnv.Api) (*MeMessageExtensionClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "memessageextension", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating MeMessageExtensionClient: %+v", err)
	}

	return &MeMessageExtensionClient{
		Client: client,
	}, nil
}
