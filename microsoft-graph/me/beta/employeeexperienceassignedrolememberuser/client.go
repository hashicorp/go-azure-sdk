package employeeexperienceassignedrolememberuser

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EmployeeExperienceAssignedRoleMemberUserClient struct {
	Client *msgraph.Client
}

func NewEmployeeExperienceAssignedRoleMemberUserClientWithBaseURI(sdkApi sdkEnv.Api) (*EmployeeExperienceAssignedRoleMemberUserClient, error) {
	client, err := msgraph.NewClient(sdkApi, "employeeexperienceassignedrolememberuser", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating EmployeeExperienceAssignedRoleMemberUserClient: %+v", err)
	}

	return &EmployeeExperienceAssignedRoleMemberUserClient{
		Client: client,
	}, nil
}
