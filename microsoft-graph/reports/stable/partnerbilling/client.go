package partnerbilling

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PartnerBillingClient struct {
	Client *msgraph.Client
}

func NewPartnerBillingClientWithBaseURI(sdkApi sdkEnv.Api) (*PartnerBillingClient, error) {
	client, err := msgraph.NewClient(sdkApi, "partnerbilling", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating PartnerBillingClient: %+v", err)
	}

	return &PartnerBillingClient{
		Client: client,
	}, nil
}
