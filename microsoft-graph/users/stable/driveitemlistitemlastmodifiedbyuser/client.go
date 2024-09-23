package driveitemlistitemlastmodifiedbyuser

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DriveItemListItemLastModifiedByUserClient struct {
	Client *msgraph.Client
}

func NewDriveItemListItemLastModifiedByUserClientWithBaseURI(sdkApi sdkEnv.Api) (*DriveItemListItemLastModifiedByUserClient, error) {
	client, err := msgraph.NewClient(sdkApi, "driveitemlistitemlastmodifiedbyuser", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating DriveItemListItemLastModifiedByUserClient: %+v", err)
	}

	return &DriveItemListItemLastModifiedByUserClient{
		Client: client,
	}, nil
}
