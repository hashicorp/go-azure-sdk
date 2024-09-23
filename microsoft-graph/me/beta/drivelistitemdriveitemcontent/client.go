package drivelistitemdriveitemcontent

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DriveListItemDriveItemContentClient struct {
	Client *msgraph.Client
}

func NewDriveListItemDriveItemContentClientWithBaseURI(sdkApi sdkEnv.Api) (*DriveListItemDriveItemContentClient, error) {
	client, err := msgraph.NewClient(sdkApi, "drivelistitemdriveitemcontent", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating DriveListItemDriveItemContentClient: %+v", err)
	}

	return &DriveListItemDriveItemContentClient{
		Client: client,
	}, nil
}
