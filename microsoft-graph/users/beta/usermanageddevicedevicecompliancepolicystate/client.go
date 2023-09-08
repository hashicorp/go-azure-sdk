package usermanageddevicedevicecompliancepolicystate

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserManagedDeviceDeviceCompliancePolicyStateClient struct {
	Client *msgraph.Client
}

func NewUserManagedDeviceDeviceCompliancePolicyStateClientWithBaseURI(api sdkEnv.Api) (*UserManagedDeviceDeviceCompliancePolicyStateClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "usermanageddevicedevicecompliancepolicystate", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating UserManagedDeviceDeviceCompliancePolicyStateClient: %+v", err)
	}

	return &UserManagedDeviceDeviceCompliancePolicyStateClient{
		Client: client,
	}, nil
}
