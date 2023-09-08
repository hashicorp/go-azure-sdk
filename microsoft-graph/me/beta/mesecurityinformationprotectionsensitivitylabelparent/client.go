package mesecurityinformationprotectionsensitivitylabelparent

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MeSecurityInformationProtectionSensitivityLabelParentClient struct {
	Client *msgraph.Client
}

func NewMeSecurityInformationProtectionSensitivityLabelParentClientWithBaseURI(api sdkEnv.Api) (*MeSecurityInformationProtectionSensitivityLabelParentClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "mesecurityinformationprotectionsensitivitylabelparent", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating MeSecurityInformationProtectionSensitivityLabelParentClient: %+v", err)
	}

	return &MeSecurityInformationProtectionSensitivityLabelParentClient{
		Client: client,
	}, nil
}
