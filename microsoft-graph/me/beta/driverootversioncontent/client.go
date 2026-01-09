package driverootversioncontent

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DriveRootVersionContentClient struct {
	Client *msgraph.Client
}

func NewDriveRootVersionContentClientWithBaseURI(sdkApi sdkEnv.Api) (*DriveRootVersionContentClient, error) {
	client, err := msgraph.NewClient(sdkApi, "driverootversioncontent", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating DriveRootVersionContentClient: %+v", err)
	}

	return &DriveRootVersionContentClient{
		Client: client,
	}, nil
}
