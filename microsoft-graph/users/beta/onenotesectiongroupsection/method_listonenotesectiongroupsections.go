package onenotesectiongroupsection

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

type ListOnenoteSectionGroupSectionsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.OnenoteSection
}

type ListOnenoteSectionGroupSectionsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.OnenoteSection
}

type ListOnenoteSectionGroupSectionsOperationOptions struct {
	Count    *bool
	Expand   *odata.Expand
	Filter   *string
	Metadata *odata.Metadata
	OrderBy  *odata.OrderBy
	Search   *string
	Select   *[]string
	Skip     *int64
	Top      *int64
}

func DefaultListOnenoteSectionGroupSectionsOperationOptions() ListOnenoteSectionGroupSectionsOperationOptions {
	return ListOnenoteSectionGroupSectionsOperationOptions{}
}

func (o ListOnenoteSectionGroupSectionsOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListOnenoteSectionGroupSectionsOperationOptions) ToOData() *odata.Query {
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

func (o ListOnenoteSectionGroupSectionsOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListOnenoteSectionGroupSectionsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListOnenoteSectionGroupSectionsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListOnenoteSectionGroupSections - Get sections from users. The sections in the section group. Read-only. Nullable.
func (c OnenoteSectionGroupSectionClient) ListOnenoteSectionGroupSections(ctx context.Context, id beta.UserIdOnenoteSectionGroupId, options ListOnenoteSectionGroupSectionsOperationOptions) (result ListOnenoteSectionGroupSectionsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListOnenoteSectionGroupSectionsCustomPager{},
		Path:          fmt.Sprintf("%s/sections", id.ID()),
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
		Values *[]beta.OnenoteSection `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListOnenoteSectionGroupSectionsComplete retrieves all the results into a single object
func (c OnenoteSectionGroupSectionClient) ListOnenoteSectionGroupSectionsComplete(ctx context.Context, id beta.UserIdOnenoteSectionGroupId, options ListOnenoteSectionGroupSectionsOperationOptions) (ListOnenoteSectionGroupSectionsCompleteResult, error) {
	return c.ListOnenoteSectionGroupSectionsCompleteMatchingPredicate(ctx, id, options, OnenoteSectionOperationPredicate{})
}

// ListOnenoteSectionGroupSectionsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c OnenoteSectionGroupSectionClient) ListOnenoteSectionGroupSectionsCompleteMatchingPredicate(ctx context.Context, id beta.UserIdOnenoteSectionGroupId, options ListOnenoteSectionGroupSectionsOperationOptions, predicate OnenoteSectionOperationPredicate) (result ListOnenoteSectionGroupSectionsCompleteResult, err error) {
	items := make([]beta.OnenoteSection, 0)

	resp, err := c.ListOnenoteSectionGroupSections(ctx, id, options)
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

	result = ListOnenoteSectionGroupSectionsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
