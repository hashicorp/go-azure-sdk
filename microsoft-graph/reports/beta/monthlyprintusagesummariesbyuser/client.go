package monthlyprintusagesummariesbyuser

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MonthlyPrintUsageSummariesByUserClient struct {
	Client *msgraph.Client
}

func NewMonthlyPrintUsageSummariesByUserClientWithBaseURI(sdkApi sdkEnv.Api) (*MonthlyPrintUsageSummariesByUserClient, error) {
	client, err := msgraph.NewClient(sdkApi, "monthlyprintusagesummariesbyuser", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating MonthlyPrintUsageSummariesByUserClient: %+v", err)
	}

	return &MonthlyPrintUsageSummariesByUserClient{
		Client: client,
	}, nil
}
