package driverootlistitemactivitydriveitemcontent

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DriveRootListItemActivityDriveItemContentClient struct {
	Client *msgraph.Client
}

func NewDriveRootListItemActivityDriveItemContentClientWithBaseURI(sdkApi sdkEnv.Api) (*DriveRootListItemActivityDriveItemContentClient, error) {
	client, err := msgraph.NewClient(sdkApi, "driverootlistitemactivitydriveitemcontent", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating DriveRootListItemActivityDriveItemContentClient: %+v", err)
	}

	return &DriveRootListItemActivityDriveItemContentClient{
		Client: client,
	}, nil
}
