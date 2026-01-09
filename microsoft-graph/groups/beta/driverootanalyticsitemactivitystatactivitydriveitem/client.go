package driverootanalyticsitemactivitystatactivitydriveitem

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DriveRootAnalyticsItemActivityStatActivityDriveItemClient struct {
	Client *msgraph.Client
}

func NewDriveRootAnalyticsItemActivityStatActivityDriveItemClientWithBaseURI(sdkApi sdkEnv.Api) (*DriveRootAnalyticsItemActivityStatActivityDriveItemClient, error) {
	client, err := msgraph.NewClient(sdkApi, "driverootanalyticsitemactivitystatactivitydriveitem", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating DriveRootAnalyticsItemActivityStatActivityDriveItemClient: %+v", err)
	}

	return &DriveRootAnalyticsItemActivityStatActivityDriveItemClient{
		Client: client,
	}, nil
}
