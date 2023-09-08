package mesecurityinformationprotectionsensitivitylabel

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MeSecurityInformationProtectionSensitivityLabelClient struct {
	Client *msgraph.Client
}

func NewMeSecurityInformationProtectionSensitivityLabelClientWithBaseURI(api sdkEnv.Api) (*MeSecurityInformationProtectionSensitivityLabelClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "mesecurityinformationprotectionsensitivitylabel", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating MeSecurityInformationProtectionSensitivityLabelClient: %+v", err)
	}

	return &MeSecurityInformationProtectionSensitivityLabelClient{
		Client: client,
	}, nil
}
