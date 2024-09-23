package driveitemlistitemdriveitem

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DriveItemListItemDriveItemClient struct {
	Client *msgraph.Client
}

func NewDriveItemListItemDriveItemClientWithBaseURI(sdkApi sdkEnv.Api) (*DriveItemListItemDriveItemClient, error) {
	client, err := msgraph.NewClient(sdkApi, "driveitemlistitemdriveitem", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating DriveItemListItemDriveItemClient: %+v", err)
	}

	return &DriveItemListItemDriveItemClient{
		Client: client,
	}, nil
}
