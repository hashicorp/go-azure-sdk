package driveitemanalyticslastsevenday

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DriveItemAnalyticsLastSevenDayClient struct {
	Client *msgraph.Client
}

func NewDriveItemAnalyticsLastSevenDayClientWithBaseURI(sdkApi sdkEnv.Api) (*DriveItemAnalyticsLastSevenDayClient, error) {
	client, err := msgraph.NewClient(sdkApi, "driveitemanalyticslastsevenday", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating DriveItemAnalyticsLastSevenDayClient: %+v", err)
	}

	return &DriveItemAnalyticsLastSevenDayClient{
		Client: client,
	}, nil
}
