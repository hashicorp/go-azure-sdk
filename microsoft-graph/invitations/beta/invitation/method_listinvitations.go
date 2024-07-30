package invitation

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

type ListInvitationsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.Invitation
}

type ListInvitationsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.Invitation
}

type ListInvitationsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListInvitationsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListInvitations ...
func (c InvitationClient) ListInvitations(ctx context.Context) (result ListInvitationsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListInvitationsCustomPager{},
		Path:       "/invitations",
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
		Values *[]beta.Invitation `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListInvitationsComplete retrieves all the results into a single object
func (c InvitationClient) ListInvitationsComplete(ctx context.Context) (ListInvitationsCompleteResult, error) {
	return c.ListInvitationsCompleteMatchingPredicate(ctx, InvitationOperationPredicate{})
}

// ListInvitationsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c InvitationClient) ListInvitationsCompleteMatchingPredicate(ctx context.Context, predicate InvitationOperationPredicate) (result ListInvitationsCompleteResult, err error) {
	items := make([]beta.Invitation, 0)

	resp, err := c.ListInvitations(ctx)
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

	result = ListInvitationsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
