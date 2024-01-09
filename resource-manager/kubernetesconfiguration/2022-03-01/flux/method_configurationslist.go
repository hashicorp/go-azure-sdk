package flux

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

type ConfigurationsListOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]FluxConfiguration
}

type ConfigurationsListCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []FluxConfiguration
}

// ConfigurationsList ...
func (c FluxClient) ConfigurationsList(ctx context.Context, id commonids.ScopeId) (result ConfigurationsListOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       fmt.Sprintf("%s/providers/Microsoft.KubernetesConfiguration/fluxConfigurations", id.ID()),
	}

	req, err := c.Client.NewRequest(ctx, opts)
	if err != nil {
		return
	}

	var resp *client.Response
	resp, err = req.ExecutePaged(ctx)
	if resp != nil {
		result.OData = resp.OData
		result.HttpResponse = resp.Response
	}
	if err != nil {
		return
	}

	var values struct {
		Values *[]FluxConfiguration `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ConfigurationsListComplete retrieves all the results into a single object
func (c FluxClient) ConfigurationsListComplete(ctx context.Context, id commonids.ScopeId) (ConfigurationsListCompleteResult, error) {
	return c.ConfigurationsListCompleteMatchingPredicate(ctx, id, FluxConfigurationOperationPredicate{})
}

// ConfigurationsListCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c FluxClient) ConfigurationsListCompleteMatchingPredicate(ctx context.Context, id commonids.ScopeId, predicate FluxConfigurationOperationPredicate) (result ConfigurationsListCompleteResult, err error) {
	items := make([]FluxConfiguration, 0)

	resp, err := c.ConfigurationsList(ctx, id)
	if err != nil {
		err = fmt.Errorf("loading results: %+v", err)
		return
	}
	if resp.Model != nil {
		for _, v := range *resp.Model {
			if predicate.Matches(v) {
				items = append(items, v)
			}
		}
	}

	result = ConfigurationsListCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
