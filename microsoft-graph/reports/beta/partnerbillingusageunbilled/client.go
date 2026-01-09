package partnerbillingusageunbilled

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PartnerBillingUsageUnbilledClient struct {
	Client *msgraph.Client
}

func NewPartnerBillingUsageUnbilledClientWithBaseURI(sdkApi sdkEnv.Api) (*PartnerBillingUsageUnbilledClient, error) {
	client, err := msgraph.NewClient(sdkApi, "partnerbillingusageunbilled", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating PartnerBillingUsageUnbilledClient: %+v", err)
	}

	return &PartnerBillingUsageUnbilledClient{
		Client: client,
	}, nil
}
