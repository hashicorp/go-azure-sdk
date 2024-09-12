package drivelistitemactivity

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DriveListItemActivityClient struct {
	Client *msgraph.Client
}

func NewDriveListItemActivityClientWithBaseURI(sdkApi sdkEnv.Api) (*DriveListItemActivityClient, error) {
	client, err := msgraph.NewClient(sdkApi, "drivelistitemactivity", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating DriveListItemActivityClient: %+v", err)
	}

	return &DriveListItemActivityClient{
		Client: client,
	}, nil
}
