package driverootextension

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DriveRootExtensionClient struct {
	Client *msgraph.Client
}

func NewDriveRootExtensionClientWithBaseURI(sdkApi sdkEnv.Api) (*DriveRootExtensionClient, error) {
	client, err := msgraph.NewClient(sdkApi, "driverootextension", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating DriveRootExtensionClient: %+v", err)
	}

	return &DriveRootExtensionClient{
		Client: client,
	}, nil
}
