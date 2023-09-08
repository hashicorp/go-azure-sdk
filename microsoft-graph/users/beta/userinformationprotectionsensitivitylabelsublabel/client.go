package userinformationprotectionsensitivitylabelsublabel

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserInformationProtectionSensitivityLabelSublabelClient struct {
	Client *msgraph.Client
}

func NewUserInformationProtectionSensitivityLabelSublabelClientWithBaseURI(api sdkEnv.Api) (*UserInformationProtectionSensitivityLabelSublabelClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "userinformationprotectionsensitivitylabelsublabel", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating UserInformationProtectionSensitivityLabelSublabelClient: %+v", err)
	}

	return &UserInformationProtectionSensitivityLabelSublabelClient{
		Client: client,
	}, nil
}
