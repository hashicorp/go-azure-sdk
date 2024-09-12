package drivelistcontenttypebase

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DriveListContentTypeBaseClient struct {
	Client *msgraph.Client
}

func NewDriveListContentTypeBaseClientWithBaseURI(sdkApi sdkEnv.Api) (*DriveListContentTypeBaseClient, error) {
	client, err := msgraph.NewClient(sdkApi, "drivelistcontenttypebase", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating DriveListContentTypeBaseClient: %+v", err)
	}

	return &DriveListContentTypeBaseClient{
		Client: client,
	}, nil
}
