package drivelistitemactivitydriveitem

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DriveListItemActivityDriveItemClient struct {
	Client *msgraph.Client
}

func NewDriveListItemActivityDriveItemClientWithBaseURI(sdkApi sdkEnv.Api) (*DriveListItemActivityDriveItemClient, error) {
	client, err := msgraph.NewClient(sdkApi, "drivelistitemactivitydriveitem", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating DriveListItemActivityDriveItemClient: %+v", err)
	}

	return &DriveListItemActivityDriveItemClient{
		Client: client,
	}, nil
}
