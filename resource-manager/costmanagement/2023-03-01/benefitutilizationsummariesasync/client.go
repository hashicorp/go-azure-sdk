package benefitutilizationsummariesasync

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	"github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type BenefitUtilizationSummariesAsyncClient struct {
	Client *resourcemanager.Client
}

func NewBenefitUtilizationSummariesAsyncClientWithBaseURI(api environments.Api) (*BenefitUtilizationSummariesAsyncClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(api, "benefitutilizationsummariesasync", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating BenefitUtilizationSummariesAsyncClient: %+v", err)
	}

	return &BenefitUtilizationSummariesAsyncClient{
		Client: client,
	}, nil
}
