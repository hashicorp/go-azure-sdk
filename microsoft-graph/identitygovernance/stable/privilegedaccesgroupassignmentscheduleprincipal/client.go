package privilegedaccesgroupassignmentscheduleprincipal

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PrivilegedAccesGroupAssignmentSchedulePrincipalClient struct {
	Client *msgraph.Client
}

func NewPrivilegedAccesGroupAssignmentSchedulePrincipalClientWithBaseURI(sdkApi sdkEnv.Api) (*PrivilegedAccesGroupAssignmentSchedulePrincipalClient, error) {
	client, err := msgraph.NewClient(sdkApi, "privilegedaccesgroupassignmentscheduleprincipal", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating PrivilegedAccesGroupAssignmentSchedulePrincipalClient: %+v", err)
	}

	return &PrivilegedAccesGroupAssignmentSchedulePrincipalClient{
		Client: client,
	}, nil
}
