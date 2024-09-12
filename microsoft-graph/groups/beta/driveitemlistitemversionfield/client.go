package driveitemlistitemversionfield

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DriveItemListItemVersionFieldClient struct {
	Client *msgraph.Client
}

func NewDriveItemListItemVersionFieldClientWithBaseURI(sdkApi sdkEnv.Api) (*DriveItemListItemVersionFieldClient, error) {
	client, err := msgraph.NewClient(sdkApi, "driveitemlistitemversionfield", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating DriveItemListItemVersionFieldClient: %+v", err)
	}

	return &DriveItemListItemVersionFieldClient{
		Client: client,
	}, nil
}
