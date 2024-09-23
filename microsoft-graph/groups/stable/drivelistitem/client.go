package drivelistitem

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DriveListItemClient struct {
	Client *msgraph.Client
}

func NewDriveListItemClientWithBaseURI(sdkApi sdkEnv.Api) (*DriveListItemClient, error) {
	client, err := msgraph.NewClient(sdkApi, "drivelistitem", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating DriveListItemClient: %+v", err)
	}

	return &DriveListItemClient{
		Client: client,
	}, nil
}
