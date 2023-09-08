package usersecurityinformationprotectionsensitivitylabel

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserSecurityInformationProtectionSensitivityLabelClient struct {
	Client *msgraph.Client
}

func NewUserSecurityInformationProtectionSensitivityLabelClientWithBaseURI(api sdkEnv.Api) (*UserSecurityInformationProtectionSensitivityLabelClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "usersecurityinformationprotectionsensitivitylabel", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating UserSecurityInformationProtectionSensitivityLabelClient: %+v", err)
	}

	return &UserSecurityInformationProtectionSensitivityLabelClient{
		Client: client,
	}, nil
}
