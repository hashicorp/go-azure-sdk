package driverootanalyticslastsevenday

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DriveRootAnalyticsLastSevenDayClient struct {
	Client *msgraph.Client
}

func NewDriveRootAnalyticsLastSevenDayClientWithBaseURI(sdkApi sdkEnv.Api) (*DriveRootAnalyticsLastSevenDayClient, error) {
	client, err := msgraph.NewClient(sdkApi, "driverootanalyticslastsevenday", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating DriveRootAnalyticsLastSevenDayClient: %+v", err)
	}

	return &DriveRootAnalyticsLastSevenDayClient{
		Client: client,
	}, nil
}
