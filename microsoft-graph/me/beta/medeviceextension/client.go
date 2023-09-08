package medeviceextension

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MeDeviceExtensionClient struct {
	Client *msgraph.Client
}

func NewMeDeviceExtensionClientWithBaseURI(api sdkEnv.Api) (*MeDeviceExtensionClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "medeviceextension", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating MeDeviceExtensionClient: %+v", err)
	}

	return &MeDeviceExtensionClient{
		Client: client,
	}, nil
}
