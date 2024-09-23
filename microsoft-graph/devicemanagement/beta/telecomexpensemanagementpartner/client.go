package telecomexpensemanagementpartner

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TelecomExpenseManagementPartnerClient struct {
	Client *msgraph.Client
}

func NewTelecomExpenseManagementPartnerClientWithBaseURI(sdkApi sdkEnv.Api) (*TelecomExpenseManagementPartnerClient, error) {
	client, err := msgraph.NewClient(sdkApi, "telecomexpensemanagementpartner", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating TelecomExpenseManagementPartnerClient: %+v", err)
	}

	return &TelecomExpenseManagementPartnerClient{
		Client: client,
	}, nil
}
