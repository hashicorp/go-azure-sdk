package driverootthumbnail

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DriveRootThumbnailClient struct {
	Client *msgraph.Client
}

func NewDriveRootThumbnailClientWithBaseURI(sdkApi sdkEnv.Api) (*DriveRootThumbnailClient, error) {
	client, err := msgraph.NewClient(sdkApi, "driverootthumbnail", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating DriveRootThumbnailClient: %+v", err)
	}

	return &DriveRootThumbnailClient{
		Client: client,
	}, nil
}
