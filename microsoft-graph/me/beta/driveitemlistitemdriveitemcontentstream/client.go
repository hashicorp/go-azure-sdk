package driveitemlistitemdriveitemcontentstream

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DriveItemListItemDriveItemContentStreamClient struct {
	Client *msgraph.Client
}

func NewDriveItemListItemDriveItemContentStreamClientWithBaseURI(sdkApi sdkEnv.Api) (*DriveItemListItemDriveItemContentStreamClient, error) {
	client, err := msgraph.NewClient(sdkApi, "driveitemlistitemdriveitemcontentstream", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating DriveItemListItemDriveItemContentStreamClient: %+v", err)
	}

	return &DriveItemListItemDriveItemContentStreamClient{
		Client: client,
	}, nil
}
