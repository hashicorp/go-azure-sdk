package termsofuseagreementfileversion

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

type ListTermsOfUseAgreementFileVersionsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]stable.AgreementFileVersion
}

type ListTermsOfUseAgreementFileVersionsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []stable.AgreementFileVersion
}

type ListTermsOfUseAgreementFileVersionsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListTermsOfUseAgreementFileVersionsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListTermsOfUseAgreementFileVersions ...
func (c TermsOfUseAgreementFileVersionClient) ListTermsOfUseAgreementFileVersions(ctx context.Context, id IdentityGovernanceTermsOfUseAgreementIdFileId) (result ListTermsOfUseAgreementFileVersionsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListTermsOfUseAgreementFileVersionsCustomPager{},
		Path:       fmt.Sprintf("%s/versions", id.ID()),
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
		Values *[]stable.AgreementFileVersion `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListTermsOfUseAgreementFileVersionsComplete retrieves all the results into a single object
func (c TermsOfUseAgreementFileVersionClient) ListTermsOfUseAgreementFileVersionsComplete(ctx context.Context, id IdentityGovernanceTermsOfUseAgreementIdFileId) (ListTermsOfUseAgreementFileVersionsCompleteResult, error) {
	return c.ListTermsOfUseAgreementFileVersionsCompleteMatchingPredicate(ctx, id, AgreementFileVersionOperationPredicate{})
}

// ListTermsOfUseAgreementFileVersionsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c TermsOfUseAgreementFileVersionClient) ListTermsOfUseAgreementFileVersionsCompleteMatchingPredicate(ctx context.Context, id IdentityGovernanceTermsOfUseAgreementIdFileId, predicate AgreementFileVersionOperationPredicate) (result ListTermsOfUseAgreementFileVersionsCompleteResult, err error) {
	items := make([]stable.AgreementFileVersion, 0)

	resp, err := c.ListTermsOfUseAgreementFileVersions(ctx, id)
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

	result = ListTermsOfUseAgreementFileVersionsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
