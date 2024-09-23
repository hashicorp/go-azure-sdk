package driveitemlistitemdriveitemcontent

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DriveItemListItemDriveItemContentClient struct {
	Client *msgraph.Client
}

func NewDriveItemListItemDriveItemContentClientWithBaseURI(sdkApi sdkEnv.Api) (*DriveItemListItemDriveItemContentClient, error) {
	client, err := msgraph.NewClient(sdkApi, "driveitemlistitemdriveitemcontent", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating DriveItemListItemDriveItemContentClient: %+v", err)
	}

	return &DriveItemListItemDriveItemContentClient{
		Client: client,
	}, nil
}
