package drivefollowing

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DriveFollowingClient struct {
	Client *msgraph.Client
}

func NewDriveFollowingClientWithBaseURI(sdkApi sdkEnv.Api) (*DriveFollowingClient, error) {
	client, err := msgraph.NewClient(sdkApi, "drivefollowing", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating DriveFollowingClient: %+v", err)
	}

	return &DriveFollowingClient{
		Client: client,
	}, nil
}
