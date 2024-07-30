package informationprotectionsensitivitylabelsublabel

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

type ListInformationProtectionSensitivityLabelSublabelsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.SensitivityLabel
}

type ListInformationProtectionSensitivityLabelSublabelsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.SensitivityLabel
}

type ListInformationProtectionSensitivityLabelSublabelsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListInformationProtectionSensitivityLabelSublabelsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListInformationProtectionSensitivityLabelSublabels ...
func (c InformationProtectionSensitivityLabelSublabelClient) ListInformationProtectionSensitivityLabelSublabels(ctx context.Context, id MeInformationProtectionSensitivityLabelId) (result ListInformationProtectionSensitivityLabelSublabelsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListInformationProtectionSensitivityLabelSublabelsCustomPager{},
		Path:       fmt.Sprintf("%s/sublabels", id.ID()),
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
		Values *[]beta.SensitivityLabel `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListInformationProtectionSensitivityLabelSublabelsComplete retrieves all the results into a single object
func (c InformationProtectionSensitivityLabelSublabelClient) ListInformationProtectionSensitivityLabelSublabelsComplete(ctx context.Context, id MeInformationProtectionSensitivityLabelId) (ListInformationProtectionSensitivityLabelSublabelsCompleteResult, error) {
	return c.ListInformationProtectionSensitivityLabelSublabelsCompleteMatchingPredicate(ctx, id, SensitivityLabelOperationPredicate{})
}

// ListInformationProtectionSensitivityLabelSublabelsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c InformationProtectionSensitivityLabelSublabelClient) ListInformationProtectionSensitivityLabelSublabelsCompleteMatchingPredicate(ctx context.Context, id MeInformationProtectionSensitivityLabelId, predicate SensitivityLabelOperationPredicate) (result ListInformationProtectionSensitivityLabelSublabelsCompleteResult, err error) {
	items := make([]beta.SensitivityLabel, 0)

	resp, err := c.ListInformationProtectionSensitivityLabelSublabels(ctx, id)
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

	result = ListInformationProtectionSensitivityLabelSublabelsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
