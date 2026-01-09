package driveitemretentionlabel

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DriveItemRetentionLabelClient struct {
	Client *msgraph.Client
}

func NewDriveItemRetentionLabelClientWithBaseURI(sdkApi sdkEnv.Api) (*DriveItemRetentionLabelClient, error) {
	client, err := msgraph.NewClient(sdkApi, "driveitemretentionlabel", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating DriveItemRetentionLabelClient: %+v", err)
	}

	return &DriveItemRetentionLabelClient{
		Client: client,
	}, nil
}
