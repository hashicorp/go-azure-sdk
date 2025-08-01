package siteinformationprotectionsensitivitylabelright

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SiteInformationProtectionSensitivityLabelRightClient struct {
	Client *msgraph.Client
}

func NewSiteInformationProtectionSensitivityLabelRightClientWithBaseURI(sdkApi sdkEnv.Api) (*SiteInformationProtectionSensitivityLabelRightClient, error) {
	client, err := msgraph.NewClient(sdkApi, "siteinformationprotectionsensitivitylabelright", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating SiteInformationProtectionSensitivityLabelRightClient: %+v", err)
	}

	return &SiteInformationProtectionSensitivityLabelRightClient{
		Client: client,
	}, nil
}
