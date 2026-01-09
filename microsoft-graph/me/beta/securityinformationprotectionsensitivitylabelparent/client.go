package securityinformationprotectionsensitivitylabelparent

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityInformationProtectionSensitivityLabelParentClient struct {
	Client *msgraph.Client
}

func NewSecurityInformationProtectionSensitivityLabelParentClientWithBaseURI(sdkApi sdkEnv.Api) (*SecurityInformationProtectionSensitivityLabelParentClient, error) {
	client, err := msgraph.NewClient(sdkApi, "securityinformationprotectionsensitivitylabelparent", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating SecurityInformationProtectionSensitivityLabelParentClient: %+v", err)
	}

	return &SecurityInformationProtectionSensitivityLabelParentClient{
		Client: client,
	}, nil
}
