package driveactivitydriveitem

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DriveActivityDriveItemClient struct {
	Client *msgraph.Client
}

func NewDriveActivityDriveItemClientWithBaseURI(sdkApi sdkEnv.Api) (*DriveActivityDriveItemClient, error) {
	client, err := msgraph.NewClient(sdkApi, "driveactivitydriveitem", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating DriveActivityDriveItemClient: %+v", err)
	}

	return &DriveActivityDriveItemClient{
		Client: client,
	}, nil
}
