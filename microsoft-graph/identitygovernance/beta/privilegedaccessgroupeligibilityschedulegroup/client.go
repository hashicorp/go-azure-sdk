package privilegedaccessgroupeligibilityschedulegroup

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PrivilegedAccessGroupEligibilityScheduleGroupClient struct {
	Client *msgraph.Client
}

func NewPrivilegedAccessGroupEligibilityScheduleGroupClientWithBaseURI(sdkApi sdkEnv.Api) (*PrivilegedAccessGroupEligibilityScheduleGroupClient, error) {
	client, err := msgraph.NewClient(sdkApi, "privilegedaccessgroupeligibilityschedulegroup", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating PrivilegedAccessGroupEligibilityScheduleGroupClient: %+v", err)
	}

	return &PrivilegedAccessGroupEligibilityScheduleGroupClient{
		Client: client,
	}, nil
}
