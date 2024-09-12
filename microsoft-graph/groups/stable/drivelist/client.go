package drivelist

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DriveListClient struct {
	Client *msgraph.Client
}

func NewDriveListClientWithBaseURI(sdkApi sdkEnv.Api) (*DriveListClient, error) {
	client, err := msgraph.NewClient(sdkApi, "drivelist", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating DriveListClient: %+v", err)
	}

	return &DriveListClient{
		Client: client,
	}, nil
}
