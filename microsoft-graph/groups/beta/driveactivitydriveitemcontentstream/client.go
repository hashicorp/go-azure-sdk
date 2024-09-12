package driveactivitydriveitemcontentstream

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DriveActivityDriveItemContentStreamClient struct {
	Client *msgraph.Client
}

func NewDriveActivityDriveItemContentStreamClientWithBaseURI(sdkApi sdkEnv.Api) (*DriveActivityDriveItemContentStreamClient, error) {
	client, err := msgraph.NewClient(sdkApi, "driveactivitydriveitemcontentstream", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating DriveActivityDriveItemContentStreamClient: %+v", err)
	}

	return &DriveActivityDriveItemContentStreamClient{
		Client: client,
	}, nil
}
