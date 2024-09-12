package drivelistcontenttypecolumnposition

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DriveListContentTypeColumnPositionClient struct {
	Client *msgraph.Client
}

func NewDriveListContentTypeColumnPositionClientWithBaseURI(sdkApi sdkEnv.Api) (*DriveListContentTypeColumnPositionClient, error) {
	client, err := msgraph.NewClient(sdkApi, "drivelistcontenttypecolumnposition", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating DriveListContentTypeColumnPositionClient: %+v", err)
	}

	return &DriveListContentTypeColumnPositionClient{
		Client: client,
	}, nil
}
