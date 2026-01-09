package drivecreatedbyuser

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DriveCreatedByUserClient struct {
	Client *msgraph.Client
}

func NewDriveCreatedByUserClientWithBaseURI(sdkApi sdkEnv.Api) (*DriveCreatedByUserClient, error) {
	client, err := msgraph.NewClient(sdkApi, "drivecreatedbyuser", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating DriveCreatedByUserClient: %+v", err)
	}

	return &DriveCreatedByUserClient{
		Client: client,
	}, nil
}
