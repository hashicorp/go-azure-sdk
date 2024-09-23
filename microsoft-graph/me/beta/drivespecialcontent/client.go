package drivespecialcontent

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DriveSpecialContentClient struct {
	Client *msgraph.Client
}

func NewDriveSpecialContentClientWithBaseURI(sdkApi sdkEnv.Api) (*DriveSpecialContentClient, error) {
	client, err := msgraph.NewClient(sdkApi, "drivespecialcontent", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating DriveSpecialContentClient: %+v", err)
	}

	return &DriveSpecialContentClient{
		Client: client,
	}, nil
}
