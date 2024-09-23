package driverootanalytics

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DriveRootAnalyticsClient struct {
	Client *msgraph.Client
}

func NewDriveRootAnalyticsClientWithBaseURI(sdkApi sdkEnv.Api) (*DriveRootAnalyticsClient, error) {
	client, err := msgraph.NewClient(sdkApi, "driverootanalytics", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating DriveRootAnalyticsClient: %+v", err)
	}

	return &DriveRootAnalyticsClient{
		Client: client,
	}, nil
}
