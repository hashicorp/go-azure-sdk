package serviceconfigurationrecord

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ServiceConfigurationRecordClient struct {
	Client *msgraph.Client
}

func NewServiceConfigurationRecordClientWithBaseURI(sdkApi sdkEnv.Api) (*ServiceConfigurationRecordClient, error) {
	client, err := msgraph.NewClient(sdkApi, "serviceconfigurationrecord", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating ServiceConfigurationRecordClient: %+v", err)
	}

	return &ServiceConfigurationRecordClient{
		Client: client,
	}, nil
}
