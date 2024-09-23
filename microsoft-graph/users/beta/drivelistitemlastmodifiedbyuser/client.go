package drivelistitemlastmodifiedbyuser

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DriveListItemLastModifiedByUserClient struct {
	Client *msgraph.Client
}

func NewDriveListItemLastModifiedByUserClientWithBaseURI(sdkApi sdkEnv.Api) (*DriveListItemLastModifiedByUserClient, error) {
	client, err := msgraph.NewClient(sdkApi, "drivelistitemlastmodifiedbyuser", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating DriveListItemLastModifiedByUserClient: %+v", err)
	}

	return &DriveListItemLastModifiedByUserClient{
		Client: client,
	}, nil
}
