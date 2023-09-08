package groupsiteinformationprotectionsensitivitypolicysetting

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GroupSiteInformationProtectionSensitivityPolicySettingClient struct {
	Client *msgraph.Client
}

func NewGroupSiteInformationProtectionSensitivityPolicySettingClientWithBaseURI(api sdkEnv.Api) (*GroupSiteInformationProtectionSensitivityPolicySettingClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "groupsiteinformationprotectionsensitivitypolicysetting", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating GroupSiteInformationProtectionSensitivityPolicySettingClient: %+v", err)
	}

	return &GroupSiteInformationProtectionSensitivityPolicySettingClient{
		Client: client,
	}, nil
}
