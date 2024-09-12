package monthlyprintusagebyuser

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MonthlyPrintUsageByUserClient struct {
	Client *msgraph.Client
}

func NewMonthlyPrintUsageByUserClientWithBaseURI(sdkApi sdkEnv.Api) (*MonthlyPrintUsageByUserClient, error) {
	client, err := msgraph.NewClient(sdkApi, "monthlyprintusagebyuser", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating MonthlyPrintUsageByUserClient: %+v", err)
	}

	return &MonthlyPrintUsageByUserClient{
		Client: client,
	}, nil
}
