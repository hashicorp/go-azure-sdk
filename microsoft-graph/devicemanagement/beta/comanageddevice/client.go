package comanageddevice

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ComanagedDeviceClient struct {
	Client *msgraph.Client
}

func NewComanagedDeviceClientWithBaseURI(sdkApi sdkEnv.Api) (*ComanagedDeviceClient, error) {
	client, err := msgraph.NewClient(sdkApi, "comanageddevice", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating ComanagedDeviceClient: %+v", err)
	}

	return &ComanagedDeviceClient{
		Client: client,
	}, nil
}
