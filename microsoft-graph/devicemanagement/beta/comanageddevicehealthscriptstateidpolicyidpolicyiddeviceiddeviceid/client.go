package comanageddevicehealthscriptstateidpolicyidpolicyiddeviceiddeviceid

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ComanagedDeviceHealthScriptStateIdPolicyIdPolicyIdDeviceIdDeviceIdClient struct {
	Client *msgraph.Client
}

func NewComanagedDeviceHealthScriptStateIdPolicyIdPolicyIdDeviceIdDeviceIdClientWithBaseURI(sdkApi sdkEnv.Api) (*ComanagedDeviceHealthScriptStateIdPolicyIdPolicyIdDeviceIdDeviceIdClient, error) {
	client, err := msgraph.NewClient(sdkApi, "comanageddevicehealthscriptstateidpolicyidpolicyiddeviceiddeviceid", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating ComanagedDeviceHealthScriptStateIdPolicyIdPolicyIdDeviceIdDeviceIdClient: %+v", err)
	}

	return &ComanagedDeviceHealthScriptStateIdPolicyIdPolicyIdDeviceIdDeviceIdClient{
		Client: client,
	}, nil
}
