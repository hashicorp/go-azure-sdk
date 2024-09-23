package driveactivity

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DriveActivityClient struct {
	Client *msgraph.Client
}

func NewDriveActivityClientWithBaseURI(sdkApi sdkEnv.Api) (*DriveActivityClient, error) {
	client, err := msgraph.NewClient(sdkApi, "driveactivity", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating DriveActivityClient: %+v", err)
	}

	return &DriveActivityClient{
		Client: client,
	}, nil
}
