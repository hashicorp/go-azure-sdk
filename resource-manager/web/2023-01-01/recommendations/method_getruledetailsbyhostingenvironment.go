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

type GetRuleDetailsByHostingEnvironmentOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *RecommendationRule
}

type GetRuleDetailsByHostingEnvironmentOperationOptions struct {
	RecommendationId *string
	UpdateSeen       *bool
}

func DefaultGetRuleDetailsByHostingEnvironmentOperationOptions() GetRuleDetailsByHostingEnvironmentOperationOptions {
	return GetRuleDetailsByHostingEnvironmentOperationOptions{}
}

func (o GetRuleDetailsByHostingEnvironmentOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o GetRuleDetailsByHostingEnvironmentOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	return &out
}

func (o GetRuleDetailsByHostingEnvironmentOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}
	if o.RecommendationId != nil {
		out.Append("recommendationId", fmt.Sprintf("%v", *o.RecommendationId))
	}
	if o.UpdateSeen != nil {
		out.Append("updateSeen", fmt.Sprintf("%v", *o.UpdateSeen))
	}
	return &out
}

// GetRuleDetailsByHostingEnvironment ...
func (c RecommendationsClient) GetRuleDetailsByHostingEnvironment(ctx context.Context, id HostingEnvironmentRecommendationId, options GetRuleDetailsByHostingEnvironmentOperationOptions) (result GetRuleDetailsByHostingEnvironmentOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
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

	if err = resp.Unmarshal(&result.Model); err != nil {
		return
	}

	return
}
