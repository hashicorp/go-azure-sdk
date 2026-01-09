package driverootlistitemactivitylistitem

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DriveRootListItemActivityListItemClient struct {
	Client *msgraph.Client
}

func NewDriveRootListItemActivityListItemClientWithBaseURI(sdkApi sdkEnv.Api) (*DriveRootListItemActivityListItemClient, error) {
	client, err := msgraph.NewClient(sdkApi, "driverootlistitemactivitylistitem", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating DriveRootListItemActivityListItemClient: %+v", err)
	}

	return &DriveRootListItemActivityListItemClient{
		Client: client,
	}, nil
}
