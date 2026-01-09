package siteonenotenotebooksection

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/stable"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListSiteOnenoteNotebookSectionsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]stable.OnenoteSection
}

type ListSiteOnenoteNotebookSectionsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []stable.OnenoteSection
}

type ListSiteOnenoteNotebookSectionsOperationOptions struct {
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

func DefaultListSiteOnenoteNotebookSectionsOperationOptions() ListSiteOnenoteNotebookSectionsOperationOptions {
	return ListSiteOnenoteNotebookSectionsOperationOptions{}
}

func (o ListSiteOnenoteNotebookSectionsOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListSiteOnenoteNotebookSectionsOperationOptions) ToOData() *odata.Query {
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

func (o ListSiteOnenoteNotebookSectionsOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListSiteOnenoteNotebookSectionsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListSiteOnenoteNotebookSectionsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListSiteOnenoteNotebookSections - Get sections from groups. The sections in the notebook. Read-only. Nullable.
func (c SiteOnenoteNotebookSectionClient) ListSiteOnenoteNotebookSections(ctx context.Context, id stable.GroupIdSiteIdOnenoteNotebookId, options ListSiteOnenoteNotebookSectionsOperationOptions) (result ListSiteOnenoteNotebookSectionsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListSiteOnenoteNotebookSectionsCustomPager{},
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
		Values *[]stable.OnenoteSection `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListSiteOnenoteNotebookSectionsComplete retrieves all the results into a single object
func (c SiteOnenoteNotebookSectionClient) ListSiteOnenoteNotebookSectionsComplete(ctx context.Context, id stable.GroupIdSiteIdOnenoteNotebookId, options ListSiteOnenoteNotebookSectionsOperationOptions) (ListSiteOnenoteNotebookSectionsCompleteResult, error) {
	return c.ListSiteOnenoteNotebookSectionsCompleteMatchingPredicate(ctx, id, options, OnenoteSectionOperationPredicate{})
}

// ListSiteOnenoteNotebookSectionsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c SiteOnenoteNotebookSectionClient) ListSiteOnenoteNotebookSectionsCompleteMatchingPredicate(ctx context.Context, id stable.GroupIdSiteIdOnenoteNotebookId, options ListSiteOnenoteNotebookSectionsOperationOptions, predicate OnenoteSectionOperationPredicate) (result ListSiteOnenoteNotebookSectionsCompleteResult, err error) {
	items := make([]stable.OnenoteSection, 0)

	resp, err := c.ListSiteOnenoteNotebookSections(ctx, id, options)
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

	result = ListSiteOnenoteNotebookSectionsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
