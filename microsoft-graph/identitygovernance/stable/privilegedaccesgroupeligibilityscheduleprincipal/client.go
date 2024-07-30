package privilegedaccesgroupeligibilityscheduleprincipal

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PrivilegedAccesGroupEligibilitySchedulePrincipalClient struct {
	Client *msgraph.Client
}

func NewPrivilegedAccesGroupEligibilitySchedulePrincipalClientWithBaseURI(sdkApi sdkEnv.Api) (*PrivilegedAccesGroupEligibilitySchedulePrincipalClient, error) {
	client, err := msgraph.NewClient(sdkApi, "privilegedaccesgroupeligibilityscheduleprincipal", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating PrivilegedAccesGroupEligibilitySchedulePrincipalClient: %+v", err)
	}

	return &PrivilegedAccesGroupEligibilitySchedulePrincipalClient{
		Client: client,
	}, nil
}
