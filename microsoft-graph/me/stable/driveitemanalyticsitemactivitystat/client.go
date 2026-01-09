package driveitemanalyticsitemactivitystat

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DriveItemAnalyticsItemActivityStatClient struct {
	Client *msgraph.Client
}

func NewDriveItemAnalyticsItemActivityStatClientWithBaseURI(sdkApi sdkEnv.Api) (*DriveItemAnalyticsItemActivityStatClient, error) {
	client, err := msgraph.NewClient(sdkApi, "driveitemanalyticsitemactivitystat", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating DriveItemAnalyticsItemActivityStatClient: %+v", err)
	}

	return &DriveItemAnalyticsItemActivityStatClient{
		Client: client,
	}, nil
}
