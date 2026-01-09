package privilegedaccessgroupassignmentscheduleprincipal

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PrivilegedAccessGroupAssignmentSchedulePrincipalClient struct {
	Client *msgraph.Client
}

func NewPrivilegedAccessGroupAssignmentSchedulePrincipalClientWithBaseURI(sdkApi sdkEnv.Api) (*PrivilegedAccessGroupAssignmentSchedulePrincipalClient, error) {
	client, err := msgraph.NewClient(sdkApi, "privilegedaccessgroupassignmentscheduleprincipal", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating PrivilegedAccessGroupAssignmentSchedulePrincipalClient: %+v", err)
	}

	return &PrivilegedAccessGroupAssignmentSchedulePrincipalClient{
		Client: client,
	}, nil
}
