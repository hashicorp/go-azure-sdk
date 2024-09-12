package drivelistitemcreatedbyuser

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DriveListItemCreatedByUserClient struct {
	Client *msgraph.Client
}

func NewDriveListItemCreatedByUserClientWithBaseURI(sdkApi sdkEnv.Api) (*DriveListItemCreatedByUserClient, error) {
	client, err := msgraph.NewClient(sdkApi, "drivelistitemcreatedbyuser", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating DriveListItemCreatedByUserClient: %+v", err)
	}

	return &DriveListItemCreatedByUserClient{
		Client: client,
	}, nil
}
