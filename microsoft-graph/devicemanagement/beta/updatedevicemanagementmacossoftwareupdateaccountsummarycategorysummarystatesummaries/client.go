package updatedevicemanagementmacossoftwareupdateaccountsummarycategorysummarystatesummaries

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UpdateDeviceManagementMacOSSoftwareUpdateAccountSummaryCategorySummaryStateSummariesClient struct {
	Client *msgraph.Client
}

func NewUpdateDeviceManagementMacOSSoftwareUpdateAccountSummaryCategorySummaryStateSummariesClientWithBaseURI(sdkApi sdkEnv.Api) (*UpdateDeviceManagementMacOSSoftwareUpdateAccountSummaryCategorySummaryStateSummariesClient, error) {
	client, err := msgraph.NewClient(sdkApi, "updatedevicemanagementmacossoftwareupdateaccountsummarycategorysummarystatesummaries", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating UpdateDeviceManagementMacOSSoftwareUpdateAccountSummaryCategorySummaryStateSummariesClient: %+v", err)
	}

	return &UpdateDeviceManagementMacOSSoftwareUpdateAccountSummaryCategorySummaryStateSummariesClient{
		Client: client,
	}, nil
}
