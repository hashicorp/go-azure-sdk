package driveitemlistitemactivitydriveitemcontentstream

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DriveItemListItemActivityDriveItemContentStreamClient struct {
	Client *msgraph.Client
}

func NewDriveItemListItemActivityDriveItemContentStreamClientWithBaseURI(sdkApi sdkEnv.Api) (*DriveItemListItemActivityDriveItemContentStreamClient, error) {
	client, err := msgraph.NewClient(sdkApi, "driveitemlistitemactivitydriveitemcontentstream", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating DriveItemListItemActivityDriveItemContentStreamClient: %+v", err)
	}

	return &DriveItemListItemActivityDriveItemContentStreamClient{
		Client: client,
	}, nil
}
