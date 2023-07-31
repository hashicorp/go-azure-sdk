package quotabyperiodkeys

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type QuotaByPeriodKeysClient struct {
	Client *resourcemanager.Client
}

func NewQuotaByPeriodKeysClientWithBaseURI(api sdkEnv.Api) (*QuotaByPeriodKeysClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(api, "quotabyperiodkeys", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating QuotaByPeriodKeysClient: %+v", err)
	}

	return &QuotaByPeriodKeysClient{
		Client: client,
	}, nil
}
