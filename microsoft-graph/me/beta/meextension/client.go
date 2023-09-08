package meextension

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MeExtensionClient struct {
	Client *msgraph.Client
}

func NewMeExtensionClientWithBaseURI(api sdkEnv.Api) (*MeExtensionClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "meextension", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating MeExtensionClient: %+v", err)
	}

	return &MeExtensionClient{
		Client: client,
	}, nil
}
