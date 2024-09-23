package driveitemlistitemcreatedbyuserserviceprovisioningerror

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DriveItemListItemCreatedByUserServiceProvisioningErrorClient struct {
	Client *msgraph.Client
}

func NewDriveItemListItemCreatedByUserServiceProvisioningErrorClientWithBaseURI(sdkApi sdkEnv.Api) (*DriveItemListItemCreatedByUserServiceProvisioningErrorClient, error) {
	client, err := msgraph.NewClient(sdkApi, "driveitemlistitemcreatedbyuserserviceprovisioningerror", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating DriveItemListItemCreatedByUserServiceProvisioningErrorClient: %+v", err)
	}

	return &DriveItemListItemCreatedByUserServiceProvisioningErrorClient{
		Client: client,
	}, nil
}
