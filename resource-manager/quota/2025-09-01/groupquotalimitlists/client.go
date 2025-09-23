package groupquotalimitlists

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GroupQuotaLimitListsClient struct {
	Client *resourcemanager.Client
}

func NewGroupQuotaLimitListsClientWithBaseURI(sdkApi sdkEnv.Api) (*GroupQuotaLimitListsClient, error) {
	client, err := resourcemanager.NewClient(sdkApi, "groupquotalimitlists", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating GroupQuotaLimitListsClient: %+v", err)
	}

	return &GroupQuotaLimitListsClient{
		Client: client,
	}, nil
}
