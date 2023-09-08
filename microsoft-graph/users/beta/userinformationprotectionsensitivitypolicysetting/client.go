package userinformationprotectionsensitivitypolicysetting

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserInformationProtectionSensitivityPolicySettingClient struct {
	Client *msgraph.Client
}

func NewUserInformationProtectionSensitivityPolicySettingClientWithBaseURI(api sdkEnv.Api) (*UserInformationProtectionSensitivityPolicySettingClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "userinformationprotectionsensitivitypolicysetting", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating UserInformationProtectionSensitivityPolicySettingClient: %+v", err)
	}

	return &UserInformationProtectionSensitivityPolicySettingClient{
		Client: client,
	}, nil
}
