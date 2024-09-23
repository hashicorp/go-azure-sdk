package partnerbillingoperation

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PartnerBillingOperationClient struct {
	Client *msgraph.Client
}

func NewPartnerBillingOperationClientWithBaseURI(sdkApi sdkEnv.Api) (*PartnerBillingOperationClient, error) {
	client, err := msgraph.NewClient(sdkApi, "partnerbillingoperation", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating PartnerBillingOperationClient: %+v", err)
	}

	return &PartnerBillingOperationClient{
		Client: client,
	}, nil
}
