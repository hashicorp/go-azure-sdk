package driveitemthumbnail

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DriveItemThumbnailClient struct {
	Client *msgraph.Client
}

func NewDriveItemThumbnailClientWithBaseURI(sdkApi sdkEnv.Api) (*DriveItemThumbnailClient, error) {
	client, err := msgraph.NewClient(sdkApi, "driveitemthumbnail", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating DriveItemThumbnailClient: %+v", err)
	}

	return &DriveItemThumbnailClient{
		Client: client,
	}, nil
}
