package driveactivitylistitem

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DriveActivityListItemClient struct {
	Client *msgraph.Client
}

func NewDriveActivityListItemClientWithBaseURI(sdkApi sdkEnv.Api) (*DriveActivityListItemClient, error) {
	client, err := msgraph.NewClient(sdkApi, "driveactivitylistitem", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating DriveActivityListItemClient: %+v", err)
	}

	return &DriveActivityListItemClient{
		Client: client,
	}, nil
}
