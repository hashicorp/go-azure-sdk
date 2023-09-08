package userinformationprotectionsensitivitylabel

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserInformationProtectionSensitivityLabelClient struct {
	Client *msgraph.Client
}

func NewUserInformationProtectionSensitivityLabelClientWithBaseURI(api sdkEnv.Api) (*UserInformationProtectionSensitivityLabelClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "userinformationprotectionsensitivitylabel", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating UserInformationProtectionSensitivityLabelClient: %+v", err)
	}

	return &UserInformationProtectionSensitivityLabelClient{
		Client: client,
	}, nil
}
