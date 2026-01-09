package partnerbillingreconciliation

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PartnerBillingReconciliationClient struct {
	Client *msgraph.Client
}

func NewPartnerBillingReconciliationClientWithBaseURI(sdkApi sdkEnv.Api) (*PartnerBillingReconciliationClient, error) {
	client, err := msgraph.NewClient(sdkApi, "partnerbillingreconciliation", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating PartnerBillingReconciliationClient: %+v", err)
	}

	return &PartnerBillingReconciliationClient{
		Client: client,
	}, nil
}
