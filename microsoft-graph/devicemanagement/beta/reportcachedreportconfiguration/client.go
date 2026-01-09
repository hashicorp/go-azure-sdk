package reportcachedreportconfiguration

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ReportCachedReportConfigurationClient struct {
	Client *msgraph.Client
}

func NewReportCachedReportConfigurationClientWithBaseURI(sdkApi sdkEnv.Api) (*ReportCachedReportConfigurationClient, error) {
	client, err := msgraph.NewClient(sdkApi, "reportcachedreportconfiguration", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating ReportCachedReportConfigurationClient: %+v", err)
	}

	return &ReportCachedReportConfigurationClient{
		Client: client,
	}, nil
}
