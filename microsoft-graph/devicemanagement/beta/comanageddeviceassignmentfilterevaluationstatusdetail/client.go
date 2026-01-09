package comanageddeviceassignmentfilterevaluationstatusdetail

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ComanagedDeviceAssignmentFilterEvaluationStatusDetailClient struct {
	Client *msgraph.Client
}

func NewComanagedDeviceAssignmentFilterEvaluationStatusDetailClientWithBaseURI(sdkApi sdkEnv.Api) (*ComanagedDeviceAssignmentFilterEvaluationStatusDetailClient, error) {
	client, err := msgraph.NewClient(sdkApi, "comanageddeviceassignmentfilterevaluationstatusdetail", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating ComanagedDeviceAssignmentFilterEvaluationStatusDetailClient: %+v", err)
	}

	return &ComanagedDeviceAssignmentFilterEvaluationStatusDetailClient{
		Client: client,
	}, nil
}
