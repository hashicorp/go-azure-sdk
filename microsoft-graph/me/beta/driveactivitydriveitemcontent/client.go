package driveactivitydriveitemcontent

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DriveActivityDriveItemContentClient struct {
	Client *msgraph.Client
}

func NewDriveActivityDriveItemContentClientWithBaseURI(sdkApi sdkEnv.Api) (*DriveActivityDriveItemContentClient, error) {
	client, err := msgraph.NewClient(sdkApi, "driveactivitydriveitemcontent", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating DriveActivityDriveItemContentClient: %+v", err)
	}

	return &DriveActivityDriveItemContentClient{
		Client: client,
	}, nil
}
