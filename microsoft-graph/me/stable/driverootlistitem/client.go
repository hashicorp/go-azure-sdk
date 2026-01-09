package driverootlistitem

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DriveRootListItemClient struct {
	Client *msgraph.Client
}

func NewDriveRootListItemClientWithBaseURI(sdkApi sdkEnv.Api) (*DriveRootListItemClient, error) {
	client, err := msgraph.NewClient(sdkApi, "driverootlistitem", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating DriveRootListItemClient: %+v", err)
	}

	return &DriveRootListItemClient{
		Client: client,
	}, nil
}
