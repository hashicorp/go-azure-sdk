package datasecurityandgovernancesensitivitylabelsublabel

import (
	"context"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UpdateDataSecurityAndGovernanceSensitivityLabelSublabelOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type UpdateDataSecurityAndGovernanceSensitivityLabelSublabelOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultUpdateDataSecurityAndGovernanceSensitivityLabelSublabelOperationOptions() UpdateDataSecurityAndGovernanceSensitivityLabelSublabelOperationOptions {
	return UpdateDataSecurityAndGovernanceSensitivityLabelSublabelOperationOptions{}
}

func (o UpdateDataSecurityAndGovernanceSensitivityLabelSublabelOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o UpdateDataSecurityAndGovernanceSensitivityLabelSublabelOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o UpdateDataSecurityAndGovernanceSensitivityLabelSublabelOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// UpdateDataSecurityAndGovernanceSensitivityLabelSublabel - Update the navigation property sublabels in me
func (c DataSecurityAndGovernanceSensitivityLabelSublabelClient) UpdateDataSecurityAndGovernanceSensitivityLabelSublabel(ctx context.Context, id beta.MeDataSecurityAndGovernanceSensitivityLabelIdSublabelId, input beta.SensitivityLabel, options UpdateDataSecurityAndGovernanceSensitivityLabelSublabelOperationOptions) (result UpdateDataSecurityAndGovernanceSensitivityLabelSublabelOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusAccepted,
			http.StatusNoContent,
			http.StatusOK,
		},
		HttpMethod:    http.MethodPatch,
		OptionsObject: options,
		Path:          id.ID(),
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

	return
}
