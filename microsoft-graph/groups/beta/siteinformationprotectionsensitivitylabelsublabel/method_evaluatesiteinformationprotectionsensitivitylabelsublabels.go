package siteinformationprotectionsensitivitylabelsublabel

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

type EvaluateSiteInformationProtectionSensitivityLabelSublabelsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *beta.EvaluateLabelJobResponse
}

type EvaluateSiteInformationProtectionSensitivityLabelSublabelsOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultEvaluateSiteInformationProtectionSensitivityLabelSublabelsOperationOptions() EvaluateSiteInformationProtectionSensitivityLabelSublabelsOperationOptions {
	return EvaluateSiteInformationProtectionSensitivityLabelSublabelsOperationOptions{}
}

func (o EvaluateSiteInformationProtectionSensitivityLabelSublabelsOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o EvaluateSiteInformationProtectionSensitivityLabelSublabelsOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o EvaluateSiteInformationProtectionSensitivityLabelSublabelsOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// EvaluateSiteInformationProtectionSensitivityLabelSublabels - Invoke action evaluate
func (c SiteInformationProtectionSensitivityLabelSublabelClient) EvaluateSiteInformationProtectionSensitivityLabelSublabels(ctx context.Context, id beta.GroupIdSiteIdInformationProtectionSensitivityLabelId, input EvaluateSiteInformationProtectionSensitivityLabelSublabelsRequest, options EvaluateSiteInformationProtectionSensitivityLabelSublabelsOperationOptions) (result EvaluateSiteInformationProtectionSensitivityLabelSublabelsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodPost,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/sublabels/evaluate", id.ID()),
		RetryFunc:     options.RetryFunc,
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

	var model beta.EvaluateLabelJobResponse
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
