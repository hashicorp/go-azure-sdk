package meinformationprotectionsensitivitylabelsublabel

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MeInformationProtectionSensitivityLabelSublabelClient struct {
	Client *msgraph.Client
}

func NewMeInformationProtectionSensitivityLabelSublabelClientWithBaseURI(api sdkEnv.Api) (*MeInformationProtectionSensitivityLabelSublabelClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "meinformationprotectionsensitivitylabelsublabel", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating MeInformationProtectionSensitivityLabelSublabelClient: %+v", err)
	}

	return &MeInformationProtectionSensitivityLabelSublabelClient{
		Client: client,
	}, nil
}
