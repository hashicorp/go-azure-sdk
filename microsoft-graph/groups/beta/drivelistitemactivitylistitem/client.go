package drivelistitemactivitylistitem

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DriveListItemActivityListItemClient struct {
	Client *msgraph.Client
}

func NewDriveListItemActivityListItemClientWithBaseURI(sdkApi sdkEnv.Api) (*DriveListItemActivityListItemClient, error) {
	client, err := msgraph.NewClient(sdkApi, "drivelistitemactivitylistitem", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating DriveListItemActivityListItemClient: %+v", err)
	}

	return &DriveListItemActivityListItemClient{
		Client: client,
	}, nil
}
