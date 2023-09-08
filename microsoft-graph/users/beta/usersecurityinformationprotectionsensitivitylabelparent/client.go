package usersecurityinformationprotectionsensitivitylabelparent

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserSecurityInformationProtectionSensitivityLabelParentClient struct {
	Client *msgraph.Client
}

func NewUserSecurityInformationProtectionSensitivityLabelParentClientWithBaseURI(api sdkEnv.Api) (*UserSecurityInformationProtectionSensitivityLabelParentClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "usersecurityinformationprotectionsensitivitylabelparent", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating UserSecurityInformationProtectionSensitivityLabelParentClient: %+v", err)
	}

	return &UserSecurityInformationProtectionSensitivityLabelParentClient{
		Client: client,
	}, nil
}
