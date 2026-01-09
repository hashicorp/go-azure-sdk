package drivefollowingcontentstream

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DriveFollowingContentStreamClient struct {
	Client *msgraph.Client
}

func NewDriveFollowingContentStreamClientWithBaseURI(sdkApi sdkEnv.Api) (*DriveFollowingContentStreamClient, error) {
	client, err := msgraph.NewClient(sdkApi, "drivefollowingcontentstream", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating DriveFollowingContentStreamClient: %+v", err)
	}

	return &DriveFollowingContentStreamClient{
		Client: client,
	}, nil
}
