package driverootcontentstream

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DriveRootContentStreamClient struct {
	Client *msgraph.Client
}

func NewDriveRootContentStreamClientWithBaseURI(sdkApi sdkEnv.Api) (*DriveRootContentStreamClient, error) {
	client, err := msgraph.NewClient(sdkApi, "driverootcontentstream", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating DriveRootContentStreamClient: %+v", err)
	}

	return &DriveRootContentStreamClient{
		Client: client,
	}, nil
}
