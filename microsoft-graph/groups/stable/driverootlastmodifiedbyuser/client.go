package driverootlastmodifiedbyuser

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DriveRootLastModifiedByUserClient struct {
	Client *msgraph.Client
}

func NewDriveRootLastModifiedByUserClientWithBaseURI(sdkApi sdkEnv.Api) (*DriveRootLastModifiedByUserClient, error) {
	client, err := msgraph.NewClient(sdkApi, "driverootlastmodifiedbyuser", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating DriveRootLastModifiedByUserClient: %+v", err)
	}

	return &DriveRootLastModifiedByUserClient{
		Client: client,
	}, nil
}
