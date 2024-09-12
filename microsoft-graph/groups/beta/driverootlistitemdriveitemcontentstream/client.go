package driverootlistitemdriveitemcontentstream

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DriveRootListItemDriveItemContentStreamClient struct {
	Client *msgraph.Client
}

func NewDriveRootListItemDriveItemContentStreamClientWithBaseURI(sdkApi sdkEnv.Api) (*DriveRootListItemDriveItemContentStreamClient, error) {
	client, err := msgraph.NewClient(sdkApi, "driverootlistitemdriveitemcontentstream", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating DriveRootListItemDriveItemContentStreamClient: %+v", err)
	}

	return &DriveRootListItemDriveItemContentStreamClient{
		Client: client,
	}, nil
}
