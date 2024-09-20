package siteinformationprotectionpolicylabel

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

type ExtractSiteInformationProtectionPolicyLabelsLabelOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *beta.InformationProtectionContentLabel
}

type ExtractSiteInformationProtectionPolicyLabelsLabelOperationOptions struct {
	Metadata *odata.Metadata
}

func DefaultExtractSiteInformationProtectionPolicyLabelsLabelOperationOptions() ExtractSiteInformationProtectionPolicyLabelsLabelOperationOptions {
	return ExtractSiteInformationProtectionPolicyLabelsLabelOperationOptions{}
}

func (o ExtractSiteInformationProtectionPolicyLabelsLabelOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ExtractSiteInformationProtectionPolicyLabelsLabelOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o ExtractSiteInformationProtectionPolicyLabelsLabelOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// ExtractSiteInformationProtectionPolicyLabelsLabel - Invoke action extractLabel. Using the metadata that exists on an
// already-labeled piece of information, resolve the metadata to a specific sensitivity label. The contentInfo input is
// resolved to informationProtectionContentLabel.
func (c SiteInformationProtectionPolicyLabelClient) ExtractSiteInformationProtectionPolicyLabelsLabel(ctx context.Context, id beta.GroupIdSiteId, input ExtractSiteInformationProtectionPolicyLabelsLabelRequest, options ExtractSiteInformationProtectionPolicyLabelsLabelOperationOptions) (result ExtractSiteInformationProtectionPolicyLabelsLabelOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodPost,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/informationProtection/policy/labels/extractLabel", id.ID()),
	}

	req, err := c.Client.NewRequest(ctx, opts)
	if err != nil {
		return
	}

	if err = req.Marshal(input); err != nil {
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

	var model beta.InformationProtectionContentLabel
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
