package sitedocumentprocessingjob

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

type ListSiteDocumentProcessingJobsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.DocumentProcessingJob
}

type ListSiteDocumentProcessingJobsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.DocumentProcessingJob
}

type ListSiteDocumentProcessingJobsOperationOptions struct {
	Count   *bool
	Expand  *odata.Expand
	Filter  *string
	OrderBy *odata.OrderBy
	Search  *string
	Select  *[]string
	Skip    *int64
	Top     *int64
}

func DefaultListSiteDocumentProcessingJobsOperationOptions() ListSiteDocumentProcessingJobsOperationOptions {
	return ListSiteDocumentProcessingJobsOperationOptions{}
}

func (o ListSiteDocumentProcessingJobsOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListSiteDocumentProcessingJobsOperationOptions) ToOData() *odata.Query {
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

func (o ListSiteDocumentProcessingJobsOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListSiteDocumentProcessingJobsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListSiteDocumentProcessingJobsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListSiteDocumentProcessingJobs - Get documentProcessingJobs from groups. The document processing jobs running on this
// site.
func (c SiteDocumentProcessingJobClient) ListSiteDocumentProcessingJobs(ctx context.Context, id beta.GroupIdSiteId, options ListSiteDocumentProcessingJobsOperationOptions) (result ListSiteDocumentProcessingJobsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListSiteDocumentProcessingJobsCustomPager{},
		Path:          fmt.Sprintf("%s/documentProcessingJobs", id.ID()),
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
		Values *[]beta.DocumentProcessingJob `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListSiteDocumentProcessingJobsComplete retrieves all the results into a single object
func (c SiteDocumentProcessingJobClient) ListSiteDocumentProcessingJobsComplete(ctx context.Context, id beta.GroupIdSiteId, options ListSiteDocumentProcessingJobsOperationOptions) (ListSiteDocumentProcessingJobsCompleteResult, error) {
	return c.ListSiteDocumentProcessingJobsCompleteMatchingPredicate(ctx, id, options, DocumentProcessingJobOperationPredicate{})
}

// ListSiteDocumentProcessingJobsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c SiteDocumentProcessingJobClient) ListSiteDocumentProcessingJobsCompleteMatchingPredicate(ctx context.Context, id beta.GroupIdSiteId, options ListSiteDocumentProcessingJobsOperationOptions, predicate DocumentProcessingJobOperationPredicate) (result ListSiteDocumentProcessingJobsCompleteResult, err error) {
	items := make([]beta.DocumentProcessingJob, 0)

	resp, err := c.ListSiteDocumentProcessingJobs(ctx, id, options)
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

	result = ListSiteDocumentProcessingJobsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
