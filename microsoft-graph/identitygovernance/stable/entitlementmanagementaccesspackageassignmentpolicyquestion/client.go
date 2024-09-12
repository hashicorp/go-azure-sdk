package entitlementmanagementaccesspackageassignmentpolicyquestion

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EntitlementManagementAccessPackageAssignmentPolicyQuestionClient struct {
	Client *msgraph.Client
}

func NewEntitlementManagementAccessPackageAssignmentPolicyQuestionClientWithBaseURI(sdkApi sdkEnv.Api) (*EntitlementManagementAccessPackageAssignmentPolicyQuestionClient, error) {
	client, err := msgraph.NewClient(sdkApi, "entitlementmanagementaccesspackageassignmentpolicyquestion", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating EntitlementManagementAccessPackageAssignmentPolicyQuestionClient: %+v", err)
	}

	return &EntitlementManagementAccessPackageAssignmentPolicyQuestionClient{
		Client: client,
	}, nil
}
