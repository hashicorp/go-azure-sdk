package driveitemactivity

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DriveItemActivityClient struct {
	Client *msgraph.Client
}

func NewDriveItemActivityClientWithBaseURI(sdkApi sdkEnv.Api) (*DriveItemActivityClient, error) {
	client, err := msgraph.NewClient(sdkApi, "driveitemactivity", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating DriveItemActivityClient: %+v", err)
	}

	return &DriveItemActivityClient{
		Client: client,
	}, nil
}
