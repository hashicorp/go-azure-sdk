package driveitemversioncontent

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DriveItemVersionContentClient struct {
	Client *msgraph.Client
}

func NewDriveItemVersionContentClientWithBaseURI(sdkApi sdkEnv.Api) (*DriveItemVersionContentClient, error) {
	client, err := msgraph.NewClient(sdkApi, "driveitemversioncontent", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating DriveItemVersionContentClient: %+v", err)
	}

	return &DriveItemVersionContentClient{
		Client: client,
	}, nil
}
