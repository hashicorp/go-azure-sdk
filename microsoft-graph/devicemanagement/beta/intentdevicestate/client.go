package intentdevicestate

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IntentDeviceStateClient struct {
	Client *msgraph.Client
}

func NewIntentDeviceStateClientWithBaseURI(sdkApi sdkEnv.Api) (*IntentDeviceStateClient, error) {
	client, err := msgraph.NewClient(sdkApi, "intentdevicestate", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating IntentDeviceStateClient: %+v", err)
	}

	return &IntentDeviceStateClient{
		Client: client,
	}, nil
}
