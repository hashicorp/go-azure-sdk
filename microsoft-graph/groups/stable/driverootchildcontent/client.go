package driverootchildcontent

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DriveRootChildContentClient struct {
	Client *msgraph.Client
}

func NewDriveRootChildContentClientWithBaseURI(sdkApi sdkEnv.Api) (*DriveRootChildContentClient, error) {
	client, err := msgraph.NewClient(sdkApi, "driverootchildcontent", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating DriveRootChildContentClient: %+v", err)
	}

	return &DriveRootChildContentClient{
		Client: client,
	}, nil
}
