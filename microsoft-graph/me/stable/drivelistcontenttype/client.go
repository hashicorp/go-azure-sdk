package drivelistcontenttype

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DriveListContentTypeClient struct {
	Client *msgraph.Client
}

func NewDriveListContentTypeClientWithBaseURI(sdkApi sdkEnv.Api) (*DriveListContentTypeClient, error) {
	client, err := msgraph.NewClient(sdkApi, "drivelistcontenttype", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating DriveListContentTypeClient: %+v", err)
	}

	return &DriveListContentTypeClient{
		Client: client,
	}, nil
}
