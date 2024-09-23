package securityinformationprotectionlabelpolicysetting

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityInformationProtectionLabelPolicySettingClient struct {
	Client *msgraph.Client
}

func NewSecurityInformationProtectionLabelPolicySettingClientWithBaseURI(sdkApi sdkEnv.Api) (*SecurityInformationProtectionLabelPolicySettingClient, error) {
	client, err := msgraph.NewClient(sdkApi, "securityinformationprotectionlabelpolicysetting", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating SecurityInformationProtectionLabelPolicySettingClient: %+v", err)
	}

	return &SecurityInformationProtectionLabelPolicySettingClient{
		Client: client,
	}, nil
}
