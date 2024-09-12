package driverootchildcontentstream

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DriveRootChildContentStreamClient struct {
	Client *msgraph.Client
}

func NewDriveRootChildContentStreamClientWithBaseURI(sdkApi sdkEnv.Api) (*DriveRootChildContentStreamClient, error) {
	client, err := msgraph.NewClient(sdkApi, "driverootchildcontentstream", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating DriveRootChildContentStreamClient: %+v", err)
	}

	return &DriveRootChildContentStreamClient{
		Client: client,
	}, nil
}
