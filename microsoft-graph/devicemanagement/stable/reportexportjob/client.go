package reportexportjob

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ReportExportJobClient struct {
	Client *msgraph.Client
}

func NewReportExportJobClientWithBaseURI(sdkApi sdkEnv.Api) (*ReportExportJobClient, error) {
	client, err := msgraph.NewClient(sdkApi, "reportexportjob", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating ReportExportJobClient: %+v", err)
	}

	return &ReportExportJobClient{
		Client: client,
	}, nil
}
