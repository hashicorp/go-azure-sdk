package driveitempermission

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DriveItemPermissionClient struct {
	Client *msgraph.Client
}

func NewDriveItemPermissionClientWithBaseURI(sdkApi sdkEnv.Api) (*DriveItemPermissionClient, error) {
	client, err := msgraph.NewClient(sdkApi, "driveitempermission", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating DriveItemPermissionClient: %+v", err)
	}

	return &DriveItemPermissionClient{
		Client: client,
	}, nil
}
