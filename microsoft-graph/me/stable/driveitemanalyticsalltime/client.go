package driveitemanalyticsalltime

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DriveItemAnalyticsAllTimeClient struct {
	Client *msgraph.Client
}

func NewDriveItemAnalyticsAllTimeClientWithBaseURI(sdkApi sdkEnv.Api) (*DriveItemAnalyticsAllTimeClient, error) {
	client, err := msgraph.NewClient(sdkApi, "driveitemanalyticsalltime", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating DriveItemAnalyticsAllTimeClient: %+v", err)
	}

	return &DriveItemAnalyticsAllTimeClient{
		Client: client,
	}, nil
}
