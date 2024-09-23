package driverootactivity

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DriveRootActivityClient struct {
	Client *msgraph.Client
}

func NewDriveRootActivityClientWithBaseURI(sdkApi sdkEnv.Api) (*DriveRootActivityClient, error) {
	client, err := msgraph.NewClient(sdkApi, "driverootactivity", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating DriveRootActivityClient: %+v", err)
	}

	return &DriveRootActivityClient{
		Client: client,
	}, nil
}
