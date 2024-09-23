package driverootlistitemversion

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DriveRootListItemVersionClient struct {
	Client *msgraph.Client
}

func NewDriveRootListItemVersionClientWithBaseURI(sdkApi sdkEnv.Api) (*DriveRootListItemVersionClient, error) {
	client, err := msgraph.NewClient(sdkApi, "driverootlistitemversion", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating DriveRootListItemVersionClient: %+v", err)
	}

	return &DriveRootListItemVersionClient{
		Client: client,
	}, nil
}
