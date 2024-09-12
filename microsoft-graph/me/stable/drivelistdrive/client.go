package drivelistdrive

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DriveListDriveClient struct {
	Client *msgraph.Client
}

func NewDriveListDriveClientWithBaseURI(sdkApi sdkEnv.Api) (*DriveListDriveClient, error) {
	client, err := msgraph.NewClient(sdkApi, "drivelistdrive", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating DriveListDriveClient: %+v", err)
	}

	return &DriveListDriveClient{
		Client: client,
	}, nil
}
