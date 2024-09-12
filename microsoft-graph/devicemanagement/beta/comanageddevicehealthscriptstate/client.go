package comanageddevicehealthscriptstate

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ComanagedDeviceHealthScriptStateClient struct {
	Client *msgraph.Client
}

func NewComanagedDeviceHealthScriptStateClientWithBaseURI(sdkApi sdkEnv.Api) (*ComanagedDeviceHealthScriptStateClient, error) {
	client, err := msgraph.NewClient(sdkApi, "comanageddevicehealthscriptstate", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating ComanagedDeviceHealthScriptStateClient: %+v", err)
	}

	return &ComanagedDeviceHealthScriptStateClient{
		Client: client,
	}, nil
}
