package detectedappmanageddevice

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DetectedAppManagedDeviceClient struct {
	Client *msgraph.Client
}

func NewDetectedAppManagedDeviceClientWithBaseURI(sdkApi sdkEnv.Api) (*DetectedAppManagedDeviceClient, error) {
	client, err := msgraph.NewClient(sdkApi, "detectedappmanageddevice", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating DetectedAppManagedDeviceClient: %+v", err)
	}

	return &DetectedAppManagedDeviceClient{
		Client: client,
	}, nil
}
