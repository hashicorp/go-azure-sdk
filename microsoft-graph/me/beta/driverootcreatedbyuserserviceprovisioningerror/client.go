package driverootcreatedbyuserserviceprovisioningerror

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DriveRootCreatedByUserServiceProvisioningErrorClient struct {
	Client *msgraph.Client
}

func NewDriveRootCreatedByUserServiceProvisioningErrorClientWithBaseURI(sdkApi sdkEnv.Api) (*DriveRootCreatedByUserServiceProvisioningErrorClient, error) {
	client, err := msgraph.NewClient(sdkApi, "driverootcreatedbyuserserviceprovisioningerror", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating DriveRootCreatedByUserServiceProvisioningErrorClient: %+v", err)
	}

	return &DriveRootCreatedByUserServiceProvisioningErrorClient{
		Client: client,
	}, nil
}
