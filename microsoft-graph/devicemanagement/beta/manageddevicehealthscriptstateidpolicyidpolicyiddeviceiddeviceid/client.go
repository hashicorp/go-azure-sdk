package manageddevicehealthscriptstateidpolicyidpolicyiddeviceiddeviceid

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ManagedDeviceHealthScriptStateIdPolicyIdPolicyIdDeviceIdDeviceIdClient struct {
	Client *msgraph.Client
}

func NewManagedDeviceHealthScriptStateIdPolicyIdPolicyIdDeviceIdDeviceIdClientWithBaseURI(sdkApi sdkEnv.Api) (*ManagedDeviceHealthScriptStateIdPolicyIdPolicyIdDeviceIdDeviceIdClient, error) {
	client, err := msgraph.NewClient(sdkApi, "manageddevicehealthscriptstateidpolicyidpolicyiddeviceiddeviceid", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating ManagedDeviceHealthScriptStateIdPolicyIdPolicyIdDeviceIdDeviceIdClient: %+v", err)
	}

	return &ManagedDeviceHealthScriptStateIdPolicyIdPolicyIdDeviceIdDeviceIdClient{
		Client: client,
	}, nil
}
