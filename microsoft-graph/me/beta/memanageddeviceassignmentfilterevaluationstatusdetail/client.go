package memanageddeviceassignmentfilterevaluationstatusdetail

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MeManagedDeviceAssignmentFilterEvaluationStatusDetailClient struct {
	Client *msgraph.Client
}

func NewMeManagedDeviceAssignmentFilterEvaluationStatusDetailClientWithBaseURI(api sdkEnv.Api) (*MeManagedDeviceAssignmentFilterEvaluationStatusDetailClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "memanageddeviceassignmentfilterevaluationstatusdetail", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating MeManagedDeviceAssignmentFilterEvaluationStatusDetailClient: %+v", err)
	}

	return &MeManagedDeviceAssignmentFilterEvaluationStatusDetailClient{
		Client: client,
	}, nil
}
