package driverootlistitemdriveitem

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DriveRootListItemDriveItemClient struct {
	Client *msgraph.Client
}

func NewDriveRootListItemDriveItemClientWithBaseURI(sdkApi sdkEnv.Api) (*DriveRootListItemDriveItemClient, error) {
	client, err := msgraph.NewClient(sdkApi, "driverootlistitemdriveitem", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating DriveRootListItemDriveItemClient: %+v", err)
	}

	return &DriveRootListItemDriveItemClient{
		Client: client,
	}, nil
}
