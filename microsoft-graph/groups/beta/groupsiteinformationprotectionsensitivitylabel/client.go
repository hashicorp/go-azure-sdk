package groupsiteinformationprotectionsensitivitylabel

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GroupSiteInformationProtectionSensitivityLabelClient struct {
	Client *msgraph.Client
}

func NewGroupSiteInformationProtectionSensitivityLabelClientWithBaseURI(api sdkEnv.Api) (*GroupSiteInformationProtectionSensitivityLabelClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "groupsiteinformationprotectionsensitivitylabel", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating GroupSiteInformationProtectionSensitivityLabelClient: %+v", err)
	}

	return &GroupSiteInformationProtectionSensitivityLabelClient{
		Client: client,
	}, nil
}
