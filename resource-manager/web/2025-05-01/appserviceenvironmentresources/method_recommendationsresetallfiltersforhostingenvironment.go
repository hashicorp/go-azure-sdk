package appserviceenvironmentresources

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/commonids"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RecommendationsResetAllFiltersForHostingEnvironmentOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type RecommendationsResetAllFiltersForHostingEnvironmentOperationOptions struct {
	EnvironmentName *string
}

func DefaultRecommendationsResetAllFiltersForHostingEnvironmentOperationOptions() RecommendationsResetAllFiltersForHostingEnvironmentOperationOptions {
	return RecommendationsResetAllFiltersForHostingEnvironmentOperationOptions{}
}

func (o RecommendationsResetAllFiltersForHostingEnvironmentOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o RecommendationsResetAllFiltersForHostingEnvironmentOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}

	return &out
}

func (o RecommendationsResetAllFiltersForHostingEnvironmentOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}
	if o.EnvironmentName != nil {
		out.Append("environmentName", fmt.Sprintf("%v", *o.EnvironmentName))
	}
	return &out
}

// RecommendationsResetAllFiltersForHostingEnvironment ...
func (c AppServiceEnvironmentResourcesClient) RecommendationsResetAllFiltersForHostingEnvironment(ctx context.Context, id commonids.AppServiceEnvironmentId, options RecommendationsResetAllFiltersForHostingEnvironmentOperationOptions) (result RecommendationsResetAllFiltersForHostingEnvironmentOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusNoContent,
		},
		HttpMethod:    http.MethodPost,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/recommendations/reset", id.ID()),
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
