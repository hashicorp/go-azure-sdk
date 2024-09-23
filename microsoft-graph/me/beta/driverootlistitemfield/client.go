package driverootlistitemfield

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DriveRootListItemFieldClient struct {
	Client *msgraph.Client
}

func NewDriveRootListItemFieldClientWithBaseURI(sdkApi sdkEnv.Api) (*DriveRootListItemFieldClient, error) {
	client, err := msgraph.NewClient(sdkApi, "driverootlistitemfield", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating DriveRootListItemFieldClient: %+v", err)
	}

	return &DriveRootListItemFieldClient{
		Client: client,
	}, nil
}
