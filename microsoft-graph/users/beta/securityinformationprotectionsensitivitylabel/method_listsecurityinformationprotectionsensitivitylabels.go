package securityinformationprotectionsensitivitylabel

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

type ListSecurityInformationProtectionSensitivityLabelsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.SecuritySensitivityLabel
}

type ListSecurityInformationProtectionSensitivityLabelsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.SecuritySensitivityLabel
}

type ListSecurityInformationProtectionSensitivityLabelsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListSecurityInformationProtectionSensitivityLabelsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListSecurityInformationProtectionSensitivityLabels ...
func (c SecurityInformationProtectionSensitivityLabelClient) ListSecurityInformationProtectionSensitivityLabels(ctx context.Context, id UserId) (result ListSecurityInformationProtectionSensitivityLabelsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListSecurityInformationProtectionSensitivityLabelsCustomPager{},
		Path:       fmt.Sprintf("%s/security/informationProtection/sensitivityLabels", id.ID()),
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
		Values *[]beta.SecuritySensitivityLabel `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListSecurityInformationProtectionSensitivityLabelsComplete retrieves all the results into a single object
func (c SecurityInformationProtectionSensitivityLabelClient) ListSecurityInformationProtectionSensitivityLabelsComplete(ctx context.Context, id UserId) (ListSecurityInformationProtectionSensitivityLabelsCompleteResult, error) {
	return c.ListSecurityInformationProtectionSensitivityLabelsCompleteMatchingPredicate(ctx, id, SecuritySensitivityLabelOperationPredicate{})
}

// ListSecurityInformationProtectionSensitivityLabelsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c SecurityInformationProtectionSensitivityLabelClient) ListSecurityInformationProtectionSensitivityLabelsCompleteMatchingPredicate(ctx context.Context, id UserId, predicate SecuritySensitivityLabelOperationPredicate) (result ListSecurityInformationProtectionSensitivityLabelsCompleteResult, err error) {
	items := make([]beta.SecuritySensitivityLabel, 0)

	resp, err := c.ListSecurityInformationProtectionSensitivityLabels(ctx, id)
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

	result = ListSecurityInformationProtectionSensitivityLabelsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
