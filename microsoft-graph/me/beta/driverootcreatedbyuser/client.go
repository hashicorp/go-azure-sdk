package driverootcreatedbyuser

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DriveRootCreatedByUserClient struct {
	Client *msgraph.Client
}

func NewDriveRootCreatedByUserClientWithBaseURI(sdkApi sdkEnv.Api) (*DriveRootCreatedByUserClient, error) {
	client, err := msgraph.NewClient(sdkApi, "driverootcreatedbyuser", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating DriveRootCreatedByUserClient: %+v", err)
	}

	return &DriveRootCreatedByUserClient{
		Client: client,
	}, nil
}
