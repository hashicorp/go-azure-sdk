package siteinformationprotectionsensitivitypolicysetting

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeleteSiteInformationProtectionSensitivityPolicySettingOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type DeleteSiteInformationProtectionSensitivityPolicySettingOperationOptions struct {
	IfMatch  *string
	Metadata *odata.Metadata
}

func DefaultDeleteSiteInformationProtectionSensitivityPolicySettingOperationOptions() DeleteSiteInformationProtectionSensitivityPolicySettingOperationOptions {
	return DeleteSiteInformationProtectionSensitivityPolicySettingOperationOptions{}
}

func (o DeleteSiteInformationProtectionSensitivityPolicySettingOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}
	if o.IfMatch != nil {
		out.Append("If-Match", fmt.Sprintf("%v", *o.IfMatch))
	}
	return &out
}

func (o DeleteSiteInformationProtectionSensitivityPolicySettingOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o DeleteSiteInformationProtectionSensitivityPolicySettingOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// DeleteSiteInformationProtectionSensitivityPolicySetting - Delete navigation property sensitivityPolicySettings for
// groups
func (c SiteInformationProtectionSensitivityPolicySettingClient) DeleteSiteInformationProtectionSensitivityPolicySetting(ctx context.Context, id beta.GroupIdSiteId, options DeleteSiteInformationProtectionSensitivityPolicySettingOperationOptions) (result DeleteSiteInformationProtectionSensitivityPolicySettingOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusNoContent,
		},
		HttpMethod:    http.MethodDelete,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/informationProtection/sensitivityPolicySettings", id.ID()),
	}

	req, err := c.Client.NewRequest(ctx, opts)
	if err != nil {
		return
	}

	var resp *client.Response
	resp, err = req.Execute(ctx)
	if resp != nil {
		result.OData = resp.OData
		result.HttpResponse = resp.Response
	}
	if err != nil {
		return
	}

	return
}
