package manageddeviceassignmentfilterevaluationstatusdetail

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ManagedDeviceAssignmentFilterEvaluationStatusDetailClient struct {
	Client *msgraph.Client
}

func NewManagedDeviceAssignmentFilterEvaluationStatusDetailClientWithBaseURI(sdkApi sdkEnv.Api) (*ManagedDeviceAssignmentFilterEvaluationStatusDetailClient, error) {
	client, err := msgraph.NewClient(sdkApi, "manageddeviceassignmentfilterevaluationstatusdetail", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating ManagedDeviceAssignmentFilterEvaluationStatusDetailClient: %+v", err)
	}

	return &ManagedDeviceAssignmentFilterEvaluationStatusDetailClient{
		Client: client,
	}, nil
}
