package driveitemlastmodifiedbyuserserviceprovisioningerror

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DriveItemLastModifiedByUserServiceProvisioningErrorClient struct {
	Client *msgraph.Client
}

func NewDriveItemLastModifiedByUserServiceProvisioningErrorClientWithBaseURI(sdkApi sdkEnv.Api) (*DriveItemLastModifiedByUserServiceProvisioningErrorClient, error) {
	client, err := msgraph.NewClient(sdkApi, "driveitemlastmodifiedbyuserserviceprovisioningerror", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating DriveItemLastModifiedByUserServiceProvisioningErrorClient: %+v", err)
	}

	return &DriveItemLastModifiedByUserServiceProvisioningErrorClient{
		Client: client,
	}, nil
}
