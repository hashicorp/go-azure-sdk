package medrive

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MeDriveClient struct {
	Client *msgraph.Client
}

func NewMeDriveClientWithBaseURI(api sdkEnv.Api) (*MeDriveClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "medrive", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating MeDriveClient: %+v", err)
	}

	return &MeDriveClient{
		Client: client,
	}, nil
}
