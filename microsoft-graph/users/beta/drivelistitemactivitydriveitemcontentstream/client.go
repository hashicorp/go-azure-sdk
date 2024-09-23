package drivelistitemactivitydriveitemcontentstream

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DriveListItemActivityDriveItemContentStreamClient struct {
	Client *msgraph.Client
}

func NewDriveListItemActivityDriveItemContentStreamClientWithBaseURI(sdkApi sdkEnv.Api) (*DriveListItemActivityDriveItemContentStreamClient, error) {
	client, err := msgraph.NewClient(sdkApi, "drivelistitemactivitydriveitemcontentstream", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating DriveListItemActivityDriveItemContentStreamClient: %+v", err)
	}

	return &DriveListItemActivityDriveItemContentStreamClient{
		Client: client,
	}, nil
}
