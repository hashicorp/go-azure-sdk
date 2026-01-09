package driverootlistitemversionfield

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DriveRootListItemVersionFieldClient struct {
	Client *msgraph.Client
}

func NewDriveRootListItemVersionFieldClientWithBaseURI(sdkApi sdkEnv.Api) (*DriveRootListItemVersionFieldClient, error) {
	client, err := msgraph.NewClient(sdkApi, "driverootlistitemversionfield", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating DriveRootListItemVersionFieldClient: %+v", err)
	}

	return &DriveRootListItemVersionFieldClient{
		Client: client,
	}, nil
}
