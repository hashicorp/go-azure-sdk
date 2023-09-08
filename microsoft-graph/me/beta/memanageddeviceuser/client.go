package memanageddeviceuser

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MeManagedDeviceUserClient struct {
	Client *msgraph.Client
}

func NewMeManagedDeviceUserClientWithBaseURI(api sdkEnv.Api) (*MeManagedDeviceUserClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "memanageddeviceuser", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating MeManagedDeviceUserClient: %+v", err)
	}

	return &MeManagedDeviceUserClient{
		Client: client,
	}, nil
}
