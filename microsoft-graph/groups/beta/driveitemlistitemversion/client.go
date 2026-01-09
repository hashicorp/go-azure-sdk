package driveitemlistitemversion

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DriveItemListItemVersionClient struct {
	Client *msgraph.Client
}

func NewDriveItemListItemVersionClientWithBaseURI(sdkApi sdkEnv.Api) (*DriveItemListItemVersionClient, error) {
	client, err := msgraph.NewClient(sdkApi, "driveitemlistitemversion", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating DriveItemListItemVersionClient: %+v", err)
	}

	return &DriveItemListItemVersionClient{
		Client: client,
	}, nil
}
