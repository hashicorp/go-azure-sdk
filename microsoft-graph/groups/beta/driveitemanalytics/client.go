package driveitemanalytics

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DriveItemAnalyticsClient struct {
	Client *msgraph.Client
}

func NewDriveItemAnalyticsClientWithBaseURI(sdkApi sdkEnv.Api) (*DriveItemAnalyticsClient, error) {
	client, err := msgraph.NewClient(sdkApi, "driveitemanalytics", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating DriveItemAnalyticsClient: %+v", err)
	}

	return &DriveItemAnalyticsClient{
		Client: client,
	}, nil
}
