package driveitemchild

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DriveItemChildClient struct {
	Client *msgraph.Client
}

func NewDriveItemChildClientWithBaseURI(sdkApi sdkEnv.Api) (*DriveItemChildClient, error) {
	client, err := msgraph.NewClient(sdkApi, "driveitemchild", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating DriveItemChildClient: %+v", err)
	}

	return &DriveItemChildClient{
		Client: client,
	}, nil
}
