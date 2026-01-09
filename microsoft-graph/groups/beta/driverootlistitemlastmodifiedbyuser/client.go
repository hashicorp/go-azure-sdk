package driverootlistitemlastmodifiedbyuser

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DriveRootListItemLastModifiedByUserClient struct {
	Client *msgraph.Client
}

func NewDriveRootListItemLastModifiedByUserClientWithBaseURI(sdkApi sdkEnv.Api) (*DriveRootListItemLastModifiedByUserClient, error) {
	client, err := msgraph.NewClient(sdkApi, "driverootlistitemlastmodifiedbyuser", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating DriveRootListItemLastModifiedByUserClient: %+v", err)
	}

	return &DriveRootListItemLastModifiedByUserClient{
		Client: client,
	}, nil
}
