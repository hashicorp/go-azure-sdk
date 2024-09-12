package driveitemanalyticsitemactivitystatactivitydriveitem

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DriveItemAnalyticsItemActivityStatActivityDriveItemClient struct {
	Client *msgraph.Client
}

func NewDriveItemAnalyticsItemActivityStatActivityDriveItemClientWithBaseURI(sdkApi sdkEnv.Api) (*DriveItemAnalyticsItemActivityStatActivityDriveItemClient, error) {
	client, err := msgraph.NewClient(sdkApi, "driveitemanalyticsitemactivitystatactivitydriveitem", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating DriveItemAnalyticsItemActivityStatActivityDriveItemClient: %+v", err)
	}

	return &DriveItemAnalyticsItemActivityStatActivityDriveItemClient{
		Client: client,
	}, nil
}
