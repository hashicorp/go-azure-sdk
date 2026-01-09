package privilegedaccessgroupeligibilityscheduleprincipal

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PrivilegedAccessGroupEligibilitySchedulePrincipalClient struct {
	Client *msgraph.Client
}

func NewPrivilegedAccessGroupEligibilitySchedulePrincipalClientWithBaseURI(sdkApi sdkEnv.Api) (*PrivilegedAccessGroupEligibilitySchedulePrincipalClient, error) {
	client, err := msgraph.NewClient(sdkApi, "privilegedaccessgroupeligibilityscheduleprincipal", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating PrivilegedAccessGroupEligibilitySchedulePrincipalClient: %+v", err)
	}

	return &PrivilegedAccessGroupEligibilitySchedulePrincipalClient{
		Client: client,
	}, nil
}
