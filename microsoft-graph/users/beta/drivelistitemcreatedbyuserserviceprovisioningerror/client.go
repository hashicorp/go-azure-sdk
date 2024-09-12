package drivelistitemcreatedbyuserserviceprovisioningerror

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DriveListItemCreatedByUserServiceProvisioningErrorClient struct {
	Client *msgraph.Client
}

func NewDriveListItemCreatedByUserServiceProvisioningErrorClientWithBaseURI(sdkApi sdkEnv.Api) (*DriveListItemCreatedByUserServiceProvisioningErrorClient, error) {
	client, err := msgraph.NewClient(sdkApi, "drivelistitemcreatedbyuserserviceprovisioningerror", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating DriveListItemCreatedByUserServiceProvisioningErrorClient: %+v", err)
	}

	return &DriveListItemCreatedByUserServiceProvisioningErrorClient{
		Client: client,
	}, nil
}
