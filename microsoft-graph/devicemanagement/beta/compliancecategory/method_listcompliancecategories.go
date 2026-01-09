package compliancecategory

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

type ListComplianceCategoriesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.DeviceManagementConfigurationCategory
}

type ListComplianceCategoriesCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.DeviceManagementConfigurationCategory
}

type ListComplianceCategoriesOperationOptions struct {
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

func DefaultListComplianceCategoriesOperationOptions() ListComplianceCategoriesOperationOptions {
	return ListComplianceCategoriesOperationOptions{}
}

func (o ListComplianceCategoriesOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListComplianceCategoriesOperationOptions) ToOData() *odata.Query {
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

func (o ListComplianceCategoriesOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListComplianceCategoriesCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListComplianceCategoriesCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListComplianceCategories - Get complianceCategories from deviceManagement. List of all compliance categories
func (c ComplianceCategoryClient) ListComplianceCategories(ctx context.Context, options ListComplianceCategoriesOperationOptions) (result ListComplianceCategoriesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListComplianceCategoriesCustomPager{},
		Path:          "/deviceManagement/complianceCategories",
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
		Values *[]beta.DeviceManagementConfigurationCategory `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListComplianceCategoriesComplete retrieves all the results into a single object
func (c ComplianceCategoryClient) ListComplianceCategoriesComplete(ctx context.Context, options ListComplianceCategoriesOperationOptions) (ListComplianceCategoriesCompleteResult, error) {
	return c.ListComplianceCategoriesCompleteMatchingPredicate(ctx, options, DeviceManagementConfigurationCategoryOperationPredicate{})
}

// ListComplianceCategoriesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c ComplianceCategoryClient) ListComplianceCategoriesCompleteMatchingPredicate(ctx context.Context, options ListComplianceCategoriesOperationOptions, predicate DeviceManagementConfigurationCategoryOperationPredicate) (result ListComplianceCategoriesCompleteResult, err error) {
	items := make([]beta.DeviceManagementConfigurationCategory, 0)

	resp, err := c.ListComplianceCategories(ctx, options)
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

	result = ListComplianceCategoriesCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
