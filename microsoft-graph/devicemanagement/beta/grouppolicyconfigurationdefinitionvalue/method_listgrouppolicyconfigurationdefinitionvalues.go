package grouppolicyconfigurationdefinitionvalue

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListGroupPolicyConfigurationDefinitionValuesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.GroupPolicyDefinitionValue
}

type ListGroupPolicyConfigurationDefinitionValuesCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.GroupPolicyDefinitionValue
}

type ListGroupPolicyConfigurationDefinitionValuesOperationOptions struct {
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

func DefaultListGroupPolicyConfigurationDefinitionValuesOperationOptions() ListGroupPolicyConfigurationDefinitionValuesOperationOptions {
	return ListGroupPolicyConfigurationDefinitionValuesOperationOptions{}
}

func (o ListGroupPolicyConfigurationDefinitionValuesOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListGroupPolicyConfigurationDefinitionValuesOperationOptions) ToOData() *odata.Query {
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

func (o ListGroupPolicyConfigurationDefinitionValuesOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListGroupPolicyConfigurationDefinitionValuesCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListGroupPolicyConfigurationDefinitionValuesCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListGroupPolicyConfigurationDefinitionValues - Get definitionValues from deviceManagement. The list of enabled or
// disabled group policy definition values for the configuration.
func (c GroupPolicyConfigurationDefinitionValueClient) ListGroupPolicyConfigurationDefinitionValues(ctx context.Context, id beta.DeviceManagementGroupPolicyConfigurationId, options ListGroupPolicyConfigurationDefinitionValuesOperationOptions) (result ListGroupPolicyConfigurationDefinitionValuesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListGroupPolicyConfigurationDefinitionValuesCustomPager{},
		Path:          fmt.Sprintf("%s/definitionValues", id.ID()),
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
		Values *[]beta.GroupPolicyDefinitionValue `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListGroupPolicyConfigurationDefinitionValuesComplete retrieves all the results into a single object
func (c GroupPolicyConfigurationDefinitionValueClient) ListGroupPolicyConfigurationDefinitionValuesComplete(ctx context.Context, id beta.DeviceManagementGroupPolicyConfigurationId, options ListGroupPolicyConfigurationDefinitionValuesOperationOptions) (ListGroupPolicyConfigurationDefinitionValuesCompleteResult, error) {
	return c.ListGroupPolicyConfigurationDefinitionValuesCompleteMatchingPredicate(ctx, id, options, GroupPolicyDefinitionValueOperationPredicate{})
}

// ListGroupPolicyConfigurationDefinitionValuesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c GroupPolicyConfigurationDefinitionValueClient) ListGroupPolicyConfigurationDefinitionValuesCompleteMatchingPredicate(ctx context.Context, id beta.DeviceManagementGroupPolicyConfigurationId, options ListGroupPolicyConfigurationDefinitionValuesOperationOptions, predicate GroupPolicyDefinitionValueOperationPredicate) (result ListGroupPolicyConfigurationDefinitionValuesCompleteResult, err error) {
	items := make([]beta.GroupPolicyDefinitionValue, 0)

	resp, err := c.ListGroupPolicyConfigurationDefinitionValues(ctx, id, options)
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

	result = ListGroupPolicyConfigurationDefinitionValuesCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
