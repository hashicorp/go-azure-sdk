package drivelistitemdriveitemcontentstream

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DriveListItemDriveItemContentStreamClient struct {
	Client *msgraph.Client
}

func NewDriveListItemDriveItemContentStreamClientWithBaseURI(sdkApi sdkEnv.Api) (*DriveListItemDriveItemContentStreamClient, error) {
	client, err := msgraph.NewClient(sdkApi, "drivelistitemdriveitemcontentstream", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating DriveListItemDriveItemContentStreamClient: %+v", err)
	}

	return &DriveListItemDriveItemContentStreamClient{
		Client: client,
	}, nil
}
