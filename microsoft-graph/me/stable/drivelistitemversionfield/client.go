package drivelistitemversionfield

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DriveListItemVersionFieldClient struct {
	Client *msgraph.Client
}

func NewDriveListItemVersionFieldClientWithBaseURI(sdkApi sdkEnv.Api) (*DriveListItemVersionFieldClient, error) {
	client, err := msgraph.NewClient(sdkApi, "drivelistitemversionfield", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating DriveListItemVersionFieldClient: %+v", err)
	}

	return &DriveListItemVersionFieldClient{
		Client: client,
	}, nil
}
