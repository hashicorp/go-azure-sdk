package siteonenotenotebooksectiongroupsection

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

type ListSiteOnenoteNotebookSectionGroupSectionsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.OnenoteSection
}

type ListSiteOnenoteNotebookSectionGroupSectionsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.OnenoteSection
}

type ListSiteOnenoteNotebookSectionGroupSectionsOperationOptions struct {
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

func DefaultListSiteOnenoteNotebookSectionGroupSectionsOperationOptions() ListSiteOnenoteNotebookSectionGroupSectionsOperationOptions {
	return ListSiteOnenoteNotebookSectionGroupSectionsOperationOptions{}
}

func (o ListSiteOnenoteNotebookSectionGroupSectionsOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListSiteOnenoteNotebookSectionGroupSectionsOperationOptions) ToOData() *odata.Query {
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

func (o ListSiteOnenoteNotebookSectionGroupSectionsOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListSiteOnenoteNotebookSectionGroupSectionsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListSiteOnenoteNotebookSectionGroupSectionsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListSiteOnenoteNotebookSectionGroupSections - Get sections from groups. The sections in the section group. Read-only.
// Nullable.
func (c SiteOnenoteNotebookSectionGroupSectionClient) ListSiteOnenoteNotebookSectionGroupSections(ctx context.Context, id beta.GroupIdSiteIdOnenoteNotebookIdSectionGroupId, options ListSiteOnenoteNotebookSectionGroupSectionsOperationOptions) (result ListSiteOnenoteNotebookSectionGroupSectionsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListSiteOnenoteNotebookSectionGroupSectionsCustomPager{},
		Path:          fmt.Sprintf("%s/sections", id.ID()),
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
		Values *[]beta.OnenoteSection `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListSiteOnenoteNotebookSectionGroupSectionsComplete retrieves all the results into a single object
func (c SiteOnenoteNotebookSectionGroupSectionClient) ListSiteOnenoteNotebookSectionGroupSectionsComplete(ctx context.Context, id beta.GroupIdSiteIdOnenoteNotebookIdSectionGroupId, options ListSiteOnenoteNotebookSectionGroupSectionsOperationOptions) (ListSiteOnenoteNotebookSectionGroupSectionsCompleteResult, error) {
	return c.ListSiteOnenoteNotebookSectionGroupSectionsCompleteMatchingPredicate(ctx, id, options, OnenoteSectionOperationPredicate{})
}

// ListSiteOnenoteNotebookSectionGroupSectionsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c SiteOnenoteNotebookSectionGroupSectionClient) ListSiteOnenoteNotebookSectionGroupSectionsCompleteMatchingPredicate(ctx context.Context, id beta.GroupIdSiteIdOnenoteNotebookIdSectionGroupId, options ListSiteOnenoteNotebookSectionGroupSectionsOperationOptions, predicate OnenoteSectionOperationPredicate) (result ListSiteOnenoteNotebookSectionGroupSectionsCompleteResult, err error) {
	items := make([]beta.OnenoteSection, 0)

	resp, err := c.ListSiteOnenoteNotebookSectionGroupSections(ctx, id, options)
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

	result = ListSiteOnenoteNotebookSectionGroupSectionsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
