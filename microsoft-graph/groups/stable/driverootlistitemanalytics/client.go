package driverootlistitemanalytics

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DriveRootListItemAnalyticsClient struct {
	Client *msgraph.Client
}

func NewDriveRootListItemAnalyticsClientWithBaseURI(sdkApi sdkEnv.Api) (*DriveRootListItemAnalyticsClient, error) {
	client, err := msgraph.NewClient(sdkApi, "driverootlistitemanalytics", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating DriveRootListItemAnalyticsClient: %+v", err)
	}

	return &DriveRootListItemAnalyticsClient{
		Client: client,
	}, nil
}
