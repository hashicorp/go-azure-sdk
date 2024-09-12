package drivelistlastmodifiedbyuser

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DriveListLastModifiedByUserClient struct {
	Client *msgraph.Client
}

func NewDriveListLastModifiedByUserClientWithBaseURI(sdkApi sdkEnv.Api) (*DriveListLastModifiedByUserClient, error) {
	client, err := msgraph.NewClient(sdkApi, "drivelistlastmodifiedbyuser", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating DriveListLastModifiedByUserClient: %+v", err)
	}

	return &DriveListLastModifiedByUserClient{
		Client: client,
	}, nil
}
