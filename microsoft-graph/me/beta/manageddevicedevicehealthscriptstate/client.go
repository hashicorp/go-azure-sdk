package manageddevicedevicehealthscriptstate

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ManagedDeviceDeviceHealthScriptStateClient struct {
	Client *msgraph.Client
}

func NewManagedDeviceDeviceHealthScriptStateClientWithBaseURI(sdkApi sdkEnv.Api) (*ManagedDeviceDeviceHealthScriptStateClient, error) {
	client, err := msgraph.NewClient(sdkApi, "manageddevicedevicehealthscriptstate", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating ManagedDeviceDeviceHealthScriptStateClient: %+v", err)
	}

	return &ManagedDeviceDeviceHealthScriptStateClient{
		Client: client,
	}, nil
}
