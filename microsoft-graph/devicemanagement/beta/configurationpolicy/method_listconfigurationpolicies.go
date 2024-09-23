package configurationpolicy

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

type ListConfigurationPoliciesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.DeviceManagementConfigurationPolicy
}

type ListConfigurationPoliciesCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.DeviceManagementConfigurationPolicy
}

type ListConfigurationPoliciesOperationOptions struct {
	Count     *bool
	Expand    *odata.Expand
	Filter    *string
	Metadata  *odata.Metadata
	OrderBy   *odata.OrderBy
	RetryFunc client.RequestRetryFunc
	Search    *string
	Select    *[]string
	Skip      *int64
	Top       *int64
}

func DefaultListConfigurationPoliciesOperationOptions() ListConfigurationPoliciesOperationOptions {
	return ListConfigurationPoliciesOperationOptions{}
}

func (o ListConfigurationPoliciesOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListConfigurationPoliciesOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Count != nil {
		out.Count = *o.Count
	}
	if o.Expand != nil {
		out.Expand = *o.Expand
	}
	if o.Filter != nil {
		out.Filter = *o.Filter
	}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	if o.OrderBy != nil {
		out.OrderBy = *o.OrderBy
	}
	if o.Search != nil {
		out.Search = *o.Search
	}
	if o.Select != nil {
		out.Select = *o.Select
	}
	if o.Skip != nil {
		out.Skip = int(*o.Skip)
	}
	if o.Top != nil {
		out.Top = int(*o.Top)
	}
	return &out
}

func (o ListConfigurationPoliciesOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListConfigurationPoliciesCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListConfigurationPoliciesCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListConfigurationPolicies - Get configurationPolicies from deviceManagement. List of all Configuration policies
func (c ConfigurationPolicyClient) ListConfigurationPolicies(ctx context.Context, options ListConfigurationPoliciesOperationOptions) (result ListConfigurationPoliciesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListConfigurationPoliciesCustomPager{},
		Path:          "/deviceManagement/configurationPolicies",
		RetryFunc:     options.RetryFunc,
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
		Values *[]beta.DeviceManagementConfigurationPolicy `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListConfigurationPoliciesComplete retrieves all the results into a single object
func (c ConfigurationPolicyClient) ListConfigurationPoliciesComplete(ctx context.Context, options ListConfigurationPoliciesOperationOptions) (ListConfigurationPoliciesCompleteResult, error) {
	return c.ListConfigurationPoliciesCompleteMatchingPredicate(ctx, options, DeviceManagementConfigurationPolicyOperationPredicate{})
}

// ListConfigurationPoliciesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c ConfigurationPolicyClient) ListConfigurationPoliciesCompleteMatchingPredicate(ctx context.Context, options ListConfigurationPoliciesOperationOptions, predicate DeviceManagementConfigurationPolicyOperationPredicate) (result ListConfigurationPoliciesCompleteResult, err error) {
	items := make([]beta.DeviceManagementConfigurationPolicy, 0)

	resp, err := c.ListConfigurationPolicies(ctx, options)
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

	result = ListConfigurationPoliciesCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
