package drivelistitemactivitydriveitemcontent

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DriveListItemActivityDriveItemContentClient struct {
	Client *msgraph.Client
}

func NewDriveListItemActivityDriveItemContentClientWithBaseURI(sdkApi sdkEnv.Api) (*DriveListItemActivityDriveItemContentClient, error) {
	client, err := msgraph.NewClient(sdkApi, "drivelistitemactivitydriveitemcontent", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating DriveListItemActivityDriveItemContentClient: %+v", err)
	}

	return &DriveListItemActivityDriveItemContentClient{
		Client: client,
	}, nil
}
