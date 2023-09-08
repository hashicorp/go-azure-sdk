package mecontactextension

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MeContactExtensionClient struct {
	Client *msgraph.Client
}

func NewMeContactExtensionClientWithBaseURI(api sdkEnv.Api) (*MeContactExtensionClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "mecontactextension", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating MeContactExtensionClient: %+v", err)
	}

	return &MeContactExtensionClient{
		Client: client,
	}, nil
}
