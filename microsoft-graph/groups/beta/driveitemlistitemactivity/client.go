package driveitemlistitemactivity

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DriveItemListItemActivityClient struct {
	Client *msgraph.Client
}

func NewDriveItemListItemActivityClientWithBaseURI(sdkApi sdkEnv.Api) (*DriveItemListItemActivityClient, error) {
	client, err := msgraph.NewClient(sdkApi, "driveitemlistitemactivity", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating DriveItemListItemActivityClient: %+v", err)
	}

	return &DriveItemListItemActivityClient{
		Client: client,
	}, nil
}
