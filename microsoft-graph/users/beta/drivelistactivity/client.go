package drivelistactivity

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DriveListActivityClient struct {
	Client *msgraph.Client
}

func NewDriveListActivityClientWithBaseURI(sdkApi sdkEnv.Api) (*DriveListActivityClient, error) {
	client, err := msgraph.NewClient(sdkApi, "drivelistactivity", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating DriveListActivityClient: %+v", err)
	}

	return &DriveListActivityClient{
		Client: client,
	}, nil
}
