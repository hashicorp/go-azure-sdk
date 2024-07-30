package securityinformationprotectionsensitivitylabel

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityInformationProtectionSensitivityLabelClient struct {
	Client *msgraph.Client
}

func NewSecurityInformationProtectionSensitivityLabelClientWithBaseURI(sdkApi sdkEnv.Api) (*SecurityInformationProtectionSensitivityLabelClient, error) {
	client, err := msgraph.NewClient(sdkApi, "securityinformationprotectionsensitivitylabel", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating SecurityInformationProtectionSensitivityLabelClient: %+v", err)
	}

	return &SecurityInformationProtectionSensitivityLabelClient{
		Client: client,
	}, nil
}
