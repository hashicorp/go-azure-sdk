package driverootlistitemcreatedbyuser

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DriveRootListItemCreatedByUserClient struct {
	Client *msgraph.Client
}

func NewDriveRootListItemCreatedByUserClientWithBaseURI(sdkApi sdkEnv.Api) (*DriveRootListItemCreatedByUserClient, error) {
	client, err := msgraph.NewClient(sdkApi, "driverootlistitemcreatedbyuser", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating DriveRootListItemCreatedByUserClient: %+v", err)
	}

	return &DriveRootListItemCreatedByUserClient{
		Client: client,
	}, nil
}
