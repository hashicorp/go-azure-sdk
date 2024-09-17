package billingroledefinitions

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type BillingRoleDefinitionsClient struct {
	Client *resourcemanager.Client
}

func NewBillingRoleDefinitionsClientWithBaseURI(sdkApi sdkEnv.Api) (*BillingRoleDefinitionsClient, error) {
	client, err := resourcemanager.NewClient(sdkApi, "billingroledefinitions", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating BillingRoleDefinitionsClient: %+v", err)
	}

	return &BillingRoleDefinitionsClient{
		Client: client,
	}, nil
}
