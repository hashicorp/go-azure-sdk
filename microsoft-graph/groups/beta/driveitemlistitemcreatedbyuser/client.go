package driveitemlistitemcreatedbyuser

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DriveItemListItemCreatedByUserClient struct {
	Client *msgraph.Client
}

func NewDriveItemListItemCreatedByUserClientWithBaseURI(sdkApi sdkEnv.Api) (*DriveItemListItemCreatedByUserClient, error) {
	client, err := msgraph.NewClient(sdkApi, "driveitemlistitemcreatedbyuser", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating DriveItemListItemCreatedByUserClient: %+v", err)
	}

	return &DriveItemListItemCreatedByUserClient{
		Client: client,
	}, nil
}
