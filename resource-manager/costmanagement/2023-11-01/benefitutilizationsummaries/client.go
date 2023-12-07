package benefitutilizationsummaries

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type BenefitUtilizationSummariesClient struct {
	Client *resourcemanager.Client
}

func NewBenefitUtilizationSummariesClientWithBaseURI(sdkApi sdkEnv.Api) (*BenefitUtilizationSummariesClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(sdkApi, "benefitutilizationsummaries", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating BenefitUtilizationSummariesClient: %+v", err)
	}

	return &BenefitUtilizationSummariesClient{
		Client: client,
	}, nil
}
