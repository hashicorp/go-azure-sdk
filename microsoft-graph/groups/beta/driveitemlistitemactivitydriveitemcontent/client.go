package driveitemlistitemactivitydriveitemcontent

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DriveItemListItemActivityDriveItemContentClient struct {
	Client *msgraph.Client
}

func NewDriveItemListItemActivityDriveItemContentClientWithBaseURI(sdkApi sdkEnv.Api) (*DriveItemListItemActivityDriveItemContentClient, error) {
	client, err := msgraph.NewClient(sdkApi, "driveitemlistitemactivitydriveitemcontent", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating DriveItemListItemActivityDriveItemContentClient: %+v", err)
	}

	return &DriveItemListItemActivityDriveItemContentClient{
		Client: client,
	}, nil
}
