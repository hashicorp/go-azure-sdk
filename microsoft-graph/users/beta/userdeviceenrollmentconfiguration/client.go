package userdeviceenrollmentconfiguration

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserDeviceEnrollmentConfigurationClient struct {
	Client *msgraph.Client
}

func NewUserDeviceEnrollmentConfigurationClientWithBaseURI(api sdkEnv.Api) (*UserDeviceEnrollmentConfigurationClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "userdeviceenrollmentconfiguration", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating UserDeviceEnrollmentConfigurationClient: %+v", err)
	}

	return &UserDeviceEnrollmentConfigurationClient{
		Client: client,
	}, nil
}
