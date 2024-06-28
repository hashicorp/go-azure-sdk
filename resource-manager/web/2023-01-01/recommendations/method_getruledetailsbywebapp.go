package recommendations

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GetRuleDetailsByWebAppOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *RecommendationRule
}

type GetRuleDetailsByWebAppOperationOptions struct {
	RecommendationId *string
	UpdateSeen       *bool
}

func DefaultGetRuleDetailsByWebAppOperationOptions() GetRuleDetailsByWebAppOperationOptions {
	return GetRuleDetailsByWebAppOperationOptions{}
}

func (o GetRuleDetailsByWebAppOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o GetRuleDetailsByWebAppOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	return &out
}

func (o GetRuleDetailsByWebAppOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}
	if o.RecommendationId != nil {
		out.Append("recommendationId", fmt.Sprintf("%v", *o.RecommendationId))
	}
	if o.UpdateSeen != nil {
		out.Append("updateSeen", fmt.Sprintf("%v", *o.UpdateSeen))
	}
	return &out
}

// GetRuleDetailsByWebApp ...
func (c RecommendationsClient) GetRuleDetailsByWebApp(ctx context.Context, id SiteRecommendationId, options GetRuleDetailsByWebAppOperationOptions) (result GetRuleDetailsByWebAppOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,

		Path:          id.ID(),
		OptionsObject: options,
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

	var model RecommendationRule
	result.Model = &model

	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
