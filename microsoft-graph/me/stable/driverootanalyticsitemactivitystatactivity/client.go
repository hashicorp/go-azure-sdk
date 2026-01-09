package driverootanalyticsitemactivitystatactivity

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DriveRootAnalyticsItemActivityStatActivityClient struct {
	Client *msgraph.Client
}

func NewDriveRootAnalyticsItemActivityStatActivityClientWithBaseURI(sdkApi sdkEnv.Api) (*DriveRootAnalyticsItemActivityStatActivityClient, error) {
	client, err := msgraph.NewClient(sdkApi, "driverootanalyticsitemactivitystatactivity", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating DriveRootAnalyticsItemActivityStatActivityClient: %+v", err)
	}

	return &DriveRootAnalyticsItemActivityStatActivityClient{
		Client: client,
	}, nil
}
