package drivelistcontenttypebasetype

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DriveListContentTypeBaseTypeClient struct {
	Client *msgraph.Client
}

func NewDriveListContentTypeBaseTypeClientWithBaseURI(sdkApi sdkEnv.Api) (*DriveListContentTypeBaseTypeClient, error) {
	client, err := msgraph.NewClient(sdkApi, "drivelistcontenttypebasetype", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating DriveListContentTypeBaseTypeClient: %+v", err)
	}

	return &DriveListContentTypeBaseTypeClient{
		Client: client,
	}, nil
}
