package lifecycleworkflowtaskreport

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type LifecycleWorkflowTaskReportClient struct {
	Client *msgraph.Client
}

func NewLifecycleWorkflowTaskReportClientWithBaseURI(sdkApi sdkEnv.Api) (*LifecycleWorkflowTaskReportClient, error) {
	client, err := msgraph.NewClient(sdkApi, "lifecycleworkflowtaskreport", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating LifecycleWorkflowTaskReportClient: %+v", err)
	}

	return &LifecycleWorkflowTaskReportClient{
		Client: client,
	}, nil
}
