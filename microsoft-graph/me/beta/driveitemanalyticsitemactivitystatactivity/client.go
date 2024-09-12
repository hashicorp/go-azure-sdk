package driveitemanalyticsitemactivitystatactivity

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DriveItemAnalyticsItemActivityStatActivityClient struct {
	Client *msgraph.Client
}

func NewDriveItemAnalyticsItemActivityStatActivityClientWithBaseURI(sdkApi sdkEnv.Api) (*DriveItemAnalyticsItemActivityStatActivityClient, error) {
	client, err := msgraph.NewClient(sdkApi, "driveitemanalyticsitemactivitystatactivity", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating DriveItemAnalyticsItemActivityStatActivityClient: %+v", err)
	}

	return &DriveItemAnalyticsItemActivityStatActivityClient{
		Client: client,
	}, nil
}
