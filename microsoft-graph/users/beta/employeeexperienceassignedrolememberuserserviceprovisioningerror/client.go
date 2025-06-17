package employeeexperienceassignedrolememberuserserviceprovisioningerror

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EmployeeExperienceAssignedRoleMemberUserServiceProvisioningErrorClient struct {
	Client *msgraph.Client
}

func NewEmployeeExperienceAssignedRoleMemberUserServiceProvisioningErrorClientWithBaseURI(sdkApi sdkEnv.Api) (*EmployeeExperienceAssignedRoleMemberUserServiceProvisioningErrorClient, error) {
	client, err := msgraph.NewClient(sdkApi, "employeeexperienceassignedrolememberuserserviceprovisioningerror", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating EmployeeExperienceAssignedRoleMemberUserServiceProvisioningErrorClient: %+v", err)
	}

	return &EmployeeExperienceAssignedRoleMemberUserServiceProvisioningErrorClient{
		Client: client,
	}, nil
}
