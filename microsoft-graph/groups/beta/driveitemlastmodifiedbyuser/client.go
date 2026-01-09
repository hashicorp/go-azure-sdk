package driveitemlastmodifiedbyuser

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DriveItemLastModifiedByUserClient struct {
	Client *msgraph.Client
}

func NewDriveItemLastModifiedByUserClientWithBaseURI(sdkApi sdkEnv.Api) (*DriveItemLastModifiedByUserClient, error) {
	client, err := msgraph.NewClient(sdkApi, "driveitemlastmodifiedbyuser", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating DriveItemLastModifiedByUserClient: %+v", err)
	}

	return &DriveItemLastModifiedByUserClient{
		Client: client,
	}, nil
}
