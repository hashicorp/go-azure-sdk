package driverootlistitemlastmodifiedbyuserserviceprovisioningerror

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DriveRootListItemLastModifiedByUserServiceProvisioningErrorClient struct {
	Client *msgraph.Client
}

func NewDriveRootListItemLastModifiedByUserServiceProvisioningErrorClientWithBaseURI(sdkApi sdkEnv.Api) (*DriveRootListItemLastModifiedByUserServiceProvisioningErrorClient, error) {
	client, err := msgraph.NewClient(sdkApi, "driverootlistitemlastmodifiedbyuserserviceprovisioningerror", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating DriveRootListItemLastModifiedByUserServiceProvisioningErrorClient: %+v", err)
	}

	return &DriveRootListItemLastModifiedByUserServiceProvisioningErrorClient{
		Client: client,
	}, nil
}
