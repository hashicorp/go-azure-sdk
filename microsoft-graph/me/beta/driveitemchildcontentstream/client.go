package driveitemchildcontentstream

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DriveItemChildContentStreamClient struct {
	Client *msgraph.Client
}

func NewDriveItemChildContentStreamClientWithBaseURI(sdkApi sdkEnv.Api) (*DriveItemChildContentStreamClient, error) {
	client, err := msgraph.NewClient(sdkApi, "driveitemchildcontentstream", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating DriveItemChildContentStreamClient: %+v", err)
	}

	return &DriveItemChildContentStreamClient{
		Client: client,
	}, nil
}
