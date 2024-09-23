package driveitemlistitempermission

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DriveItemListItemPermissionClient struct {
	Client *msgraph.Client
}

func NewDriveItemListItemPermissionClientWithBaseURI(sdkApi sdkEnv.Api) (*DriveItemListItemPermissionClient, error) {
	client, err := msgraph.NewClient(sdkApi, "driveitemlistitempermission", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating DriveItemListItemPermissionClient: %+v", err)
	}

	return &DriveItemListItemPermissionClient{
		Client: client,
	}, nil
}
