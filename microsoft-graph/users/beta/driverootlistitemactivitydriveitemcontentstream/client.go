package driverootlistitemactivitydriveitemcontentstream

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DriveRootListItemActivityDriveItemContentStreamClient struct {
	Client *msgraph.Client
}

func NewDriveRootListItemActivityDriveItemContentStreamClientWithBaseURI(sdkApi sdkEnv.Api) (*DriveRootListItemActivityDriveItemContentStreamClient, error) {
	client, err := msgraph.NewClient(sdkApi, "driverootlistitemactivitydriveitemcontentstream", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating DriveRootListItemActivityDriveItemContentStreamClient: %+v", err)
	}

	return &DriveRootListItemActivityDriveItemContentStreamClient{
		Client: client,
	}, nil
}
