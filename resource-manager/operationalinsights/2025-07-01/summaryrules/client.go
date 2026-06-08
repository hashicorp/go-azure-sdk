package summaryrules

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SummaryRulesClient struct {
	Client *resourcemanager.Client
}

func NewSummaryRulesClientWithBaseURI(sdkApi sdkEnv.Api) (*SummaryRulesClient, error) {
	client, err := resourcemanager.NewClient(sdkApi, "summaryrules", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating SummaryRulesClient: %+v", err)
	}

	return &SummaryRulesClient{
		Client: client,
	}, nil
}
