package lifecycleworkflowcustomtaskextensionlastmodifiedby

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type LifecycleWorkflowCustomTaskExtensionLastModifiedByClient struct {
	Client *msgraph.Client
}

func NewLifecycleWorkflowCustomTaskExtensionLastModifiedByClientWithBaseURI(sdkApi sdkEnv.Api) (*LifecycleWorkflowCustomTaskExtensionLastModifiedByClient, error) {
	client, err := msgraph.NewClient(sdkApi, "lifecycleworkflowcustomtaskextensionlastmodifiedby", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating LifecycleWorkflowCustomTaskExtensionLastModifiedByClient: %+v", err)
	}

	return &LifecycleWorkflowCustomTaskExtensionLastModifiedByClient{
		Client: client,
	}, nil
}
