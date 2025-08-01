package siteinformationprotectionsensitivitylabelsublabelright

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SiteInformationProtectionSensitivityLabelSublabelRightClient struct {
	Client *msgraph.Client
}

func NewSiteInformationProtectionSensitivityLabelSublabelRightClientWithBaseURI(sdkApi sdkEnv.Api) (*SiteInformationProtectionSensitivityLabelSublabelRightClient, error) {
	client, err := msgraph.NewClient(sdkApi, "siteinformationprotectionsensitivitylabelsublabelright", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating SiteInformationProtectionSensitivityLabelSublabelRightClient: %+v", err)
	}

	return &SiteInformationProtectionSensitivityLabelSublabelRightClient{
		Client: client,
	}, nil
}
