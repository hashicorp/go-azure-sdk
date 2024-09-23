package onenotenotebooksectiongroupsection

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/stable"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListOnenoteNotebookSectionGroupSectionsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]stable.OnenoteSection
}

type ListOnenoteNotebookSectionGroupSectionsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []stable.OnenoteSection
}

type ListOnenoteNotebookSectionGroupSectionsOperationOptions struct {
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

func DefaultListOnenoteNotebookSectionGroupSectionsOperationOptions() ListOnenoteNotebookSectionGroupSectionsOperationOptions {
	return ListOnenoteNotebookSectionGroupSectionsOperationOptions{}
}

func (o ListOnenoteNotebookSectionGroupSectionsOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListOnenoteNotebookSectionGroupSectionsOperationOptions) ToOData() *odata.Query {
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

func (o ListOnenoteNotebookSectionGroupSectionsOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListOnenoteNotebookSectionGroupSectionsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListOnenoteNotebookSectionGroupSectionsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListOnenoteNotebookSectionGroupSections - Get sections from me. The sections in the section group. Read-only.
// Nullable.
func (c OnenoteNotebookSectionGroupSectionClient) ListOnenoteNotebookSectionGroupSections(ctx context.Context, id stable.MeOnenoteNotebookIdSectionGroupId, options ListOnenoteNotebookSectionGroupSectionsOperationOptions) (result ListOnenoteNotebookSectionGroupSectionsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListOnenoteNotebookSectionGroupSectionsCustomPager{},
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

// ListOnenoteNotebookSectionGroupSectionsComplete retrieves all the results into a single object
func (c OnenoteNotebookSectionGroupSectionClient) ListOnenoteNotebookSectionGroupSectionsComplete(ctx context.Context, id stable.MeOnenoteNotebookIdSectionGroupId, options ListOnenoteNotebookSectionGroupSectionsOperationOptions) (ListOnenoteNotebookSectionGroupSectionsCompleteResult, error) {
	return c.ListOnenoteNotebookSectionGroupSectionsCompleteMatchingPredicate(ctx, id, options, OnenoteSectionOperationPredicate{})
}

// ListOnenoteNotebookSectionGroupSectionsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c OnenoteNotebookSectionGroupSectionClient) ListOnenoteNotebookSectionGroupSectionsCompleteMatchingPredicate(ctx context.Context, id stable.MeOnenoteNotebookIdSectionGroupId, options ListOnenoteNotebookSectionGroupSectionsOperationOptions, predicate OnenoteSectionOperationPredicate) (result ListOnenoteNotebookSectionGroupSectionsCompleteResult, err error) {
	items := make([]stable.OnenoteSection, 0)

	resp, err := c.ListOnenoteNotebookSectionGroupSections(ctx, id, options)
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

	result = ListOnenoteNotebookSectionGroupSectionsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
