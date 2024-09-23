package drivelistitempermission

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DriveListItemPermissionClient struct {
	Client *msgraph.Client
}

func NewDriveListItemPermissionClientWithBaseURI(sdkApi sdkEnv.Api) (*DriveListItemPermissionClient, error) {
	client, err := msgraph.NewClient(sdkApi, "drivelistitempermission", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating DriveListItemPermissionClient: %+v", err)
	}

	return &DriveListItemPermissionClient{
		Client: client,
	}, nil
}
