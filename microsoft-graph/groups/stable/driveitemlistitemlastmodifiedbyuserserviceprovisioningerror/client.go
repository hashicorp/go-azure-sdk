package driveitemlistitemlastmodifiedbyuserserviceprovisioningerror

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DriveItemListItemLastModifiedByUserServiceProvisioningErrorClient struct {
	Client *msgraph.Client
}

func NewDriveItemListItemLastModifiedByUserServiceProvisioningErrorClientWithBaseURI(sdkApi sdkEnv.Api) (*DriveItemListItemLastModifiedByUserServiceProvisioningErrorClient, error) {
	client, err := msgraph.NewClient(sdkApi, "driveitemlistitemlastmodifiedbyuserserviceprovisioningerror", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating DriveItemListItemLastModifiedByUserServiceProvisioningErrorClient: %+v", err)
	}

	return &DriveItemListItemLastModifiedByUserServiceProvisioningErrorClient{
		Client: client,
	}, nil
}
