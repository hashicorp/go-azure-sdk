package privilegedaccesgroupeligibilityscheduleinstanceprincipal

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PrivilegedAccesGroupEligibilityScheduleInstancePrincipalClient struct {
	Client *msgraph.Client
}

func NewPrivilegedAccesGroupEligibilityScheduleInstancePrincipalClientWithBaseURI(sdkApi sdkEnv.Api) (*PrivilegedAccesGroupEligibilityScheduleInstancePrincipalClient, error) {
	client, err := msgraph.NewClient(sdkApi, "privilegedaccesgroupeligibilityscheduleinstanceprincipal", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating PrivilegedAccesGroupEligibilityScheduleInstancePrincipalClient: %+v", err)
	}

	return &PrivilegedAccesGroupEligibilityScheduleInstancePrincipalClient{
		Client: client,
	}, nil
}
