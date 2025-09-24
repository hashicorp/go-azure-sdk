package groupquotasentities

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GroupQuotasEntitiesClient struct {
	Client *resourcemanager.Client
}

func NewGroupQuotasEntitiesClientWithBaseURI(sdkApi sdkEnv.Api) (*GroupQuotasEntitiesClient, error) {
	client, err := resourcemanager.NewClient(sdkApi, "groupquotasentities", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating GroupQuotasEntitiesClient: %+v", err)
	}

	return &GroupQuotasEntitiesClient{
		Client: client,
	}, nil
}
