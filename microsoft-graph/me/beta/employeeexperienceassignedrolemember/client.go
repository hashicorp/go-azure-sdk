package employeeexperienceassignedrolemember

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EmployeeExperienceAssignedRoleMemberClient struct {
	Client *msgraph.Client
}

func NewEmployeeExperienceAssignedRoleMemberClientWithBaseURI(sdkApi sdkEnv.Api) (*EmployeeExperienceAssignedRoleMemberClient, error) {
	client, err := msgraph.NewClient(sdkApi, "employeeexperienceassignedrolemember", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating EmployeeExperienceAssignedRoleMemberClient: %+v", err)
	}

	return &EmployeeExperienceAssignedRoleMemberClient{
		Client: client,
	}, nil
}
