package driveroot

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DriveRootClient struct {
	Client *msgraph.Client
}

func NewDriveRootClientWithBaseURI(sdkApi sdkEnv.Api) (*DriveRootClient, error) {
	client, err := msgraph.NewClient(sdkApi, "driveroot", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating DriveRootClient: %+v", err)
	}

	return &DriveRootClient{
		Client: client,
	}, nil
}
