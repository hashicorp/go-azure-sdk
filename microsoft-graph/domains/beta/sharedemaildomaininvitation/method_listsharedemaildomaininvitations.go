package sharedemaildomaininvitation

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

type ListSharedEmailDomainInvitationsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.SharedEmailDomainInvitation
}

type ListSharedEmailDomainInvitationsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.SharedEmailDomainInvitation
}

type ListSharedEmailDomainInvitationsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListSharedEmailDomainInvitationsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListSharedEmailDomainInvitations ...
func (c SharedEmailDomainInvitationClient) ListSharedEmailDomainInvitations(ctx context.Context, id DomainId) (result ListSharedEmailDomainInvitationsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListSharedEmailDomainInvitationsCustomPager{},
		Path:       fmt.Sprintf("%s/sharedEmailDomainInvitations", id.ID()),
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
		Values *[]beta.SharedEmailDomainInvitation `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListSharedEmailDomainInvitationsComplete retrieves all the results into a single object
func (c SharedEmailDomainInvitationClient) ListSharedEmailDomainInvitationsComplete(ctx context.Context, id DomainId) (ListSharedEmailDomainInvitationsCompleteResult, error) {
	return c.ListSharedEmailDomainInvitationsCompleteMatchingPredicate(ctx, id, SharedEmailDomainInvitationOperationPredicate{})
}

// ListSharedEmailDomainInvitationsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c SharedEmailDomainInvitationClient) ListSharedEmailDomainInvitationsCompleteMatchingPredicate(ctx context.Context, id DomainId, predicate SharedEmailDomainInvitationOperationPredicate) (result ListSharedEmailDomainInvitationsCompleteResult, err error) {
	items := make([]beta.SharedEmailDomainInvitation, 0)

	resp, err := c.ListSharedEmailDomainInvitations(ctx, id)
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

	result = ListSharedEmailDomainInvitationsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
