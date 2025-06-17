package driveitemextension

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DriveItemExtensionClient struct {
	Client *msgraph.Client
}

func NewDriveItemExtensionClientWithBaseURI(sdkApi sdkEnv.Api) (*DriveItemExtensionClient, error) {
	client, err := msgraph.NewClient(sdkApi, "driveitemextension", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating DriveItemExtensionClient: %+v", err)
	}

	return &DriveItemExtensionClient{
		Client: client,
	}, nil
}
