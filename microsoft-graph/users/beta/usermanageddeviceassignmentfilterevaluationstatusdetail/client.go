package usermanageddeviceassignmentfilterevaluationstatusdetail

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserManagedDeviceAssignmentFilterEvaluationStatusDetailClient struct {
	Client *msgraph.Client
}

func NewUserManagedDeviceAssignmentFilterEvaluationStatusDetailClientWithBaseURI(api sdkEnv.Api) (*UserManagedDeviceAssignmentFilterEvaluationStatusDetailClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "usermanageddeviceassignmentfilterevaluationstatusdetail", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating UserManagedDeviceAssignmentFilterEvaluationStatusDetailClient: %+v", err)
	}

	return &UserManagedDeviceAssignmentFilterEvaluationStatusDetailClient{
		Client: client,
	}, nil
}
