package groupsiteinformationprotectiondatalosspreventionpolicy

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GroupSiteInformationProtectionDataLossPreventionPolicyClient struct {
	Client *msgraph.Client
}

func NewGroupSiteInformationProtectionDataLossPreventionPolicyClientWithBaseURI(api sdkEnv.Api) (*GroupSiteInformationProtectionDataLossPreventionPolicyClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "groupsiteinformationprotectiondatalosspreventionpolicy", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating GroupSiteInformationProtectionDataLossPreventionPolicyClient: %+v", err)
	}

	return &GroupSiteInformationProtectionDataLossPreventionPolicyClient{
		Client: client,
	}, nil
}
