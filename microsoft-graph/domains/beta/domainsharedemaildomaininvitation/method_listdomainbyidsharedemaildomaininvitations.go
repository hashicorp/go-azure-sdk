package domainsharedemaildomaininvitation

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common/beta/models"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListDomainByIdSharedEmailDomainInvitationsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.SharedEmailDomainInvitationCollectionResponse
}

type ListDomainByIdSharedEmailDomainInvitationsCompleteResult struct {
	Items []models.SharedEmailDomainInvitationCollectionResponse
}

// ListDomainByIdSharedEmailDomainInvitations ...
func (c DomainSharedEmailDomainInvitationClient) ListDomainByIdSharedEmailDomainInvitations(ctx context.Context, id DomainId) (result ListDomainByIdSharedEmailDomainInvitationsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
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
		Values *[]models.SharedEmailDomainInvitationCollectionResponse `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListDomainByIdSharedEmailDomainInvitationsComplete retrieves all the results into a single object
func (c DomainSharedEmailDomainInvitationClient) ListDomainByIdSharedEmailDomainInvitationsComplete(ctx context.Context, id DomainId) (ListDomainByIdSharedEmailDomainInvitationsCompleteResult, error) {
	return c.ListDomainByIdSharedEmailDomainInvitationsCompleteMatchingPredicate(ctx, id, models.SharedEmailDomainInvitationCollectionResponseOperationPredicate{})
}

// ListDomainByIdSharedEmailDomainInvitationsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c DomainSharedEmailDomainInvitationClient) ListDomainByIdSharedEmailDomainInvitationsCompleteMatchingPredicate(ctx context.Context, id DomainId, predicate models.SharedEmailDomainInvitationCollectionResponseOperationPredicate) (result ListDomainByIdSharedEmailDomainInvitationsCompleteResult, err error) {
	items := make([]models.SharedEmailDomainInvitationCollectionResponse, 0)

	resp, err := c.ListDomainByIdSharedEmailDomainInvitations(ctx, id)
	if err != nil {
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

	result = ListDomainByIdSharedEmailDomainInvitationsCompleteResult{
		Items: items,
	}
	return
}
