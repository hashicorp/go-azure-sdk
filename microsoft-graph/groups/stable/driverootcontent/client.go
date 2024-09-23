package driverootcontent

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DriveRootContentClient struct {
	Client *msgraph.Client
}

func NewDriveRootContentClientWithBaseURI(sdkApi sdkEnv.Api) (*DriveRootContentClient, error) {
	client, err := msgraph.NewClient(sdkApi, "driverootcontent", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating DriveRootContentClient: %+v", err)
	}

	return &DriveRootContentClient{
		Client: client,
	}, nil
}
