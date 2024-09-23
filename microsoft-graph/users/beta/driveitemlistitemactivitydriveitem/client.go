package driveitemlistitemactivitydriveitem

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DriveItemListItemActivityDriveItemClient struct {
	Client *msgraph.Client
}

func NewDriveItemListItemActivityDriveItemClientWithBaseURI(sdkApi sdkEnv.Api) (*DriveItemListItemActivityDriveItemClient, error) {
	client, err := msgraph.NewClient(sdkApi, "driveitemlistitemactivitydriveitem", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating DriveItemListItemActivityDriveItemClient: %+v", err)
	}

	return &DriveItemListItemActivityDriveItemClient{
		Client: client,
	}, nil
}
