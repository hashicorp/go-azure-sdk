package memanageddevice

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MeManagedDeviceClient struct {
	Client *msgraph.Client
}

func NewMeManagedDeviceClientWithBaseURI(api sdkEnv.Api) (*MeManagedDeviceClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "memanageddevice", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating MeManagedDeviceClient: %+v", err)
	}

	return &MeManagedDeviceClient{
		Client: client,
	}, nil
}
