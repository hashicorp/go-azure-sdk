package memanageddevicedetectedapp

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MeManagedDeviceDetectedAppClient struct {
	Client *msgraph.Client
}

func NewMeManagedDeviceDetectedAppClientWithBaseURI(api sdkEnv.Api) (*MeManagedDeviceDetectedAppClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "memanageddevicedetectedapp", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating MeManagedDeviceDetectedAppClient: %+v", err)
	}

	return &MeManagedDeviceDetectedAppClient{
		Client: client,
	}, nil
}
