package partnerbillingusage

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PartnerBillingUsageClient struct {
	Client *msgraph.Client
}

func NewPartnerBillingUsageClientWithBaseURI(sdkApi sdkEnv.Api) (*PartnerBillingUsageClient, error) {
	client, err := msgraph.NewClient(sdkApi, "partnerbillingusage", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating PartnerBillingUsageClient: %+v", err)
	}

	return &PartnerBillingUsageClient{
		Client: client,
	}, nil
}
