package drivelistcontenttypecolumnsourcecolumn

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DriveListContentTypeColumnSourceColumnClient struct {
	Client *msgraph.Client
}

func NewDriveListContentTypeColumnSourceColumnClientWithBaseURI(sdkApi sdkEnv.Api) (*DriveListContentTypeColumnSourceColumnClient, error) {
	client, err := msgraph.NewClient(sdkApi, "drivelistcontenttypecolumnsourcecolumn", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating DriveListContentTypeColumnSourceColumnClient: %+v", err)
	}

	return &DriveListContentTypeColumnSourceColumnClient{
		Client: client,
	}, nil
}
