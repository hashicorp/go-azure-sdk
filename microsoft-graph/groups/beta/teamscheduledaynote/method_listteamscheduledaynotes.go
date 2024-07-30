package teamscheduledaynote

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

type ListTeamScheduleDayNotesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.DayNote
}

type ListTeamScheduleDayNotesCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.DayNote
}

type ListTeamScheduleDayNotesCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListTeamScheduleDayNotesCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListTeamScheduleDayNotes ...
func (c TeamScheduleDayNoteClient) ListTeamScheduleDayNotes(ctx context.Context, id GroupId) (result ListTeamScheduleDayNotesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListTeamScheduleDayNotesCustomPager{},
		Path:       fmt.Sprintf("%s/team/schedule/dayNotes", id.ID()),
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
		Values *[]beta.DayNote `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListTeamScheduleDayNotesComplete retrieves all the results into a single object
func (c TeamScheduleDayNoteClient) ListTeamScheduleDayNotesComplete(ctx context.Context, id GroupId) (ListTeamScheduleDayNotesCompleteResult, error) {
	return c.ListTeamScheduleDayNotesCompleteMatchingPredicate(ctx, id, DayNoteOperationPredicate{})
}

// ListTeamScheduleDayNotesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c TeamScheduleDayNoteClient) ListTeamScheduleDayNotesCompleteMatchingPredicate(ctx context.Context, id GroupId, predicate DayNoteOperationPredicate) (result ListTeamScheduleDayNotesCompleteResult, err error) {
	items := make([]beta.DayNote, 0)

	resp, err := c.ListTeamScheduleDayNotes(ctx, id)
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

	result = ListTeamScheduleDayNotesCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
