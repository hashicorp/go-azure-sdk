package groupsiteinformationprotection

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GroupSiteInformationProtectionClient struct {
	Client *msgraph.Client
}

func NewGroupSiteInformationProtectionClientWithBaseURI(api sdkEnv.Api) (*GroupSiteInformationProtectionClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "groupsiteinformationprotection", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating GroupSiteInformationProtectionClient: %+v", err)
	}

	return &GroupSiteInformationProtectionClient{
		Client: client,
	}, nil
}
