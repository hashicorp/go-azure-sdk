package driveitemlistitemactivitycontent

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DriveItemListItemActivityContentClient struct {
	Client *msgraph.Client
}

func NewDriveItemListItemActivityContentClientWithBaseURI(sdkApi sdkEnv.Api) (*DriveItemListItemActivityContentClient, error) {
	client, err := msgraph.NewClient(sdkApi, "driveitemlistitemactivitycontent", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating DriveItemListItemActivityContentClient: %+v", err)
	}

	return &DriveItemListItemActivityContentClient{
		Client: client,
	}, nil
}
