package managedhsmkeyoperationgroup

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ManagedHsmKeyOperationGroupClient struct {
	Client *resourcemanager.Client
}

func NewManagedHsmKeyOperationGroupClientWithBaseURI(sdkApi sdkEnv.Api) (*ManagedHsmKeyOperationGroupClient, error) {
	client, err := resourcemanager.NewClient(sdkApi, "managedhsmkeyoperationgroup", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating ManagedHsmKeyOperationGroupClient: %+v", err)
	}

	return &ManagedHsmKeyOperationGroupClient{
		Client: client,
	}, nil
}
