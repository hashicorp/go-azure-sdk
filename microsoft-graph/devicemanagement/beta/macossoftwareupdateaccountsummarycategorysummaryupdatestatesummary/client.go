package macossoftwareupdateaccountsummarycategorysummaryupdatestatesummary

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MacOSSoftwareUpdateAccountSummaryCategorySummaryUpdateStateSummaryClient struct {
	Client *msgraph.Client
}

func NewMacOSSoftwareUpdateAccountSummaryCategorySummaryUpdateStateSummaryClientWithBaseURI(sdkApi sdkEnv.Api) (*MacOSSoftwareUpdateAccountSummaryCategorySummaryUpdateStateSummaryClient, error) {
	client, err := msgraph.NewClient(sdkApi, "macossoftwareupdateaccountsummarycategorysummaryupdatestatesummary", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating MacOSSoftwareUpdateAccountSummaryCategorySummaryUpdateStateSummaryClient: %+v", err)
	}

	return &MacOSSoftwareUpdateAccountSummaryCategorySummaryUpdateStateSummaryClient{
		Client: client,
	}, nil
}
