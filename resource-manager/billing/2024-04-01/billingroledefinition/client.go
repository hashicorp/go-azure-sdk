package billingroledefinition

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type BillingRoleDefinitionClient struct {
	Client *resourcemanager.Client
}

func NewBillingRoleDefinitionClientWithBaseURI(sdkApi sdkEnv.Api) (*BillingRoleDefinitionClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(sdkApi, "billingroledefinition", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating BillingRoleDefinitionClient: %+v", err)
	}

	return &BillingRoleDefinitionClient{
		Client: client,
	}, nil
}
