package profilenote

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

type ListProfileNotesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.PersonAnnotation
}

type ListProfileNotesCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.PersonAnnotation
}

type ListProfileNotesCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListProfileNotesCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListProfileNotes ...
func (c ProfileNoteClient) ListProfileNotes(ctx context.Context) (result ListProfileNotesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListProfileNotesCustomPager{},
		Path:       "/me/profile/notes",
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
		Values *[]beta.PersonAnnotation `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListProfileNotesComplete retrieves all the results into a single object
func (c ProfileNoteClient) ListProfileNotesComplete(ctx context.Context) (ListProfileNotesCompleteResult, error) {
	return c.ListProfileNotesCompleteMatchingPredicate(ctx, PersonAnnotationOperationPredicate{})
}

// ListProfileNotesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c ProfileNoteClient) ListProfileNotesCompleteMatchingPredicate(ctx context.Context, predicate PersonAnnotationOperationPredicate) (result ListProfileNotesCompleteResult, err error) {
	items := make([]beta.PersonAnnotation, 0)

	resp, err := c.ListProfileNotes(ctx)
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

	result = ListProfileNotesCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
