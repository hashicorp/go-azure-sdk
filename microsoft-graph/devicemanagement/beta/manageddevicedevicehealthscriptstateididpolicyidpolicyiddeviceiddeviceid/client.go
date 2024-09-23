package manageddevicedevicehealthscriptstateididpolicyidpolicyiddeviceiddeviceid

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ManagedDeviceDeviceHealthScriptStateIdIdPolicyIdPolicyIdDeviceIdDeviceIdClient struct {
	Client *msgraph.Client
}

func NewManagedDeviceDeviceHealthScriptStateIdIdPolicyIdPolicyIdDeviceIdDeviceIdClientWithBaseURI(sdkApi sdkEnv.Api) (*ManagedDeviceDeviceHealthScriptStateIdIdPolicyIdPolicyIdDeviceIdDeviceIdClient, error) {
	client, err := msgraph.NewClient(sdkApi, "manageddevicedevicehealthscriptstateididpolicyidpolicyiddeviceiddeviceid", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating ManagedDeviceDeviceHealthScriptStateIdIdPolicyIdPolicyIdDeviceIdDeviceIdClient: %+v", err)
	}

	return &ManagedDeviceDeviceHealthScriptStateIdIdPolicyIdPolicyIdDeviceIdDeviceIdClient{
		Client: client,
	}, nil
}
