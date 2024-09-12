package driveitem

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DriveItemClient struct {
	Client *msgraph.Client
}

func NewDriveItemClientWithBaseURI(sdkApi sdkEnv.Api) (*DriveItemClient, error) {
	client, err := msgraph.NewClient(sdkApi, "driveitem", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating DriveItemClient: %+v", err)
	}

	return &DriveItemClient{
		Client: client,
	}, nil
}
