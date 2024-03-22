package recoveryservices

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RecoveryServicesClient struct {
	Client *resourcemanager.Client
}

func NewRecoveryServicesClientWithBaseURI(sdkApi sdkEnv.Api) (*RecoveryServicesClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(sdkApi, "recoveryservices", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating RecoveryServicesClient: %+v", err)
	}

	return &RecoveryServicesClient{
		Client: client,
	}, nil
}
