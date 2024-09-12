package lifecycleworkflowuserprocessingresulttasksubject

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type LifecycleWorkflowUserProcessingResultTaskSubjectClient struct {
	Client *msgraph.Client
}

func NewLifecycleWorkflowUserProcessingResultTaskSubjectClientWithBaseURI(sdkApi sdkEnv.Api) (*LifecycleWorkflowUserProcessingResultTaskSubjectClient, error) {
	client, err := msgraph.NewClient(sdkApi, "lifecycleworkflowuserprocessingresulttasksubject", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating LifecycleWorkflowUserProcessingResultTaskSubjectClient: %+v", err)
	}

	return &LifecycleWorkflowUserProcessingResultTaskSubjectClient{
		Client: client,
	}, nil
}
