package driveitemlistitemactivitylistitem

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DriveItemListItemActivityListItemClient struct {
	Client *msgraph.Client
}

func NewDriveItemListItemActivityListItemClientWithBaseURI(sdkApi sdkEnv.Api) (*DriveItemListItemActivityListItemClient, error) {
	client, err := msgraph.NewClient(sdkApi, "driveitemlistitemactivitylistitem", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating DriveItemListItemActivityListItemClient: %+v", err)
	}

	return &DriveItemListItemActivityListItemClient{
		Client: client,
	}, nil
}
