package driverootlistitemdriveitemcontent

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DriveRootListItemDriveItemContentClient struct {
	Client *msgraph.Client
}

func NewDriveRootListItemDriveItemContentClientWithBaseURI(sdkApi sdkEnv.Api) (*DriveRootListItemDriveItemContentClient, error) {
	client, err := msgraph.NewClient(sdkApi, "driverootlistitemdriveitemcontent", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating DriveRootListItemDriveItemContentClient: %+v", err)
	}

	return &DriveRootListItemDriveItemContentClient{
		Client: client,
	}, nil
}
