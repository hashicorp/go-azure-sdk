package groupsiteinformationprotectionpolicy

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GroupSiteInformationProtectionPolicyClient struct {
	Client *msgraph.Client
}

func NewGroupSiteInformationProtectionPolicyClientWithBaseURI(api sdkEnv.Api) (*GroupSiteInformationProtectionPolicyClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "groupsiteinformationprotectionpolicy", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating GroupSiteInformationProtectionPolicyClient: %+v", err)
	}

	return &GroupSiteInformationProtectionPolicyClient{
		Client: client,
	}, nil
}
