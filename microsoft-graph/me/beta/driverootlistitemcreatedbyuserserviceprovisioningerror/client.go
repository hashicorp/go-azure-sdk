package driverootlistitemcreatedbyuserserviceprovisioningerror

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DriveRootListItemCreatedByUserServiceProvisioningErrorClient struct {
	Client *msgraph.Client
}

func NewDriveRootListItemCreatedByUserServiceProvisioningErrorClientWithBaseURI(sdkApi sdkEnv.Api) (*DriveRootListItemCreatedByUserServiceProvisioningErrorClient, error) {
	client, err := msgraph.NewClient(sdkApi, "driverootlistitemcreatedbyuserserviceprovisioningerror", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating DriveRootListItemCreatedByUserServiceProvisioningErrorClient: %+v", err)
	}

	return &DriveRootListItemCreatedByUserServiceProvisioningErrorClient{
		Client: client,
	}, nil
}
