package partnerbillingusagebilled

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PartnerBillingUsageBilledClient struct {
	Client *msgraph.Client
}

func NewPartnerBillingUsageBilledClientWithBaseURI(sdkApi sdkEnv.Api) (*PartnerBillingUsageBilledClient, error) {
	client, err := msgraph.NewClient(sdkApi, "partnerbillingusagebilled", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating PartnerBillingUsageBilledClient: %+v", err)
	}

	return &PartnerBillingUsageBilledClient{
		Client: client,
	}, nil
}
