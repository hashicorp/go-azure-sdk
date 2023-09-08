package groupsiteinformationprotectionsensitivitylabelsublabel

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GroupSiteInformationProtectionSensitivityLabelSublabelClient struct {
	Client *msgraph.Client
}

func NewGroupSiteInformationProtectionSensitivityLabelSublabelClientWithBaseURI(api sdkEnv.Api) (*GroupSiteInformationProtectionSensitivityLabelSublabelClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "groupsiteinformationprotectionsensitivitylabelsublabel", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating GroupSiteInformationProtectionSensitivityLabelSublabelClient: %+v", err)
	}

	return &GroupSiteInformationProtectionSensitivityLabelSublabelClient{
		Client: client,
	}, nil
}
