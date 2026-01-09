package macossoftwareupdateaccountsummarycategorysummary

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MacOSSoftwareUpdateAccountSummaryCategorySummaryClient struct {
	Client *msgraph.Client
}

func NewMacOSSoftwareUpdateAccountSummaryCategorySummaryClientWithBaseURI(sdkApi sdkEnv.Api) (*MacOSSoftwareUpdateAccountSummaryCategorySummaryClient, error) {
	client, err := msgraph.NewClient(sdkApi, "macossoftwareupdateaccountsummarycategorysummary", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating MacOSSoftwareUpdateAccountSummaryCategorySummaryClient: %+v", err)
	}

	return &MacOSSoftwareUpdateAccountSummaryCategorySummaryClient{
		Client: client,
	}, nil
}
