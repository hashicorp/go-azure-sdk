package lifecycleworkflowdeleteditemworkflowtaskreport

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type LifecycleWorkflowDeletedItemWorkflowTaskReportClient struct {
	Client *msgraph.Client
}

func NewLifecycleWorkflowDeletedItemWorkflowTaskReportClientWithBaseURI(sdkApi sdkEnv.Api) (*LifecycleWorkflowDeletedItemWorkflowTaskReportClient, error) {
	client, err := msgraph.NewClient(sdkApi, "lifecycleworkflowdeleteditemworkflowtaskreport", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating LifecycleWorkflowDeletedItemWorkflowTaskReportClient: %+v", err)
	}

	return &LifecycleWorkflowDeletedItemWorkflowTaskReportClient{
		Client: client,
	}, nil
}
