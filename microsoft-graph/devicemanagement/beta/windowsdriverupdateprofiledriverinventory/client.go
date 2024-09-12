package windowsdriverupdateprofiledriverinventory

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsDriverUpdateProfileDriverInventoryClient struct {
	Client *msgraph.Client
}

func NewWindowsDriverUpdateProfileDriverInventoryClientWithBaseURI(sdkApi sdkEnv.Api) (*WindowsDriverUpdateProfileDriverInventoryClient, error) {
	client, err := msgraph.NewClient(sdkApi, "windowsdriverupdateprofiledriverinventory", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating WindowsDriverUpdateProfileDriverInventoryClient: %+v", err)
	}

	return &WindowsDriverUpdateProfileDriverInventoryClient{
		Client: client,
	}, nil
}
