package driveitemanalyticsitemactivitystatactivitydriveitemcontent

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DriveItemAnalyticsItemActivityStatActivityDriveItemContentClient struct {
	Client *msgraph.Client
}

func NewDriveItemAnalyticsItemActivityStatActivityDriveItemContentClientWithBaseURI(sdkApi sdkEnv.Api) (*DriveItemAnalyticsItemActivityStatActivityDriveItemContentClient, error) {
	client, err := msgraph.NewClient(sdkApi, "driveitemanalyticsitemactivitystatactivitydriveitemcontent", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating DriveItemAnalyticsItemActivityStatActivityDriveItemContentClient: %+v", err)
	}

	return &DriveItemAnalyticsItemActivityStatActivityDriveItemContentClient{
		Client: client,
	}, nil
}
