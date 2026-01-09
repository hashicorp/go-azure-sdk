package driverootversion

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DriveRootVersionClient struct {
	Client *msgraph.Client
}

func NewDriveRootVersionClientWithBaseURI(sdkApi sdkEnv.Api) (*DriveRootVersionClient, error) {
	client, err := msgraph.NewClient(sdkApi, "driverootversion", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating DriveRootVersionClient: %+v", err)
	}

	return &DriveRootVersionClient{
		Client: client,
	}, nil
}
