package driverootlistitemactivity

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DriveRootListItemActivityClient struct {
	Client *msgraph.Client
}

func NewDriveRootListItemActivityClientWithBaseURI(sdkApi sdkEnv.Api) (*DriveRootListItemActivityClient, error) {
	client, err := msgraph.NewClient(sdkApi, "driverootlistitemactivity", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating DriveRootListItemActivityClient: %+v", err)
	}

	return &DriveRootListItemActivityClient{
		Client: client,
	}, nil
}
