package privilegedaccessgroupassignmentschedulegroupserviceprovisioningerror

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PrivilegedAccessGroupAssignmentScheduleGroupServiceProvisioningErrorClient struct {
	Client *msgraph.Client
}

func NewPrivilegedAccessGroupAssignmentScheduleGroupServiceProvisioningErrorClientWithBaseURI(sdkApi sdkEnv.Api) (*PrivilegedAccessGroupAssignmentScheduleGroupServiceProvisioningErrorClient, error) {
	client, err := msgraph.NewClient(sdkApi, "privilegedaccessgroupassignmentschedulegroupserviceprovisioningerror", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating PrivilegedAccessGroupAssignmentScheduleGroupServiceProvisioningErrorClient: %+v", err)
	}

	return &PrivilegedAccessGroupAssignmentScheduleGroupServiceProvisioningErrorClient{
		Client: client,
	}, nil
}
