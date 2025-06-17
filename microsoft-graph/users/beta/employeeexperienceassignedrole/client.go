package employeeexperienceassignedrole

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EmployeeExperienceAssignedRoleClient struct {
	Client *msgraph.Client
}

func NewEmployeeExperienceAssignedRoleClientWithBaseURI(sdkApi sdkEnv.Api) (*EmployeeExperienceAssignedRoleClient, error) {
	client, err := msgraph.NewClient(sdkApi, "employeeexperienceassignedrole", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating EmployeeExperienceAssignedRoleClient: %+v", err)
	}

	return &EmployeeExperienceAssignedRoleClient{
		Client: client,
	}, nil
}
