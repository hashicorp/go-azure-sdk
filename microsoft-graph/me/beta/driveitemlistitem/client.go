package driveitemlistitem

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DriveItemListItemClient struct {
	Client *msgraph.Client
}

func NewDriveItemListItemClientWithBaseURI(sdkApi sdkEnv.Api) (*DriveItemListItemClient, error) {
	client, err := msgraph.NewClient(sdkApi, "driveitemlistitem", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating DriveItemListItemClient: %+v", err)
	}

	return &DriveItemListItemClient{
		Client: client,
	}, nil
}
