package compliancemanagementpartner

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

type ListComplianceManagementPartnersOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.ComplianceManagementPartner
}

type ListComplianceManagementPartnersCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.ComplianceManagementPartner
}

type ListComplianceManagementPartnersOperationOptions struct {
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

func DefaultListComplianceManagementPartnersOperationOptions() ListComplianceManagementPartnersOperationOptions {
	return ListComplianceManagementPartnersOperationOptions{}
}

func (o ListComplianceManagementPartnersOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListComplianceManagementPartnersOperationOptions) ToOData() *odata.Query {
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

func (o ListComplianceManagementPartnersOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListComplianceManagementPartnersCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListComplianceManagementPartnersCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListComplianceManagementPartners - Get complianceManagementPartners from deviceManagement. The list of Compliance
// Management Partners configured by the tenant.
func (c ComplianceManagementPartnerClient) ListComplianceManagementPartners(ctx context.Context, options ListComplianceManagementPartnersOperationOptions) (result ListComplianceManagementPartnersOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListComplianceManagementPartnersCustomPager{},
		Path:          "/deviceManagement/complianceManagementPartners",
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
		Values *[]beta.ComplianceManagementPartner `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListComplianceManagementPartnersComplete retrieves all the results into a single object
func (c ComplianceManagementPartnerClient) ListComplianceManagementPartnersComplete(ctx context.Context, options ListComplianceManagementPartnersOperationOptions) (ListComplianceManagementPartnersCompleteResult, error) {
	return c.ListComplianceManagementPartnersCompleteMatchingPredicate(ctx, options, ComplianceManagementPartnerOperationPredicate{})
}

// ListComplianceManagementPartnersCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c ComplianceManagementPartnerClient) ListComplianceManagementPartnersCompleteMatchingPredicate(ctx context.Context, options ListComplianceManagementPartnersOperationOptions, predicate ComplianceManagementPartnerOperationPredicate) (result ListComplianceManagementPartnersCompleteResult, err error) {
	items := make([]beta.ComplianceManagementPartner, 0)

	resp, err := c.ListComplianceManagementPartners(ctx, options)
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

	result = ListComplianceManagementPartnersCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
