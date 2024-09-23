package driverootlastmodifiedbyuserserviceprovisioningerror

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DriveRootLastModifiedByUserServiceProvisioningErrorClient struct {
	Client *msgraph.Client
}

func NewDriveRootLastModifiedByUserServiceProvisioningErrorClientWithBaseURI(sdkApi sdkEnv.Api) (*DriveRootLastModifiedByUserServiceProvisioningErrorClient, error) {
	client, err := msgraph.NewClient(sdkApi, "driverootlastmodifiedbyuserserviceprovisioningerror", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating DriveRootLastModifiedByUserServiceProvisioningErrorClient: %+v", err)
	}

	return &DriveRootLastModifiedByUserServiceProvisioningErrorClient{
		Client: client,
	}, nil
}
