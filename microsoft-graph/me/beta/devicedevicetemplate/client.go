package devicedevicetemplate

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceDeviceTemplateClient struct {
	Client *msgraph.Client
}

func NewDeviceDeviceTemplateClientWithBaseURI(sdkApi sdkEnv.Api) (*DeviceDeviceTemplateClient, error) {
	client, err := msgraph.NewClient(sdkApi, "devicedevicetemplate", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating DeviceDeviceTemplateClient: %+v", err)
	}

	return &DeviceDeviceTemplateClient{
		Client: client,
	}, nil
}
