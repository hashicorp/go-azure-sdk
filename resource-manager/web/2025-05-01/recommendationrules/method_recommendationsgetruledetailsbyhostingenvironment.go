package recommendationrules

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RecommendationsGetRuleDetailsByHostingEnvironmentOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *RecommendationRule
}

type RecommendationsGetRuleDetailsByHostingEnvironmentOperationOptions struct {
	RecommendationId *string
	UpdateSeen       *bool
}

func DefaultRecommendationsGetRuleDetailsByHostingEnvironmentOperationOptions() RecommendationsGetRuleDetailsByHostingEnvironmentOperationOptions {
	return RecommendationsGetRuleDetailsByHostingEnvironmentOperationOptions{}
}

func (o RecommendationsGetRuleDetailsByHostingEnvironmentOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o RecommendationsGetRuleDetailsByHostingEnvironmentOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}

	return &out
}

func (o RecommendationsGetRuleDetailsByHostingEnvironmentOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}
	if o.RecommendationId != nil {
		out.Append("recommendationId", fmt.Sprintf("%v", *o.RecommendationId))
	}
	if o.UpdateSeen != nil {
		out.Append("updateSeen", fmt.Sprintf("%v", *o.UpdateSeen))
	}
	return &out
}

// RecommendationsGetRuleDetailsByHostingEnvironment ...
func (c RecommendationRulesClient) RecommendationsGetRuleDetailsByHostingEnvironment(ctx context.Context, id HostingEnvironmentRecommendationId, options RecommendationsGetRuleDetailsByHostingEnvironmentOperationOptions) (result RecommendationsGetRuleDetailsByHostingEnvironmentOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Path:          id.ID(),
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
