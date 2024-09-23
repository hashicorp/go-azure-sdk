package drive

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DriveClient struct {
	Client *msgraph.Client
}

func NewDriveClientWithBaseURI(sdkApi sdkEnv.Api) (*DriveClient, error) {
	client, err := msgraph.NewClient(sdkApi, "drive", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating DriveClient: %+v", err)
	}

	return &DriveClient{
		Client: client,
	}, nil
}
