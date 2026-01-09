package drivelistcreatedbyuserserviceprovisioningerror

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DriveListCreatedByUserServiceProvisioningErrorClient struct {
	Client *msgraph.Client
}

func NewDriveListCreatedByUserServiceProvisioningErrorClientWithBaseURI(sdkApi sdkEnv.Api) (*DriveListCreatedByUserServiceProvisioningErrorClient, error) {
	client, err := msgraph.NewClient(sdkApi, "drivelistcreatedbyuserserviceprovisioningerror", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating DriveListCreatedByUserServiceProvisioningErrorClient: %+v", err)
	}

	return &DriveListCreatedByUserServiceProvisioningErrorClient{
		Client: client,
	}, nil
}
