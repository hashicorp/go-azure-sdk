package conditionalaccesauthenticationcontextclassreference

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

type ListConditionalAccesAuthenticationContextClassReferencesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]stable.AuthenticationContextClassReference
}

type ListConditionalAccesAuthenticationContextClassReferencesCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []stable.AuthenticationContextClassReference
}

type ListConditionalAccesAuthenticationContextClassReferencesCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListConditionalAccesAuthenticationContextClassReferencesCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListConditionalAccesAuthenticationContextClassReferences ...
func (c ConditionalAccesAuthenticationContextClassReferenceClient) ListConditionalAccesAuthenticationContextClassReferences(ctx context.Context) (result ListConditionalAccesAuthenticationContextClassReferencesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListConditionalAccesAuthenticationContextClassReferencesCustomPager{},
		Path:       "/identity/conditionalAccess/authenticationContextClassReferences",
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
		Values *[]stable.AuthenticationContextClassReference `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListConditionalAccesAuthenticationContextClassReferencesComplete retrieves all the results into a single object
func (c ConditionalAccesAuthenticationContextClassReferenceClient) ListConditionalAccesAuthenticationContextClassReferencesComplete(ctx context.Context) (ListConditionalAccesAuthenticationContextClassReferencesCompleteResult, error) {
	return c.ListConditionalAccesAuthenticationContextClassReferencesCompleteMatchingPredicate(ctx, AuthenticationContextClassReferenceOperationPredicate{})
}

// ListConditionalAccesAuthenticationContextClassReferencesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c ConditionalAccesAuthenticationContextClassReferenceClient) ListConditionalAccesAuthenticationContextClassReferencesCompleteMatchingPredicate(ctx context.Context, predicate AuthenticationContextClassReferenceOperationPredicate) (result ListConditionalAccesAuthenticationContextClassReferencesCompleteResult, err error) {
	items := make([]stable.AuthenticationContextClassReference, 0)

	resp, err := c.ListConditionalAccesAuthenticationContextClassReferences(ctx)
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

	result = ListConditionalAccesAuthenticationContextClassReferencesCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
