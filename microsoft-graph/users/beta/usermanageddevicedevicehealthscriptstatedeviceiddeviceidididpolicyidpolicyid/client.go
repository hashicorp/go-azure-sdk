package usermanageddevicedevicehealthscriptstatedeviceiddeviceidididpolicyidpolicyid

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserManagedDeviceDeviceHealthScriptStateDeviceIdDeviceIdIdIdPolicyIdPolicyIdClient struct {
	Client *msgraph.Client
}

func NewUserManagedDeviceDeviceHealthScriptStateDeviceIdDeviceIdIdIdPolicyIdPolicyIdClientWithBaseURI(api sdkEnv.Api) (*UserManagedDeviceDeviceHealthScriptStateDeviceIdDeviceIdIdIdPolicyIdPolicyIdClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "usermanageddevicedevicehealthscriptstatedeviceiddeviceidididpolicyidpolicyid", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating UserManagedDeviceDeviceHealthScriptStateDeviceIdDeviceIdIdIdPolicyIdPolicyIdClient: %+v", err)
	}

	return &UserManagedDeviceDeviceHealthScriptStateDeviceIdDeviceIdIdIdPolicyIdPolicyIdClient{
		Client: client,
	}, nil
}
