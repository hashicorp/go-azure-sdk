package driverootlistitemactivitydriveitem

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DriveRootListItemActivityDriveItemClient struct {
	Client *msgraph.Client
}

func NewDriveRootListItemActivityDriveItemClientWithBaseURI(sdkApi sdkEnv.Api) (*DriveRootListItemActivityDriveItemClient, error) {
	client, err := msgraph.NewClient(sdkApi, "driverootlistitemactivitydriveitem", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating DriveRootListItemActivityDriveItemClient: %+v", err)
	}

	return &DriveRootListItemActivityDriveItemClient{
		Client: client,
	}, nil
}
