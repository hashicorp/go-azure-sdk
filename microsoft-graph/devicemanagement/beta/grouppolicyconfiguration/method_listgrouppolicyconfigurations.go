package grouppolicyconfiguration

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

type ListGroupPolicyConfigurationsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.GroupPolicyConfiguration
}

type ListGroupPolicyConfigurationsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.GroupPolicyConfiguration
}

type ListGroupPolicyConfigurationsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListGroupPolicyConfigurationsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListGroupPolicyConfigurations ...
func (c GroupPolicyConfigurationClient) ListGroupPolicyConfigurations(ctx context.Context) (result ListGroupPolicyConfigurationsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListGroupPolicyConfigurationsCustomPager{},
		Path:       "/deviceManagement/groupPolicyConfigurations",
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
		Values *[]beta.GroupPolicyConfiguration `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListGroupPolicyConfigurationsComplete retrieves all the results into a single object
func (c GroupPolicyConfigurationClient) ListGroupPolicyConfigurationsComplete(ctx context.Context) (ListGroupPolicyConfigurationsCompleteResult, error) {
	return c.ListGroupPolicyConfigurationsCompleteMatchingPredicate(ctx, GroupPolicyConfigurationOperationPredicate{})
}

// ListGroupPolicyConfigurationsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c GroupPolicyConfigurationClient) ListGroupPolicyConfigurationsCompleteMatchingPredicate(ctx context.Context, predicate GroupPolicyConfigurationOperationPredicate) (result ListGroupPolicyConfigurationsCompleteResult, err error) {
	items := make([]beta.GroupPolicyConfiguration, 0)

	resp, err := c.ListGroupPolicyConfigurations(ctx)
	if err != nil {
		result.LatestHttpResponse = resp.HttpResponse
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

	result = ListGroupPolicyConfigurationsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
