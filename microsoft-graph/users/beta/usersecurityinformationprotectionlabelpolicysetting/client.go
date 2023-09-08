package usersecurityinformationprotectionlabelpolicysetting

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserSecurityInformationProtectionLabelPolicySettingClient struct {
	Client *msgraph.Client
}

func NewUserSecurityInformationProtectionLabelPolicySettingClientWithBaseURI(api sdkEnv.Api) (*UserSecurityInformationProtectionLabelPolicySettingClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "usersecurityinformationprotectionlabelpolicysetting", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating UserSecurityInformationProtectionLabelPolicySettingClient: %+v", err)
	}

	return &UserSecurityInformationProtectionLabelPolicySettingClient{
		Client: client,
	}, nil
}
