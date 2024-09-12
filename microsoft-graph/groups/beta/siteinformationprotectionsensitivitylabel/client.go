package siteinformationprotectionsensitivitylabel

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SiteInformationProtectionSensitivityLabelClient struct {
	Client *msgraph.Client
}

func NewSiteInformationProtectionSensitivityLabelClientWithBaseURI(sdkApi sdkEnv.Api) (*SiteInformationProtectionSensitivityLabelClient, error) {
	client, err := msgraph.NewClient(sdkApi, "siteinformationprotectionsensitivitylabel", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating SiteInformationProtectionSensitivityLabelClient: %+v", err)
	}

	return &SiteInformationProtectionSensitivityLabelClient{
		Client: client,
	}, nil
}
