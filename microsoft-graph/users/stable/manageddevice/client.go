package manageddevice

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ManagedDeviceClient struct {
	Client *msgraph.Client
}

func NewManagedDeviceClientWithBaseURI(sdkApi sdkEnv.Api) (*ManagedDeviceClient, error) {
	client, err := msgraph.NewClient(sdkApi, "manageddevice", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating ManagedDeviceClient: %+v", err)
	}

	return &ManagedDeviceClient{
		Client: client,
	}, nil
}
