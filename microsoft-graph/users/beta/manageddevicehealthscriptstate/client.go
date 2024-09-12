package manageddevicehealthscriptstate

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ManagedDeviceHealthScriptStateClient struct {
	Client *msgraph.Client
}

func NewManagedDeviceHealthScriptStateClientWithBaseURI(sdkApi sdkEnv.Api) (*ManagedDeviceHealthScriptStateClient, error) {
	client, err := msgraph.NewClient(sdkApi, "manageddevicehealthscriptstate", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating ManagedDeviceHealthScriptStateClient: %+v", err)
	}

	return &ManagedDeviceHealthScriptStateClient{
		Client: client,
	}, nil
}
