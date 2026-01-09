package driveitemcreatedbyuser

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DriveItemCreatedByUserClient struct {
	Client *msgraph.Client
}

func NewDriveItemCreatedByUserClientWithBaseURI(sdkApi sdkEnv.Api) (*DriveItemCreatedByUserClient, error) {
	client, err := msgraph.NewClient(sdkApi, "driveitemcreatedbyuser", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating DriveItemCreatedByUserClient: %+v", err)
	}

	return &DriveItemCreatedByUserClient{
		Client: client,
	}, nil
}
