package driveitemanalyticsitemactivitystatactivitydriveitemcontentstream

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DriveItemAnalyticsItemActivityStatActivityDriveItemContentStreamClient struct {
	Client *msgraph.Client
}

func NewDriveItemAnalyticsItemActivityStatActivityDriveItemContentStreamClientWithBaseURI(sdkApi sdkEnv.Api) (*DriveItemAnalyticsItemActivityStatActivityDriveItemContentStreamClient, error) {
	client, err := msgraph.NewClient(sdkApi, "driveitemanalyticsitemactivitystatactivitydriveitemcontentstream", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating DriveItemAnalyticsItemActivityStatActivityDriveItemContentStreamClient: %+v", err)
	}

	return &DriveItemAnalyticsItemActivityStatActivityDriveItemContentStreamClient{
		Client: client,
	}, nil
}
