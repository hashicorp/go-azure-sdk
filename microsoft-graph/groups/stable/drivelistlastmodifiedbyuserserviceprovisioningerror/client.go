package drivelistlastmodifiedbyuserserviceprovisioningerror

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DriveListLastModifiedByUserServiceProvisioningErrorClient struct {
	Client *msgraph.Client
}

func NewDriveListLastModifiedByUserServiceProvisioningErrorClientWithBaseURI(sdkApi sdkEnv.Api) (*DriveListLastModifiedByUserServiceProvisioningErrorClient, error) {
	client, err := msgraph.NewClient(sdkApi, "drivelistlastmodifiedbyuserserviceprovisioningerror", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating DriveListLastModifiedByUserServiceProvisioningErrorClient: %+v", err)
	}

	return &DriveListLastModifiedByUserServiceProvisioningErrorClient{
		Client: client,
	}, nil
}
