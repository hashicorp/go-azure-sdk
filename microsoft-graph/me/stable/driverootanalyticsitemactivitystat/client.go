package driverootanalyticsitemactivitystat

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DriveRootAnalyticsItemActivityStatClient struct {
	Client *msgraph.Client
}

func NewDriveRootAnalyticsItemActivityStatClientWithBaseURI(sdkApi sdkEnv.Api) (*DriveRootAnalyticsItemActivityStatClient, error) {
	client, err := msgraph.NewClient(sdkApi, "driverootanalyticsitemactivitystat", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating DriveRootAnalyticsItemActivityStatClient: %+v", err)
	}

	return &DriveRootAnalyticsItemActivityStatClient{
		Client: client,
	}, nil
}
