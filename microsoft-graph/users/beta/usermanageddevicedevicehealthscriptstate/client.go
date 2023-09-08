package usermanageddevicedevicehealthscriptstate

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserManagedDeviceDeviceHealthScriptStateClient struct {
	Client *msgraph.Client
}

func NewUserManagedDeviceDeviceHealthScriptStateClientWithBaseURI(api sdkEnv.Api) (*UserManagedDeviceDeviceHealthScriptStateClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "usermanageddevicedevicehealthscriptstate", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating UserManagedDeviceDeviceHealthScriptStateClient: %+v", err)
	}

	return &UserManagedDeviceDeviceHealthScriptStateClient{
		Client: client,
	}, nil
}
