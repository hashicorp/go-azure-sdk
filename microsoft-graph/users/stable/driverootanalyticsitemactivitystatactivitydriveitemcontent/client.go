package driverootanalyticsitemactivitystatactivitydriveitemcontent

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DriveRootAnalyticsItemActivityStatActivityDriveItemContentClient struct {
	Client *msgraph.Client
}

func NewDriveRootAnalyticsItemActivityStatActivityDriveItemContentClientWithBaseURI(sdkApi sdkEnv.Api) (*DriveRootAnalyticsItemActivityStatActivityDriveItemContentClient, error) {
	client, err := msgraph.NewClient(sdkApi, "driverootanalyticsitemactivitystatactivitydriveitemcontent", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating DriveRootAnalyticsItemActivityStatActivityDriveItemContentClient: %+v", err)
	}

	return &DriveRootAnalyticsItemActivityStatActivityDriveItemContentClient{
		Client: client,
	}, nil
}
