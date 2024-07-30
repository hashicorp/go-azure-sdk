package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CloudPCExportJobReportName string

const (
	CloudPCExportJobReportName_CloudPCUsageCategoryReports                 CloudPCExportJobReportName = "cloudPcUsageCategoryReports"
	CloudPCExportJobReportName_DailyAggregatedRemoteConnectionReports      CloudPCExportJobReportName = "dailyAggregatedRemoteConnectionReports"
	CloudPCExportJobReportName_FrontlineLicenseUsageRealTimeReport         CloudPCExportJobReportName = "frontlineLicenseUsageRealTimeReport"
	CloudPCExportJobReportName_FrontlineLicenseUsageReport                 CloudPCExportJobReportName = "frontlineLicenseUsageReport"
	CloudPCExportJobReportName_InaccessibleCloudPCReports                  CloudPCExportJobReportName = "inaccessibleCloudPcReports"
	CloudPCExportJobReportName_NoLicenseAvailableConnectivityFailureReport CloudPCExportJobReportName = "noLicenseAvailableConnectivityFailureReport"
	CloudPCExportJobReportName_RawRemoteConnectionReports                  CloudPCExportJobReportName = "rawRemoteConnectionReports"
	CloudPCExportJobReportName_RemoteConnectionHistoricalReports           CloudPCExportJobReportName = "remoteConnectionHistoricalReports"
	CloudPCExportJobReportName_RemoteConnectionQualityReports              CloudPCExportJobReportName = "remoteConnectionQualityReports"
	CloudPCExportJobReportName_SharedUseLicenseUsageRealTimeReport         CloudPCExportJobReportName = "sharedUseLicenseUsageRealTimeReport"
	CloudPCExportJobReportName_SharedUseLicenseUsageReport                 CloudPCExportJobReportName = "sharedUseLicenseUsageReport"
	CloudPCExportJobReportName_TotalAggregatedRemoteConnectionReports      CloudPCExportJobReportName = "totalAggregatedRemoteConnectionReports"
)

func PossibleValuesForCloudPCExportJobReportName() []string {
	return []string{
		string(CloudPCExportJobReportName_CloudPCUsageCategoryReports),
		string(CloudPCExportJobReportName_DailyAggregatedRemoteConnectionReports),
		string(CloudPCExportJobReportName_FrontlineLicenseUsageRealTimeReport),
		string(CloudPCExportJobReportName_FrontlineLicenseUsageReport),
		string(CloudPCExportJobReportName_InaccessibleCloudPCReports),
		string(CloudPCExportJobReportName_NoLicenseAvailableConnectivityFailureReport),
		string(CloudPCExportJobReportName_RawRemoteConnectionReports),
		string(CloudPCExportJobReportName_RemoteConnectionHistoricalReports),
		string(CloudPCExportJobReportName_RemoteConnectionQualityReports),
		string(CloudPCExportJobReportName_SharedUseLicenseUsageRealTimeReport),
		string(CloudPCExportJobReportName_SharedUseLicenseUsageReport),
		string(CloudPCExportJobReportName_TotalAggregatedRemoteConnectionReports),
	}
}

func (s *CloudPCExportJobReportName) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseCloudPCExportJobReportName(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseCloudPCExportJobReportName(input string) (*CloudPCExportJobReportName, error) {
	vals := map[string]CloudPCExportJobReportName{
		"cloudpcusagecategoryreports":                 CloudPCExportJobReportName_CloudPCUsageCategoryReports,
		"dailyaggregatedremoteconnectionreports":      CloudPCExportJobReportName_DailyAggregatedRemoteConnectionReports,
		"frontlinelicenseusagerealtimereport":         CloudPCExportJobReportName_FrontlineLicenseUsageRealTimeReport,
		"frontlinelicenseusagereport":                 CloudPCExportJobReportName_FrontlineLicenseUsageReport,
		"inaccessiblecloudpcreports":                  CloudPCExportJobReportName_InaccessibleCloudPCReports,
		"nolicenseavailableconnectivityfailurereport": CloudPCExportJobReportName_NoLicenseAvailableConnectivityFailureReport,
		"rawremoteconnectionreports":                  CloudPCExportJobReportName_RawRemoteConnectionReports,
		"remoteconnectionhistoricalreports":           CloudPCExportJobReportName_RemoteConnectionHistoricalReports,
		"remoteconnectionqualityreports":              CloudPCExportJobReportName_RemoteConnectionQualityReports,
		"shareduselicenseusagerealtimereport":         CloudPCExportJobReportName_SharedUseLicenseUsageRealTimeReport,
		"shareduselicenseusagereport":                 CloudPCExportJobReportName_SharedUseLicenseUsageReport,
		"totalaggregatedremoteconnectionreports":      CloudPCExportJobReportName_TotalAggregatedRemoteConnectionReports,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := CloudPCExportJobReportName(input)
	return &out, nil
}
