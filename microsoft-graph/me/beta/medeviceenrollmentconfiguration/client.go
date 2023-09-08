package medeviceenrollmentconfiguration

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MeDeviceEnrollmentConfigurationClient struct {
	Client *msgraph.Client
}

func NewMeDeviceEnrollmentConfigurationClientWithBaseURI(api sdkEnv.Api) (*MeDeviceEnrollmentConfigurationClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "medeviceenrollmentconfiguration", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating MeDeviceEnrollmentConfigurationClient: %+v", err)
	}

	return &MeDeviceEnrollmentConfigurationClient{
		Client: client,
	}, nil
}
