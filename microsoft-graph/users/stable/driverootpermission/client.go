package driverootpermission

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DriveRootPermissionClient struct {
	Client *msgraph.Client
}

func NewDriveRootPermissionClientWithBaseURI(sdkApi sdkEnv.Api) (*DriveRootPermissionClient, error) {
	client, err := msgraph.NewClient(sdkApi, "driverootpermission", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating DriveRootPermissionClient: %+v", err)
	}

	return &DriveRootPermissionClient{
		Client: client,
	}, nil
}
