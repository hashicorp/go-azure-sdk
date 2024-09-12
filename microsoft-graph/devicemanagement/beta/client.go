package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/advancedthreatprotectiononboardingstatesummary"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/advancedthreatprotectiononboardingstatesummaryadvancedthreatprotectiononboardingdevicesettingstate"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/androiddeviceownerenrollmentprofile"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/androidforworkappconfigurationschema"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/androidforworkenrollmentprofile"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/androidforworksetting"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/androidmanagedstoreaccountenterprisesetting"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/androidmanagedstoreappconfigurationschema"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/applepushnotificationcertificate"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/appleuserinitiatedenrollmentprofile"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/appleuserinitiatedenrollmentprofileassignment"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/assignmentfilter"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/auditevent"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/autopilotevent"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/autopiloteventpolicystatusdetail"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/carttoclassassociation"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/category"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/categorysettingdefinition"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/certificateconnectordetail"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/chromeosonboardingsetting"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/cloudpcconnectivityissue"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/comanageddevice"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/comanageddeviceassignmentfilterevaluationstatusdetail"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/comanageddevicedetectedapp"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/comanageddevicedevicecategory"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/comanageddevicedevicecompliancepolicystate"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/comanageddevicedeviceconfigurationstate"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/comanageddevicedevicehealthscriptstate"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/comanageddevicedevicehealthscriptstateididpolicyidpolicyiddeviceiddeviceid"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/comanageddevicelogcollectionrequest"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/comanageddevicemanageddevicemobileappconfigurationstate"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/comanageddevicesecuritybaselinestate"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/comanageddevicesecuritybaselinestatesettingstate"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/comanageddeviceuser"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/comanageddevicewindowsprotectionstate"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/comanageddevicewindowsprotectionstatedetectedmalwarestate"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/comanagementeligibledevice"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/compliancecategory"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/compliancemanagementpartner"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/compliancepolicy"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/compliancepolicyassignment"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/compliancepolicyscheduledactionsforrule"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/compliancepolicyscheduledactionsforrulescheduledactionconfiguration"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/compliancepolicysetting"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/compliancepolicysettingsettingdefinition"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/compliancesetting"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/conditionalaccesssetting"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/configmanagercollection"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/configurationcategory"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/configurationpolicy"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/configurationpolicyassignment"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/configurationpolicysetting"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/configurationpolicysettingsettingdefinition"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/configurationpolicytemplate"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/configurationpolicytemplatesettingtemplate"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/configurationpolicytemplatesettingtemplatesettingdefinition"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/configurationsetting"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/datasharingconsent"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/deponboardingsetting"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/deponboardingsettingdefaultiosenrollmentprofile"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/deponboardingsettingdefaultmacosenrollmentprofile"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/deponboardingsettingenrollmentprofile"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/deponboardingsettingimportedappledeviceidentity"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/derivedcredential"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/detectedapp"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/detectedappmanageddevice"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/devicecategory"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/devicecompliancepolicy"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/devicecompliancepolicyassignment"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/devicecompliancepolicydevicesettingstatesummary"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/devicecompliancepolicydevicestatesummary"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/devicecompliancepolicydevicestatus"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/devicecompliancepolicydevicestatusoverview"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/devicecompliancepolicyscheduledactionsforrule"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/devicecompliancepolicyscheduledactionsforrulescheduledactionconfiguration"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/devicecompliancepolicysettingstatesummary"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/devicecompliancepolicysettingstatesummarydevicecompliancesettingstate"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/devicecompliancepolicyuserstatus"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/devicecompliancepolicyuserstatusoverview"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/devicecompliancescript"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/devicecompliancescriptassignment"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/devicecompliancescriptdevicerunstate"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/devicecompliancescriptdevicerunstatemanageddevice"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/devicecompliancescriptrunsummary"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/deviceconfiguration"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/deviceconfigurationassignment"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/deviceconfigurationconflictsummary"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/deviceconfigurationdevicesettingstatesummary"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/deviceconfigurationdevicestatesummary"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/deviceconfigurationdevicestatus"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/deviceconfigurationdevicestatusoverview"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/deviceconfigurationgroupassignment"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/deviceconfigurationgroupassignmentdeviceconfiguration"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/deviceconfigurationprofile"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/deviceconfigurationrestrictedappsviolation"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/deviceconfigurationsallmanageddevicecertificatestate"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/deviceconfigurationuserstatesummary"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/deviceconfigurationuserstatus"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/deviceconfigurationuserstatusoverview"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/devicecustomattributeshellscript"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/devicecustomattributeshellscriptassignment"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/devicecustomattributeshellscriptdevicerunstate"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/devicecustomattributeshellscriptdevicerunstatemanageddevice"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/devicecustomattributeshellscriptgroupassignment"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/devicecustomattributeshellscriptrunsummary"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/devicecustomattributeshellscriptuserrunstate"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/devicecustomattributeshellscriptuserrunstatedevicerunstate"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/devicecustomattributeshellscriptuserrunstatedevicerunstatemanageddevice"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/deviceenrollmentconfiguration"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/deviceenrollmentconfigurationassignment"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/devicehealthscript"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/devicehealthscriptassignment"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/devicehealthscriptdevicerunstate"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/devicehealthscriptdevicerunstatemanageddevice"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/devicehealthscriptrunsummary"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/devicemanagement"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/devicemanagementpartner"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/devicemanagementscript"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/devicemanagementscriptassignment"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/devicemanagementscriptdevicerunstate"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/devicemanagementscriptdevicerunstatemanageddevice"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/devicemanagementscriptgroupassignment"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/devicemanagementscriptrunsummary"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/devicemanagementscriptuserrunstate"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/devicemanagementscriptuserrunstatedevicerunstate"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/devicemanagementscriptuserrunstatedevicerunstatemanageddevice"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/deviceshellscript"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/deviceshellscriptassignment"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/deviceshellscriptdevicerunstate"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/deviceshellscriptdevicerunstatemanageddevice"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/deviceshellscriptgroupassignment"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/deviceshellscriptrunsummary"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/deviceshellscriptuserrunstate"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/deviceshellscriptuserrunstatedevicerunstate"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/deviceshellscriptuserrunstatedevicerunstatemanageddevice"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/domainjoinconnector"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/elevationrequest"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/embeddedsimactivationcodepool"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/embeddedsimactivationcodepoolassignment"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/embeddedsimactivationcodepooldevicestate"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/endpointprivilegemanagementprovisioningstatus"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/exchangeconnector"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/exchangeonpremisespolicy"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/exchangeonpremisespolicyconditionalaccesssetting"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/exportdevicemanagementreportsjobs"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/exportdevicemanagementvirtualendpointreportsjobs"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/grouppolicycategory"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/grouppolicycategorychild"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/grouppolicycategorydefinition"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/grouppolicycategorydefinitionfile"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/grouppolicycategoryparent"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/grouppolicyconfiguration"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/grouppolicyconfigurationassignment"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/grouppolicyconfigurationdefinitionvalue"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/grouppolicyconfigurationdefinitionvaluedefinition"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/grouppolicyconfigurationdefinitionvaluepresentationvalue"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/grouppolicyconfigurationdefinitionvaluepresentationvaluedefinitionvalue"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/grouppolicyconfigurationdefinitionvaluepresentationvaluepresentation"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/grouppolicydefinition"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/grouppolicydefinitioncategory"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/grouppolicydefinitiondefinitionfile"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/grouppolicydefinitionfile"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/grouppolicydefinitionfiledefinition"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/grouppolicydefinitionnextversiondefinition"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/grouppolicydefinitionnextversiondefinitioncategory"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/grouppolicydefinitionnextversiondefinitiondefinitionfile"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/grouppolicydefinitionnextversiondefinitionpresentation"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/grouppolicydefinitionnextversiondefinitionpresentationdefinition"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/grouppolicydefinitionnextversiondefinitionpreviousversiondefinition"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/grouppolicydefinitionnextversiondefinitionpreviousversiondefinitioncategory"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/grouppolicydefinitionnextversiondefinitionpreviousversiondefinitiondefinitionfile"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/grouppolicydefinitionnextversiondefinitionpreviousversiondefinitionpresentation"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/grouppolicydefinitionnextversiondefinitionpreviousversiondefinitionpresentationdefinition"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/grouppolicydefinitionpresentation"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/grouppolicydefinitionpresentationdefinition"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/grouppolicydefinitionpreviousversiondefinition"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/grouppolicydefinitionpreviousversiondefinitioncategory"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/grouppolicydefinitionpreviousversiondefinitiondefinitionfile"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/grouppolicydefinitionpreviousversiondefinitionnextversiondefinition"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/grouppolicydefinitionpreviousversiondefinitionnextversiondefinitioncategory"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/grouppolicydefinitionpreviousversiondefinitionnextversiondefinitiondefinitionfile"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/grouppolicydefinitionpreviousversiondefinitionnextversiondefinitionpresentation"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/grouppolicydefinitionpreviousversiondefinitionnextversiondefinitionpresentationdefinition"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/grouppolicydefinitionpreviousversiondefinitionpresentation"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/grouppolicydefinitionpreviousversiondefinitionpresentationdefinition"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/grouppolicymigrationreport"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/grouppolicymigrationreportgrouppolicysettingmapping"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/grouppolicymigrationreportunsupportedgrouppolicyextension"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/grouppolicyobjectfile"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/grouppolicyuploadeddefinitionfile"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/grouppolicyuploadeddefinitionfiledefinition"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/grouppolicyuploadeddefinitionfilegrouppolicyoperation"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/hardwareconfiguration"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/hardwareconfigurationassignment"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/hardwareconfigurationdevicerunstate"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/hardwareconfigurationrunsummary"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/hardwareconfigurationuserrunstate"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/hardwarepassworddetail"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/hardwarepasswordinfo"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/importeddeviceidentity"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/importedwindowsautopilotdeviceidentity"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/intent"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/intentassignment"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/intentcategory"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/intentcategorysetting"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/intentcategorysettingdefinition"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/intentdevicesettingstatesummary"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/intentdevicestate"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/intentdevicestatesummary"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/intentsetting"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/intentuserstate"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/intentuserstatesummary"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/intunebrandingprofile"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/intunebrandingprofileassignment"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/iosupdatestatus"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/macossoftwareupdateaccountsummary"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/macossoftwareupdateaccountsummarycategorysummary"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/manageddevice"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/manageddeviceassignmentfilterevaluationstatusdetail"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/manageddevicecleanuprule"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/manageddevicedetectedapp"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/manageddevicedevicecategory"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/manageddevicedevicecompliancepolicystate"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/manageddevicedeviceconfigurationstate"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/manageddevicedevicehealthscriptstate"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/manageddevicedevicehealthscriptstateididpolicyidpolicyiddeviceiddeviceid"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/manageddeviceencryptionstate"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/manageddevicelogcollectionrequest"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/manageddevicemanageddevicemobileappconfigurationstate"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/manageddeviceoverview"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/manageddevicesecuritybaselinestate"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/manageddevicesecuritybaselinestatesettingstate"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/manageddeviceuser"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/manageddevicewindowsosimage"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/manageddevicewindowsprotectionstate"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/manageddevicewindowsprotectionstatedetectedmalwarestate"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/microsofttunnelconfiguration"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/microsofttunnelhealththreshold"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/microsofttunnelserverlogcollectionresponse"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/microsofttunnelsite"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/microsofttunnelsitemicrosofttunnelconfiguration"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/microsofttunnelsitemicrosofttunnelserver"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/mobileapptroubleshootingevent"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/mobileapptroubleshootingeventapplogcollectionrequest"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/mobilethreatdefenseconnector"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/monitoring"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/monitoringalertrecord"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/monitoringalertrule"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/ndesconnector"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/notificationmessagetemplate"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/notificationmessagetemplatelocalizednotificationmessage"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/operationapprovalpolicy"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/operationapprovalrequest"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/privilegemanagementelevation"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/rebootdevicemanagementuserexperienceanalyticsbaselineanalyticsmetrics"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/remoteactionaudit"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/remoteassistancepartner"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/remoteassistancesetting"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/report"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/reportcachedreportconfiguration"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/resourceaccessprofile"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/resourceaccessprofileassignment"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/resourceoperation"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/reusablepolicysetting"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/reusablepolicysettingreferencingconfigurationpolicy"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/reusablepolicysettingreferencingconfigurationpolicyassignment"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/reusablepolicysettingreferencingconfigurationpolicysetting"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/reusablepolicysettingreferencingconfigurationpolicysettingsettingdefinition"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/reusablesetting"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/roleassignment"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/roleassignmentroledefinition"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/roleassignmentrolescopetag"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/roledefinition"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/roledefinitionroleassignment"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/roledefinitionroleassignmentroledefinition"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/rolescopetag"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/rolescopetagassignment"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/servicenowconnection"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/settingdefinition"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/softwareupdatestatussummary"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/telecomexpensemanagementpartner"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/template"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/templatecategory"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/templatecategoryrecommendedsetting"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/templatecategorysettingdefinition"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/templateinsight"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/templatemigratableto"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/templatemigratabletocategory"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/templatemigratabletocategoryrecommendedsetting"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/templatemigratabletocategorysettingdefinition"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/templatemigratabletosetting"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/templatesetting"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/templatesettingsettingdefinition"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/tenantattachrbac"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/termsandcondition"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/termsandconditionacceptancestatus"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/termsandconditionacceptancestatustermsandcondition"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/termsandconditionassignment"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/termsandconditiongroupassignment"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/termsandconditiongroupassignmenttermsandcondition"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/troubleshootingevent"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/updatedevicemanagementmacossoftwareupdateaccountsummarycategorysummarystatesummaries"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/userexperienceanalyticsanomaly"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/userexperienceanalyticsanomalycorrelationgroupoverview"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/userexperienceanalyticsanomalydevice"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/userexperienceanalyticsapphealthapplicationperformance"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/userexperienceanalyticsapphealthapplicationperformancebyappversion"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/userexperienceanalyticsapphealthapplicationperformancebyappversiondetail"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/userexperienceanalyticsapphealthapplicationperformancebyappversiondeviceid"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/userexperienceanalyticsapphealthapplicationperformancebyosversion"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/userexperienceanalyticsapphealthdevicemodelperformance"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/userexperienceanalyticsapphealthdeviceperformance"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/userexperienceanalyticsapphealthdeviceperformancedetail"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/userexperienceanalyticsapphealthosversionperformance"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/userexperienceanalyticsapphealthoverview"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/userexperienceanalyticsapphealthoverviewmetricvalue"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/userexperienceanalyticsbaseline"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/userexperienceanalyticsbaselineapphealthmetric"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/userexperienceanalyticsbaselinebatteryhealthmetric"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/userexperienceanalyticsbaselinebestpracticesmetric"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/userexperienceanalyticsbaselinedevicebootperformancemetric"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/userexperienceanalyticsbaselineresourceperformancemetric"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/userexperienceanalyticsbaselineworkfromanywheremetric"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/userexperienceanalyticsbatteryhealthappimpact"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/userexperienceanalyticsbatteryhealthcapacitydetail"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/userexperienceanalyticsbatteryhealthdeviceappimpact"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/userexperienceanalyticsbatteryhealthdeviceperformance"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/userexperienceanalyticsbatteryhealthdeviceruntimehistory"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/userexperienceanalyticsbatteryhealthmodelperformance"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/userexperienceanalyticsbatteryhealthosperformance"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/userexperienceanalyticsbatteryhealthruntimedetail"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/userexperienceanalyticscategory"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/userexperienceanalyticscategorymetricvalue"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/userexperienceanalyticsdevicemetrichistory"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/userexperienceanalyticsdeviceperformance"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/userexperienceanalyticsdevicescope"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/userexperienceanalyticsdevicescore"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/userexperienceanalyticsdevicestartuphistory"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/userexperienceanalyticsdevicestartupprocess"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/userexperienceanalyticsdevicestartupprocessperformance"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/userexperienceanalyticsdeviceswithoutcloudidentity"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/userexperienceanalyticsdevicetimelineevent"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/userexperienceanalyticsimpactingprocess"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/userexperienceanalyticsmetrichistory"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/userexperienceanalyticsmodelscore"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/userexperienceanalyticsnotautopilotreadydevice"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/userexperienceanalyticsoverview"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/userexperienceanalyticsremoteconnection"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/userexperienceanalyticsresourceperformance"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/userexperienceanalyticsscorehistory"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/userexperienceanalyticsworkfromanywherehardwarereadinessmetric"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/userexperienceanalyticsworkfromanywheremetric"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/userexperienceanalyticsworkfromanywheremetricmetricdevice"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/userexperienceanalyticsworkfromanywheremodelperformance"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/userpfxcertificate"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/virtualendpoint"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/virtualendpointauditevent"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/virtualendpointbulkaction"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/virtualendpointcloudpc"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/virtualendpointcrosscloudgovernmentorganizationmapping"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/virtualendpointdeviceimage"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/virtualendpointexternalpartnersetting"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/virtualendpointfrontlineserviceplan"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/virtualendpointgalleryimage"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/virtualendpointonpremisesconnection"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/virtualendpointorganizationsetting"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/virtualendpointprovisioningpolicy"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/virtualendpointprovisioningpolicyassignment"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/virtualendpointprovisioningpolicyassignmentassigneduser"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/virtualendpointprovisioningpolicyassignmentassignedusermailboxsetting"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/virtualendpointprovisioningpolicyassignmentassigneduserserviceprovisioningerror"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/virtualendpointreport"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/virtualendpointserviceplan"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/virtualendpointsnapshot"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/virtualendpointsupportedregion"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/virtualendpointusersetting"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/virtualendpointusersettingassignment"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/windowsautopilotdeploymentprofile"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/windowsautopilotdeploymentprofileassigneddevice"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/windowsautopilotdeploymentprofileassigneddevicedeploymentprofile"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/windowsautopilotdeploymentprofileassigneddeviceintendeddeploymentprofile"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/windowsautopilotdeploymentprofileassignment"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/windowsautopilotdeviceidentity"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/windowsautopilotdeviceidentitydeploymentprofile"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/windowsautopilotdeviceidentityintendeddeploymentprofile"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/windowsautopilotsetting"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/windowsdriverupdateprofile"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/windowsdriverupdateprofileassignment"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/windowsdriverupdateprofiledriverinventory"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/windowsfeatureupdateprofile"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/windowsfeatureupdateprofileassignment"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/windowsinformationprotectionapplearningsummary"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/windowsinformationprotectionnetworklearningsummary"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/windowsmalwareinformation"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/windowsmalwareinformationdevicemalwarestate"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/windowsqualityupdatepolicy"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/windowsqualityupdatepolicyassignment"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/windowsqualityupdateprofile"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/windowsqualityupdateprofileassignment"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/windowsupdatecatalogitem"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/zebrafotaartifact"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/zebrafotaconnector"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/zebrafotadeployment"
	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

type Client struct {
	AdvancedThreatProtectionOnboardingStateSummary                                                     *advancedthreatprotectiononboardingstatesummary.AdvancedThreatProtectionOnboardingStateSummaryClient
	AdvancedThreatProtectionOnboardingStateSummaryAdvancedThreatProtectionOnboardingDeviceSettingState *advancedthreatprotectiononboardingstatesummaryadvancedthreatprotectiononboardingdevicesettingstate.AdvancedThreatProtectionOnboardingStateSummaryAdvancedThreatProtectionOnboardingDeviceSettingStateClient
	AndroidDeviceOwnerEnrollmentProfile                                                                *androiddeviceownerenrollmentprofile.AndroidDeviceOwnerEnrollmentProfileClient
	AndroidForWorkAppConfigurationSchema                                                               *androidforworkappconfigurationschema.AndroidForWorkAppConfigurationSchemaClient
	AndroidForWorkEnrollmentProfile                                                                    *androidforworkenrollmentprofile.AndroidForWorkEnrollmentProfileClient
	AndroidForWorkSetting                                                                              *androidforworksetting.AndroidForWorkSettingClient
	AndroidManagedStoreAccountEnterpriseSetting                                                        *androidmanagedstoreaccountenterprisesetting.AndroidManagedStoreAccountEnterpriseSettingClient
	AndroidManagedStoreAppConfigurationSchema                                                          *androidmanagedstoreappconfigurationschema.AndroidManagedStoreAppConfigurationSchemaClient
	ApplePushNotificationCertificate                                                                   *applepushnotificationcertificate.ApplePushNotificationCertificateClient
	AppleUserInitiatedEnrollmentProfile                                                                *appleuserinitiatedenrollmentprofile.AppleUserInitiatedEnrollmentProfileClient
	AppleUserInitiatedEnrollmentProfileAssignment                                                      *appleuserinitiatedenrollmentprofileassignment.AppleUserInitiatedEnrollmentProfileAssignmentClient
	AssignmentFilter                                                                                   *assignmentfilter.AssignmentFilterClient
	AuditEvent                                                                                         *auditevent.AuditEventClient
	AutopilotEvent                                                                                     *autopilotevent.AutopilotEventClient
	AutopilotEventPolicyStatusDetail                                                                   *autopiloteventpolicystatusdetail.AutopilotEventPolicyStatusDetailClient
	CartToClassAssociation                                                                             *carttoclassassociation.CartToClassAssociationClient
	Category                                                                                           *category.CategoryClient
	CategorySettingDefinition                                                                          *categorysettingdefinition.CategorySettingDefinitionClient
	CertificateConnectorDetail                                                                         *certificateconnectordetail.CertificateConnectorDetailClient
	ChromeOSOnboardingSetting                                                                          *chromeosonboardingsetting.ChromeOSOnboardingSettingClient
	CloudPCConnectivityIssue                                                                           *cloudpcconnectivityissue.CloudPCConnectivityIssueClient
	ComanagedDevice                                                                                    *comanageddevice.ComanagedDeviceClient
	ComanagedDeviceAssignmentFilterEvaluationStatusDetail                                              *comanageddeviceassignmentfilterevaluationstatusdetail.ComanagedDeviceAssignmentFilterEvaluationStatusDetailClient
	ComanagedDeviceDetectedApp                                                                         *comanageddevicedetectedapp.ComanagedDeviceDetectedAppClient
	ComanagedDeviceDeviceCategory                                                                      *comanageddevicedevicecategory.ComanagedDeviceDeviceCategoryClient
	ComanagedDeviceDeviceCompliancePolicyState                                                         *comanageddevicedevicecompliancepolicystate.ComanagedDeviceDeviceCompliancePolicyStateClient
	ComanagedDeviceDeviceConfigurationState                                                            *comanageddevicedeviceconfigurationstate.ComanagedDeviceDeviceConfigurationStateClient
	ComanagedDeviceDeviceHealthScriptState                                                             *comanageddevicedevicehealthscriptstate.ComanagedDeviceDeviceHealthScriptStateClient
	ComanagedDeviceDeviceHealthScriptStateIdIdPolicyIdPolicyIdDeviceIdDeviceId                         *comanageddevicedevicehealthscriptstateididpolicyidpolicyiddeviceiddeviceid.ComanagedDeviceDeviceHealthScriptStateIdIdPolicyIdPolicyIdDeviceIdDeviceIdClient
	ComanagedDeviceLogCollectionRequest                                                                *comanageddevicelogcollectionrequest.ComanagedDeviceLogCollectionRequestClient
	ComanagedDeviceManagedDeviceMobileAppConfigurationState                                            *comanageddevicemanageddevicemobileappconfigurationstate.ComanagedDeviceManagedDeviceMobileAppConfigurationStateClient
	ComanagedDeviceSecurityBaselineState                                                               *comanageddevicesecuritybaselinestate.ComanagedDeviceSecurityBaselineStateClient
	ComanagedDeviceSecurityBaselineStateSettingState                                                   *comanageddevicesecuritybaselinestatesettingstate.ComanagedDeviceSecurityBaselineStateSettingStateClient
	ComanagedDeviceUser                                                                                *comanageddeviceuser.ComanagedDeviceUserClient
	ComanagedDeviceWindowsProtectionState                                                              *comanageddevicewindowsprotectionstate.ComanagedDeviceWindowsProtectionStateClient
	ComanagedDeviceWindowsProtectionStateDetectedMalwareState                                          *comanageddevicewindowsprotectionstatedetectedmalwarestate.ComanagedDeviceWindowsProtectionStateDetectedMalwareStateClient
	ComanagementEligibleDevice                                                                         *comanagementeligibledevice.ComanagementEligibleDeviceClient
	ComplianceCategory                                                                                 *compliancecategory.ComplianceCategoryClient
	ComplianceManagementPartner                                                                        *compliancemanagementpartner.ComplianceManagementPartnerClient
	CompliancePolicy                                                                                   *compliancepolicy.CompliancePolicyClient
	CompliancePolicyAssignment                                                                         *compliancepolicyassignment.CompliancePolicyAssignmentClient
	CompliancePolicyScheduledActionsForRule                                                            *compliancepolicyscheduledactionsforrule.CompliancePolicyScheduledActionsForRuleClient
	CompliancePolicyScheduledActionsForRuleScheduledActionConfiguration                                *compliancepolicyscheduledactionsforrulescheduledactionconfiguration.CompliancePolicyScheduledActionsForRuleScheduledActionConfigurationClient
	CompliancePolicySetting                                                                            *compliancepolicysetting.CompliancePolicySettingClient
	CompliancePolicySettingSettingDefinition                                                           *compliancepolicysettingsettingdefinition.CompliancePolicySettingSettingDefinitionClient
	ComplianceSetting                                                                                  *compliancesetting.ComplianceSettingClient
	ConditionalAccessSetting                                                                           *conditionalaccesssetting.ConditionalAccessSettingClient
	ConfigManagerCollection                                                                            *configmanagercollection.ConfigManagerCollectionClient
	ConfigurationCategory                                                                              *configurationcategory.ConfigurationCategoryClient
	ConfigurationPolicy                                                                                *configurationpolicy.ConfigurationPolicyClient
	ConfigurationPolicyAssignment                                                                      *configurationpolicyassignment.ConfigurationPolicyAssignmentClient
	ConfigurationPolicySetting                                                                         *configurationpolicysetting.ConfigurationPolicySettingClient
	ConfigurationPolicySettingSettingDefinition                                                        *configurationpolicysettingsettingdefinition.ConfigurationPolicySettingSettingDefinitionClient
	ConfigurationPolicyTemplate                                                                        *configurationpolicytemplate.ConfigurationPolicyTemplateClient
	ConfigurationPolicyTemplateSettingTemplate                                                         *configurationpolicytemplatesettingtemplate.ConfigurationPolicyTemplateSettingTemplateClient
	ConfigurationPolicyTemplateSettingTemplateSettingDefinition                                        *configurationpolicytemplatesettingtemplatesettingdefinition.ConfigurationPolicyTemplateSettingTemplateSettingDefinitionClient
	ConfigurationSetting                                                                               *configurationsetting.ConfigurationSettingClient
	DataSharingConsent                                                                                 *datasharingconsent.DataSharingConsentClient
	DepOnboardingSetting                                                                               *deponboardingsetting.DepOnboardingSettingClient
	DepOnboardingSettingDefaultIosEnrollmentProfile                                                    *deponboardingsettingdefaultiosenrollmentprofile.DepOnboardingSettingDefaultIosEnrollmentProfileClient
	DepOnboardingSettingDefaultMacOsEnrollmentProfile                                                  *deponboardingsettingdefaultmacosenrollmentprofile.DepOnboardingSettingDefaultMacOsEnrollmentProfileClient
	DepOnboardingSettingEnrollmentProfile                                                              *deponboardingsettingenrollmentprofile.DepOnboardingSettingEnrollmentProfileClient
	DepOnboardingSettingImportedAppleDeviceIdentity                                                    *deponboardingsettingimportedappledeviceidentity.DepOnboardingSettingImportedAppleDeviceIdentityClient
	DerivedCredential                                                                                  *derivedcredential.DerivedCredentialClient
	DetectedApp                                                                                        *detectedapp.DetectedAppClient
	DetectedAppManagedDevice                                                                           *detectedappmanageddevice.DetectedAppManagedDeviceClient
	DeviceCategory                                                                                     *devicecategory.DeviceCategoryClient
	DeviceCompliancePolicy                                                                             *devicecompliancepolicy.DeviceCompliancePolicyClient
	DeviceCompliancePolicyAssignment                                                                   *devicecompliancepolicyassignment.DeviceCompliancePolicyAssignmentClient
	DeviceCompliancePolicyDeviceSettingStateSummary                                                    *devicecompliancepolicydevicesettingstatesummary.DeviceCompliancePolicyDeviceSettingStateSummaryClient
	DeviceCompliancePolicyDeviceStateSummary                                                           *devicecompliancepolicydevicestatesummary.DeviceCompliancePolicyDeviceStateSummaryClient
	DeviceCompliancePolicyDeviceStatus                                                                 *devicecompliancepolicydevicestatus.DeviceCompliancePolicyDeviceStatusClient
	DeviceCompliancePolicyDeviceStatusOverview                                                         *devicecompliancepolicydevicestatusoverview.DeviceCompliancePolicyDeviceStatusOverviewClient
	DeviceCompliancePolicyScheduledActionsForRule                                                      *devicecompliancepolicyscheduledactionsforrule.DeviceCompliancePolicyScheduledActionsForRuleClient
	DeviceCompliancePolicyScheduledActionsForRuleScheduledActionConfiguration                          *devicecompliancepolicyscheduledactionsforrulescheduledactionconfiguration.DeviceCompliancePolicyScheduledActionsForRuleScheduledActionConfigurationClient
	DeviceCompliancePolicySettingStateSummary                                                          *devicecompliancepolicysettingstatesummary.DeviceCompliancePolicySettingStateSummaryClient
	DeviceCompliancePolicySettingStateSummaryDeviceComplianceSettingState                              *devicecompliancepolicysettingstatesummarydevicecompliancesettingstate.DeviceCompliancePolicySettingStateSummaryDeviceComplianceSettingStateClient
	DeviceCompliancePolicyUserStatus                                                                   *devicecompliancepolicyuserstatus.DeviceCompliancePolicyUserStatusClient
	DeviceCompliancePolicyUserStatusOverview                                                           *devicecompliancepolicyuserstatusoverview.DeviceCompliancePolicyUserStatusOverviewClient
	DeviceComplianceScript                                                                             *devicecompliancescript.DeviceComplianceScriptClient
	DeviceComplianceScriptAssignment                                                                   *devicecompliancescriptassignment.DeviceComplianceScriptAssignmentClient
	DeviceComplianceScriptDeviceRunState                                                               *devicecompliancescriptdevicerunstate.DeviceComplianceScriptDeviceRunStateClient
	DeviceComplianceScriptDeviceRunStateManagedDevice                                                  *devicecompliancescriptdevicerunstatemanageddevice.DeviceComplianceScriptDeviceRunStateManagedDeviceClient
	DeviceComplianceScriptRunSummary                                                                   *devicecompliancescriptrunsummary.DeviceComplianceScriptRunSummaryClient
	DeviceConfiguration                                                                                *deviceconfiguration.DeviceConfigurationClient
	DeviceConfigurationAssignment                                                                      *deviceconfigurationassignment.DeviceConfigurationAssignmentClient
	DeviceConfigurationConflictSummary                                                                 *deviceconfigurationconflictsummary.DeviceConfigurationConflictSummaryClient
	DeviceConfigurationDeviceSettingStateSummary                                                       *deviceconfigurationdevicesettingstatesummary.DeviceConfigurationDeviceSettingStateSummaryClient
	DeviceConfigurationDeviceStateSummary                                                              *deviceconfigurationdevicestatesummary.DeviceConfigurationDeviceStateSummaryClient
	DeviceConfigurationDeviceStatus                                                                    *deviceconfigurationdevicestatus.DeviceConfigurationDeviceStatusClient
	DeviceConfigurationDeviceStatusOverview                                                            *deviceconfigurationdevicestatusoverview.DeviceConfigurationDeviceStatusOverviewClient
	DeviceConfigurationGroupAssignment                                                                 *deviceconfigurationgroupassignment.DeviceConfigurationGroupAssignmentClient
	DeviceConfigurationGroupAssignmentDeviceConfiguration                                              *deviceconfigurationgroupassignmentdeviceconfiguration.DeviceConfigurationGroupAssignmentDeviceConfigurationClient
	DeviceConfigurationProfile                                                                         *deviceconfigurationprofile.DeviceConfigurationProfileClient
	DeviceConfigurationRestrictedAppsViolation                                                         *deviceconfigurationrestrictedappsviolation.DeviceConfigurationRestrictedAppsViolationClient
	DeviceConfigurationUserStateSummary                                                                *deviceconfigurationuserstatesummary.DeviceConfigurationUserStateSummaryClient
	DeviceConfigurationUserStatus                                                                      *deviceconfigurationuserstatus.DeviceConfigurationUserStatusClient
	DeviceConfigurationUserStatusOverview                                                              *deviceconfigurationuserstatusoverview.DeviceConfigurationUserStatusOverviewClient
	DeviceConfigurationsAllManagedDeviceCertificateState                                               *deviceconfigurationsallmanageddevicecertificatestate.DeviceConfigurationsAllManagedDeviceCertificateStateClient
	DeviceCustomAttributeShellScript                                                                   *devicecustomattributeshellscript.DeviceCustomAttributeShellScriptClient
	DeviceCustomAttributeShellScriptAssignment                                                         *devicecustomattributeshellscriptassignment.DeviceCustomAttributeShellScriptAssignmentClient
	DeviceCustomAttributeShellScriptDeviceRunState                                                     *devicecustomattributeshellscriptdevicerunstate.DeviceCustomAttributeShellScriptDeviceRunStateClient
	DeviceCustomAttributeShellScriptDeviceRunStateManagedDevice                                        *devicecustomattributeshellscriptdevicerunstatemanageddevice.DeviceCustomAttributeShellScriptDeviceRunStateManagedDeviceClient
	DeviceCustomAttributeShellScriptGroupAssignment                                                    *devicecustomattributeshellscriptgroupassignment.DeviceCustomAttributeShellScriptGroupAssignmentClient
	DeviceCustomAttributeShellScriptRunSummary                                                         *devicecustomattributeshellscriptrunsummary.DeviceCustomAttributeShellScriptRunSummaryClient
	DeviceCustomAttributeShellScriptUserRunState                                                       *devicecustomattributeshellscriptuserrunstate.DeviceCustomAttributeShellScriptUserRunStateClient
	DeviceCustomAttributeShellScriptUserRunStateDeviceRunState                                         *devicecustomattributeshellscriptuserrunstatedevicerunstate.DeviceCustomAttributeShellScriptUserRunStateDeviceRunStateClient
	DeviceCustomAttributeShellScriptUserRunStateDeviceRunStateManagedDevice                            *devicecustomattributeshellscriptuserrunstatedevicerunstatemanageddevice.DeviceCustomAttributeShellScriptUserRunStateDeviceRunStateManagedDeviceClient
	DeviceEnrollmentConfiguration                                                                      *deviceenrollmentconfiguration.DeviceEnrollmentConfigurationClient
	DeviceEnrollmentConfigurationAssignment                                                            *deviceenrollmentconfigurationassignment.DeviceEnrollmentConfigurationAssignmentClient
	DeviceHealthScript                                                                                 *devicehealthscript.DeviceHealthScriptClient
	DeviceHealthScriptAssignment                                                                       *devicehealthscriptassignment.DeviceHealthScriptAssignmentClient
	DeviceHealthScriptDeviceRunState                                                                   *devicehealthscriptdevicerunstate.DeviceHealthScriptDeviceRunStateClient
	DeviceHealthScriptDeviceRunStateManagedDevice                                                      *devicehealthscriptdevicerunstatemanageddevice.DeviceHealthScriptDeviceRunStateManagedDeviceClient
	DeviceHealthScriptRunSummary                                                                       *devicehealthscriptrunsummary.DeviceHealthScriptRunSummaryClient
	DeviceManagement                                                                                   *devicemanagement.DeviceManagementClient
	DeviceManagementPartner                                                                            *devicemanagementpartner.DeviceManagementPartnerClient
	DeviceManagementScript                                                                             *devicemanagementscript.DeviceManagementScriptClient
	DeviceManagementScriptAssignment                                                                   *devicemanagementscriptassignment.DeviceManagementScriptAssignmentClient
	DeviceManagementScriptDeviceRunState                                                               *devicemanagementscriptdevicerunstate.DeviceManagementScriptDeviceRunStateClient
	DeviceManagementScriptDeviceRunStateManagedDevice                                                  *devicemanagementscriptdevicerunstatemanageddevice.DeviceManagementScriptDeviceRunStateManagedDeviceClient
	DeviceManagementScriptGroupAssignment                                                              *devicemanagementscriptgroupassignment.DeviceManagementScriptGroupAssignmentClient
	DeviceManagementScriptRunSummary                                                                   *devicemanagementscriptrunsummary.DeviceManagementScriptRunSummaryClient
	DeviceManagementScriptUserRunState                                                                 *devicemanagementscriptuserrunstate.DeviceManagementScriptUserRunStateClient
	DeviceManagementScriptUserRunStateDeviceRunState                                                   *devicemanagementscriptuserrunstatedevicerunstate.DeviceManagementScriptUserRunStateDeviceRunStateClient
	DeviceManagementScriptUserRunStateDeviceRunStateManagedDevice                                      *devicemanagementscriptuserrunstatedevicerunstatemanageddevice.DeviceManagementScriptUserRunStateDeviceRunStateManagedDeviceClient
	DeviceShellScript                                                                                  *deviceshellscript.DeviceShellScriptClient
	DeviceShellScriptAssignment                                                                        *deviceshellscriptassignment.DeviceShellScriptAssignmentClient
	DeviceShellScriptDeviceRunState                                                                    *deviceshellscriptdevicerunstate.DeviceShellScriptDeviceRunStateClient
	DeviceShellScriptDeviceRunStateManagedDevice                                                       *deviceshellscriptdevicerunstatemanageddevice.DeviceShellScriptDeviceRunStateManagedDeviceClient
	DeviceShellScriptGroupAssignment                                                                   *deviceshellscriptgroupassignment.DeviceShellScriptGroupAssignmentClient
	DeviceShellScriptRunSummary                                                                        *deviceshellscriptrunsummary.DeviceShellScriptRunSummaryClient
	DeviceShellScriptUserRunState                                                                      *deviceshellscriptuserrunstate.DeviceShellScriptUserRunStateClient
	DeviceShellScriptUserRunStateDeviceRunState                                                        *deviceshellscriptuserrunstatedevicerunstate.DeviceShellScriptUserRunStateDeviceRunStateClient
	DeviceShellScriptUserRunStateDeviceRunStateManagedDevice                                           *deviceshellscriptuserrunstatedevicerunstatemanageddevice.DeviceShellScriptUserRunStateDeviceRunStateManagedDeviceClient
	DomainJoinConnector                                                                                *domainjoinconnector.DomainJoinConnectorClient
	ElevationRequest                                                                                   *elevationrequest.ElevationRequestClient
	EmbeddedSIMActivationCodePool                                                                      *embeddedsimactivationcodepool.EmbeddedSIMActivationCodePoolClient
	EmbeddedSIMActivationCodePoolAssignment                                                            *embeddedsimactivationcodepoolassignment.EmbeddedSIMActivationCodePoolAssignmentClient
	EmbeddedSIMActivationCodePoolDeviceState                                                           *embeddedsimactivationcodepooldevicestate.EmbeddedSIMActivationCodePoolDeviceStateClient
	EndpointPrivilegeManagementProvisioningStatus                                                      *endpointprivilegemanagementprovisioningstatus.EndpointPrivilegeManagementProvisioningStatusClient
	ExchangeConnector                                                                                  *exchangeconnector.ExchangeConnectorClient
	ExchangeOnPremisesPolicy                                                                           *exchangeonpremisespolicy.ExchangeOnPremisesPolicyClient
	ExchangeOnPremisesPolicyConditionalAccessSetting                                                   *exchangeonpremisespolicyconditionalaccesssetting.ExchangeOnPremisesPolicyConditionalAccessSettingClient
	ExportDeviceManagementReportsJobs                                                                  *exportdevicemanagementreportsjobs.ExportDeviceManagementReportsJobsClient
	ExportDeviceManagementVirtualEndpointReportsJobs                                                   *exportdevicemanagementvirtualendpointreportsjobs.ExportDeviceManagementVirtualEndpointReportsJobsClient
	GroupPolicyCategory                                                                                *grouppolicycategory.GroupPolicyCategoryClient
	GroupPolicyCategoryChild                                                                           *grouppolicycategorychild.GroupPolicyCategoryChildClient
	GroupPolicyCategoryDefinition                                                                      *grouppolicycategorydefinition.GroupPolicyCategoryDefinitionClient
	GroupPolicyCategoryDefinitionFile                                                                  *grouppolicycategorydefinitionfile.GroupPolicyCategoryDefinitionFileClient
	GroupPolicyCategoryParent                                                                          *grouppolicycategoryparent.GroupPolicyCategoryParentClient
	GroupPolicyConfiguration                                                                           *grouppolicyconfiguration.GroupPolicyConfigurationClient
	GroupPolicyConfigurationAssignment                                                                 *grouppolicyconfigurationassignment.GroupPolicyConfigurationAssignmentClient
	GroupPolicyConfigurationDefinitionValue                                                            *grouppolicyconfigurationdefinitionvalue.GroupPolicyConfigurationDefinitionValueClient
	GroupPolicyConfigurationDefinitionValueDefinition                                                  *grouppolicyconfigurationdefinitionvaluedefinition.GroupPolicyConfigurationDefinitionValueDefinitionClient
	GroupPolicyConfigurationDefinitionValuePresentationValue                                           *grouppolicyconfigurationdefinitionvaluepresentationvalue.GroupPolicyConfigurationDefinitionValuePresentationValueClient
	GroupPolicyConfigurationDefinitionValuePresentationValueDefinitionValue                            *grouppolicyconfigurationdefinitionvaluepresentationvaluedefinitionvalue.GroupPolicyConfigurationDefinitionValuePresentationValueDefinitionValueClient
	GroupPolicyConfigurationDefinitionValuePresentationValuePresentation                               *grouppolicyconfigurationdefinitionvaluepresentationvaluepresentation.GroupPolicyConfigurationDefinitionValuePresentationValuePresentationClient
	GroupPolicyDefinition                                                                              *grouppolicydefinition.GroupPolicyDefinitionClient
	GroupPolicyDefinitionCategory                                                                      *grouppolicydefinitioncategory.GroupPolicyDefinitionCategoryClient
	GroupPolicyDefinitionDefinitionFile                                                                *grouppolicydefinitiondefinitionfile.GroupPolicyDefinitionDefinitionFileClient
	GroupPolicyDefinitionFile                                                                          *grouppolicydefinitionfile.GroupPolicyDefinitionFileClient
	GroupPolicyDefinitionFileDefinition                                                                *grouppolicydefinitionfiledefinition.GroupPolicyDefinitionFileDefinitionClient
	GroupPolicyDefinitionNextVersionDefinition                                                         *grouppolicydefinitionnextversiondefinition.GroupPolicyDefinitionNextVersionDefinitionClient
	GroupPolicyDefinitionNextVersionDefinitionCategory                                                 *grouppolicydefinitionnextversiondefinitioncategory.GroupPolicyDefinitionNextVersionDefinitionCategoryClient
	GroupPolicyDefinitionNextVersionDefinitionDefinitionFile                                           *grouppolicydefinitionnextversiondefinitiondefinitionfile.GroupPolicyDefinitionNextVersionDefinitionDefinitionFileClient
	GroupPolicyDefinitionNextVersionDefinitionPresentation                                             *grouppolicydefinitionnextversiondefinitionpresentation.GroupPolicyDefinitionNextVersionDefinitionPresentationClient
	GroupPolicyDefinitionNextVersionDefinitionPresentationDefinition                                   *grouppolicydefinitionnextversiondefinitionpresentationdefinition.GroupPolicyDefinitionNextVersionDefinitionPresentationDefinitionClient
	GroupPolicyDefinitionNextVersionDefinitionPreviousVersionDefinition                                *grouppolicydefinitionnextversiondefinitionpreviousversiondefinition.GroupPolicyDefinitionNextVersionDefinitionPreviousVersionDefinitionClient
	GroupPolicyDefinitionNextVersionDefinitionPreviousVersionDefinitionCategory                        *grouppolicydefinitionnextversiondefinitionpreviousversiondefinitioncategory.GroupPolicyDefinitionNextVersionDefinitionPreviousVersionDefinitionCategoryClient
	GroupPolicyDefinitionNextVersionDefinitionPreviousVersionDefinitionDefinitionFile                  *grouppolicydefinitionnextversiondefinitionpreviousversiondefinitiondefinitionfile.GroupPolicyDefinitionNextVersionDefinitionPreviousVersionDefinitionDefinitionFileClient
	GroupPolicyDefinitionNextVersionDefinitionPreviousVersionDefinitionPresentation                    *grouppolicydefinitionnextversiondefinitionpreviousversiondefinitionpresentation.GroupPolicyDefinitionNextVersionDefinitionPreviousVersionDefinitionPresentationClient
	GroupPolicyDefinitionNextVersionDefinitionPreviousVersionDefinitionPresentationDefinition          *grouppolicydefinitionnextversiondefinitionpreviousversiondefinitionpresentationdefinition.GroupPolicyDefinitionNextVersionDefinitionPreviousVersionDefinitionPresentationDefinitionClient
	GroupPolicyDefinitionPresentation                                                                  *grouppolicydefinitionpresentation.GroupPolicyDefinitionPresentationClient
	GroupPolicyDefinitionPresentationDefinition                                                        *grouppolicydefinitionpresentationdefinition.GroupPolicyDefinitionPresentationDefinitionClient
	GroupPolicyDefinitionPreviousVersionDefinition                                                     *grouppolicydefinitionpreviousversiondefinition.GroupPolicyDefinitionPreviousVersionDefinitionClient
	GroupPolicyDefinitionPreviousVersionDefinitionCategory                                             *grouppolicydefinitionpreviousversiondefinitioncategory.GroupPolicyDefinitionPreviousVersionDefinitionCategoryClient
	GroupPolicyDefinitionPreviousVersionDefinitionDefinitionFile                                       *grouppolicydefinitionpreviousversiondefinitiondefinitionfile.GroupPolicyDefinitionPreviousVersionDefinitionDefinitionFileClient
	GroupPolicyDefinitionPreviousVersionDefinitionNextVersionDefinition                                *grouppolicydefinitionpreviousversiondefinitionnextversiondefinition.GroupPolicyDefinitionPreviousVersionDefinitionNextVersionDefinitionClient
	GroupPolicyDefinitionPreviousVersionDefinitionNextVersionDefinitionCategory                        *grouppolicydefinitionpreviousversiondefinitionnextversiondefinitioncategory.GroupPolicyDefinitionPreviousVersionDefinitionNextVersionDefinitionCategoryClient
	GroupPolicyDefinitionPreviousVersionDefinitionNextVersionDefinitionDefinitionFile                  *grouppolicydefinitionpreviousversiondefinitionnextversiondefinitiondefinitionfile.GroupPolicyDefinitionPreviousVersionDefinitionNextVersionDefinitionDefinitionFileClient
	GroupPolicyDefinitionPreviousVersionDefinitionNextVersionDefinitionPresentation                    *grouppolicydefinitionpreviousversiondefinitionnextversiondefinitionpresentation.GroupPolicyDefinitionPreviousVersionDefinitionNextVersionDefinitionPresentationClient
	GroupPolicyDefinitionPreviousVersionDefinitionNextVersionDefinitionPresentationDefinition          *grouppolicydefinitionpreviousversiondefinitionnextversiondefinitionpresentationdefinition.GroupPolicyDefinitionPreviousVersionDefinitionNextVersionDefinitionPresentationDefinitionClient
	GroupPolicyDefinitionPreviousVersionDefinitionPresentation                                         *grouppolicydefinitionpreviousversiondefinitionpresentation.GroupPolicyDefinitionPreviousVersionDefinitionPresentationClient
	GroupPolicyDefinitionPreviousVersionDefinitionPresentationDefinition                               *grouppolicydefinitionpreviousversiondefinitionpresentationdefinition.GroupPolicyDefinitionPreviousVersionDefinitionPresentationDefinitionClient
	GroupPolicyMigrationReport                                                                         *grouppolicymigrationreport.GroupPolicyMigrationReportClient
	GroupPolicyMigrationReportGroupPolicySettingMapping                                                *grouppolicymigrationreportgrouppolicysettingmapping.GroupPolicyMigrationReportGroupPolicySettingMappingClient
	GroupPolicyMigrationReportUnsupportedGroupPolicyExtension                                          *grouppolicymigrationreportunsupportedgrouppolicyextension.GroupPolicyMigrationReportUnsupportedGroupPolicyExtensionClient
	GroupPolicyObjectFile                                                                              *grouppolicyobjectfile.GroupPolicyObjectFileClient
	GroupPolicyUploadedDefinitionFile                                                                  *grouppolicyuploadeddefinitionfile.GroupPolicyUploadedDefinitionFileClient
	GroupPolicyUploadedDefinitionFileDefinition                                                        *grouppolicyuploadeddefinitionfiledefinition.GroupPolicyUploadedDefinitionFileDefinitionClient
	GroupPolicyUploadedDefinitionFileGroupPolicyOperation                                              *grouppolicyuploadeddefinitionfilegrouppolicyoperation.GroupPolicyUploadedDefinitionFileGroupPolicyOperationClient
	HardwareConfiguration                                                                              *hardwareconfiguration.HardwareConfigurationClient
	HardwareConfigurationAssignment                                                                    *hardwareconfigurationassignment.HardwareConfigurationAssignmentClient
	HardwareConfigurationDeviceRunState                                                                *hardwareconfigurationdevicerunstate.HardwareConfigurationDeviceRunStateClient
	HardwareConfigurationRunSummary                                                                    *hardwareconfigurationrunsummary.HardwareConfigurationRunSummaryClient
	HardwareConfigurationUserRunState                                                                  *hardwareconfigurationuserrunstate.HardwareConfigurationUserRunStateClient
	HardwarePasswordDetail                                                                             *hardwarepassworddetail.HardwarePasswordDetailClient
	HardwarePasswordInfo                                                                               *hardwarepasswordinfo.HardwarePasswordInfoClient
	ImportedDeviceIdentity                                                                             *importeddeviceidentity.ImportedDeviceIdentityClient
	ImportedWindowsAutopilotDeviceIdentity                                                             *importedwindowsautopilotdeviceidentity.ImportedWindowsAutopilotDeviceIdentityClient
	Intent                                                                                             *intent.IntentClient
	IntentAssignment                                                                                   *intentassignment.IntentAssignmentClient
	IntentCategory                                                                                     *intentcategory.IntentCategoryClient
	IntentCategorySetting                                                                              *intentcategorysetting.IntentCategorySettingClient
	IntentCategorySettingDefinition                                                                    *intentcategorysettingdefinition.IntentCategorySettingDefinitionClient
	IntentDeviceSettingStateSummary                                                                    *intentdevicesettingstatesummary.IntentDeviceSettingStateSummaryClient
	IntentDeviceState                                                                                  *intentdevicestate.IntentDeviceStateClient
	IntentDeviceStateSummary                                                                           *intentdevicestatesummary.IntentDeviceStateSummaryClient
	IntentSetting                                                                                      *intentsetting.IntentSettingClient
	IntentUserState                                                                                    *intentuserstate.IntentUserStateClient
	IntentUserStateSummary                                                                             *intentuserstatesummary.IntentUserStateSummaryClient
	IntuneBrandingProfile                                                                              *intunebrandingprofile.IntuneBrandingProfileClient
	IntuneBrandingProfileAssignment                                                                    *intunebrandingprofileassignment.IntuneBrandingProfileAssignmentClient
	IosUpdateStatus                                                                                    *iosupdatestatus.IosUpdateStatusClient
	MacOSSoftwareUpdateAccountSummary                                                                  *macossoftwareupdateaccountsummary.MacOSSoftwareUpdateAccountSummaryClient
	MacOSSoftwareUpdateAccountSummaryCategorySummary                                                   *macossoftwareupdateaccountsummarycategorysummary.MacOSSoftwareUpdateAccountSummaryCategorySummaryClient
	ManagedDevice                                                                                      *manageddevice.ManagedDeviceClient
	ManagedDeviceAssignmentFilterEvaluationStatusDetail                                                *manageddeviceassignmentfilterevaluationstatusdetail.ManagedDeviceAssignmentFilterEvaluationStatusDetailClient
	ManagedDeviceCleanupRule                                                                           *manageddevicecleanuprule.ManagedDeviceCleanupRuleClient
	ManagedDeviceDetectedApp                                                                           *manageddevicedetectedapp.ManagedDeviceDetectedAppClient
	ManagedDeviceDeviceCategory                                                                        *manageddevicedevicecategory.ManagedDeviceDeviceCategoryClient
	ManagedDeviceDeviceCompliancePolicyState                                                           *manageddevicedevicecompliancepolicystate.ManagedDeviceDeviceCompliancePolicyStateClient
	ManagedDeviceDeviceConfigurationState                                                              *manageddevicedeviceconfigurationstate.ManagedDeviceDeviceConfigurationStateClient
	ManagedDeviceDeviceHealthScriptState                                                               *manageddevicedevicehealthscriptstate.ManagedDeviceDeviceHealthScriptStateClient
	ManagedDeviceDeviceHealthScriptStateIdIdPolicyIdPolicyIdDeviceIdDeviceId                           *manageddevicedevicehealthscriptstateididpolicyidpolicyiddeviceiddeviceid.ManagedDeviceDeviceHealthScriptStateIdIdPolicyIdPolicyIdDeviceIdDeviceIdClient
	ManagedDeviceEncryptionState                                                                       *manageddeviceencryptionstate.ManagedDeviceEncryptionStateClient
	ManagedDeviceLogCollectionRequest                                                                  *manageddevicelogcollectionrequest.ManagedDeviceLogCollectionRequestClient
	ManagedDeviceManagedDeviceMobileAppConfigurationState                                              *manageddevicemanageddevicemobileappconfigurationstate.ManagedDeviceManagedDeviceMobileAppConfigurationStateClient
	ManagedDeviceOverview                                                                              *manageddeviceoverview.ManagedDeviceOverviewClient
	ManagedDeviceSecurityBaselineState                                                                 *manageddevicesecuritybaselinestate.ManagedDeviceSecurityBaselineStateClient
	ManagedDeviceSecurityBaselineStateSettingState                                                     *manageddevicesecuritybaselinestatesettingstate.ManagedDeviceSecurityBaselineStateSettingStateClient
	ManagedDeviceUser                                                                                  *manageddeviceuser.ManagedDeviceUserClient
	ManagedDeviceWindowsOSImage                                                                        *manageddevicewindowsosimage.ManagedDeviceWindowsOSImageClient
	ManagedDeviceWindowsProtectionState                                                                *manageddevicewindowsprotectionstate.ManagedDeviceWindowsProtectionStateClient
	ManagedDeviceWindowsProtectionStateDetectedMalwareState                                            *manageddevicewindowsprotectionstatedetectedmalwarestate.ManagedDeviceWindowsProtectionStateDetectedMalwareStateClient
	MicrosoftTunnelConfiguration                                                                       *microsofttunnelconfiguration.MicrosoftTunnelConfigurationClient
	MicrosoftTunnelHealthThreshold                                                                     *microsofttunnelhealththreshold.MicrosoftTunnelHealthThresholdClient
	MicrosoftTunnelServerLogCollectionResponse                                                         *microsofttunnelserverlogcollectionresponse.MicrosoftTunnelServerLogCollectionResponseClient
	MicrosoftTunnelSite                                                                                *microsofttunnelsite.MicrosoftTunnelSiteClient
	MicrosoftTunnelSiteMicrosoftTunnelConfiguration                                                    *microsofttunnelsitemicrosofttunnelconfiguration.MicrosoftTunnelSiteMicrosoftTunnelConfigurationClient
	MicrosoftTunnelSiteMicrosoftTunnelServer                                                           *microsofttunnelsitemicrosofttunnelserver.MicrosoftTunnelSiteMicrosoftTunnelServerClient
	MobileAppTroubleshootingEvent                                                                      *mobileapptroubleshootingevent.MobileAppTroubleshootingEventClient
	MobileAppTroubleshootingEventAppLogCollectionRequest                                               *mobileapptroubleshootingeventapplogcollectionrequest.MobileAppTroubleshootingEventAppLogCollectionRequestClient
	MobileThreatDefenseConnector                                                                       *mobilethreatdefenseconnector.MobileThreatDefenseConnectorClient
	Monitoring                                                                                         *monitoring.MonitoringClient
	MonitoringAlertRecord                                                                              *monitoringalertrecord.MonitoringAlertRecordClient
	MonitoringAlertRule                                                                                *monitoringalertrule.MonitoringAlertRuleClient
	NdesConnector                                                                                      *ndesconnector.NdesConnectorClient
	NotificationMessageTemplate                                                                        *notificationmessagetemplate.NotificationMessageTemplateClient
	NotificationMessageTemplateLocalizedNotificationMessage                                            *notificationmessagetemplatelocalizednotificationmessage.NotificationMessageTemplateLocalizedNotificationMessageClient
	OperationApprovalPolicy                                                                            *operationapprovalpolicy.OperationApprovalPolicyClient
	OperationApprovalRequest                                                                           *operationapprovalrequest.OperationApprovalRequestClient
	PrivilegeManagementElevation                                                                       *privilegemanagementelevation.PrivilegeManagementElevationClient
	RebootDeviceManagementUserExperienceAnalyticsBaselineAnalyticsMetrics                              *rebootdevicemanagementuserexperienceanalyticsbaselineanalyticsmetrics.RebootDeviceManagementUserExperienceAnalyticsBaselineAnalyticsMetricsClient
	RemoteActionAudit                                                                                  *remoteactionaudit.RemoteActionAuditClient
	RemoteAssistancePartner                                                                            *remoteassistancepartner.RemoteAssistancePartnerClient
	RemoteAssistanceSetting                                                                            *remoteassistancesetting.RemoteAssistanceSettingClient
	Report                                                                                             *report.ReportClient
	ReportCachedReportConfiguration                                                                    *reportcachedreportconfiguration.ReportCachedReportConfigurationClient
	ResourceAccessProfile                                                                              *resourceaccessprofile.ResourceAccessProfileClient
	ResourceAccessProfileAssignment                                                                    *resourceaccessprofileassignment.ResourceAccessProfileAssignmentClient
	ResourceOperation                                                                                  *resourceoperation.ResourceOperationClient
	ReusablePolicySetting                                                                              *reusablepolicysetting.ReusablePolicySettingClient
	ReusablePolicySettingReferencingConfigurationPolicy                                                *reusablepolicysettingreferencingconfigurationpolicy.ReusablePolicySettingReferencingConfigurationPolicyClient
	ReusablePolicySettingReferencingConfigurationPolicyAssignment                                      *reusablepolicysettingreferencingconfigurationpolicyassignment.ReusablePolicySettingReferencingConfigurationPolicyAssignmentClient
	ReusablePolicySettingReferencingConfigurationPolicySetting                                         *reusablepolicysettingreferencingconfigurationpolicysetting.ReusablePolicySettingReferencingConfigurationPolicySettingClient
	ReusablePolicySettingReferencingConfigurationPolicySettingSettingDefinition                        *reusablepolicysettingreferencingconfigurationpolicysettingsettingdefinition.ReusablePolicySettingReferencingConfigurationPolicySettingSettingDefinitionClient
	ReusableSetting                                                                                    *reusablesetting.ReusableSettingClient
	RoleAssignment                                                                                     *roleassignment.RoleAssignmentClient
	RoleAssignmentRoleDefinition                                                                       *roleassignmentroledefinition.RoleAssignmentRoleDefinitionClient
	RoleAssignmentRoleScopeTag                                                                         *roleassignmentrolescopetag.RoleAssignmentRoleScopeTagClient
	RoleDefinition                                                                                     *roledefinition.RoleDefinitionClient
	RoleDefinitionRoleAssignment                                                                       *roledefinitionroleassignment.RoleDefinitionRoleAssignmentClient
	RoleDefinitionRoleAssignmentRoleDefinition                                                         *roledefinitionroleassignmentroledefinition.RoleDefinitionRoleAssignmentRoleDefinitionClient
	RoleScopeTag                                                                                       *rolescopetag.RoleScopeTagClient
	RoleScopeTagAssignment                                                                             *rolescopetagassignment.RoleScopeTagAssignmentClient
	ServiceNowConnection                                                                               *servicenowconnection.ServiceNowConnectionClient
	SettingDefinition                                                                                  *settingdefinition.SettingDefinitionClient
	SoftwareUpdateStatusSummary                                                                        *softwareupdatestatussummary.SoftwareUpdateStatusSummaryClient
	TelecomExpenseManagementPartner                                                                    *telecomexpensemanagementpartner.TelecomExpenseManagementPartnerClient
	Template                                                                                           *template.TemplateClient
	TemplateCategory                                                                                   *templatecategory.TemplateCategoryClient
	TemplateCategoryRecommendedSetting                                                                 *templatecategoryrecommendedsetting.TemplateCategoryRecommendedSettingClient
	TemplateCategorySettingDefinition                                                                  *templatecategorysettingdefinition.TemplateCategorySettingDefinitionClient
	TemplateInsight                                                                                    *templateinsight.TemplateInsightClient
	TemplateMigratableTo                                                                               *templatemigratableto.TemplateMigratableToClient
	TemplateMigratableToCategory                                                                       *templatemigratabletocategory.TemplateMigratableToCategoryClient
	TemplateMigratableToCategoryRecommendedSetting                                                     *templatemigratabletocategoryrecommendedsetting.TemplateMigratableToCategoryRecommendedSettingClient
	TemplateMigratableToCategorySettingDefinition                                                      *templatemigratabletocategorysettingdefinition.TemplateMigratableToCategorySettingDefinitionClient
	TemplateMigratableToSetting                                                                        *templatemigratabletosetting.TemplateMigratableToSettingClient
	TemplateSetting                                                                                    *templatesetting.TemplateSettingClient
	TemplateSettingSettingDefinition                                                                   *templatesettingsettingdefinition.TemplateSettingSettingDefinitionClient
	TenantAttachRBAC                                                                                   *tenantattachrbac.TenantAttachRBACClient
	TermsAndCondition                                                                                  *termsandcondition.TermsAndConditionClient
	TermsAndConditionAcceptanceStatus                                                                  *termsandconditionacceptancestatus.TermsAndConditionAcceptanceStatusClient
	TermsAndConditionAcceptanceStatusTermsAndCondition                                                 *termsandconditionacceptancestatustermsandcondition.TermsAndConditionAcceptanceStatusTermsAndConditionClient
	TermsAndConditionAssignment                                                                        *termsandconditionassignment.TermsAndConditionAssignmentClient
	TermsAndConditionGroupAssignment                                                                   *termsandconditiongroupassignment.TermsAndConditionGroupAssignmentClient
	TermsAndConditionGroupAssignmentTermsAndCondition                                                  *termsandconditiongroupassignmenttermsandcondition.TermsAndConditionGroupAssignmentTermsAndConditionClient
	TroubleshootingEvent                                                                               *troubleshootingevent.TroubleshootingEventClient
	UpdateDeviceManagementMacOSSoftwareUpdateAccountSummaryCategorySummaryStateSummaries               *updatedevicemanagementmacossoftwareupdateaccountsummarycategorysummarystatesummaries.UpdateDeviceManagementMacOSSoftwareUpdateAccountSummaryCategorySummaryStateSummariesClient
	UserExperienceAnalyticsAnomaly                                                                     *userexperienceanalyticsanomaly.UserExperienceAnalyticsAnomalyClient
	UserExperienceAnalyticsAnomalyCorrelationGroupOverview                                             *userexperienceanalyticsanomalycorrelationgroupoverview.UserExperienceAnalyticsAnomalyCorrelationGroupOverviewClient
	UserExperienceAnalyticsAnomalyDevice                                                               *userexperienceanalyticsanomalydevice.UserExperienceAnalyticsAnomalyDeviceClient
	UserExperienceAnalyticsAppHealthApplicationPerformance                                             *userexperienceanalyticsapphealthapplicationperformance.UserExperienceAnalyticsAppHealthApplicationPerformanceClient
	UserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersion                                 *userexperienceanalyticsapphealthapplicationperformancebyappversion.UserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionClient
	UserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDetail                           *userexperienceanalyticsapphealthapplicationperformancebyappversiondetail.UserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDetailClient
	UserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDeviceId                         *userexperienceanalyticsapphealthapplicationperformancebyappversiondeviceid.UserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDeviceIdClient
	UserExperienceAnalyticsAppHealthApplicationPerformanceByOSVersion                                  *userexperienceanalyticsapphealthapplicationperformancebyosversion.UserExperienceAnalyticsAppHealthApplicationPerformanceByOSVersionClient
	UserExperienceAnalyticsAppHealthDeviceModelPerformance                                             *userexperienceanalyticsapphealthdevicemodelperformance.UserExperienceAnalyticsAppHealthDeviceModelPerformanceClient
	UserExperienceAnalyticsAppHealthDevicePerformance                                                  *userexperienceanalyticsapphealthdeviceperformance.UserExperienceAnalyticsAppHealthDevicePerformanceClient
	UserExperienceAnalyticsAppHealthDevicePerformanceDetail                                            *userexperienceanalyticsapphealthdeviceperformancedetail.UserExperienceAnalyticsAppHealthDevicePerformanceDetailClient
	UserExperienceAnalyticsAppHealthOSVersionPerformance                                               *userexperienceanalyticsapphealthosversionperformance.UserExperienceAnalyticsAppHealthOSVersionPerformanceClient
	UserExperienceAnalyticsAppHealthOverview                                                           *userexperienceanalyticsapphealthoverview.UserExperienceAnalyticsAppHealthOverviewClient
	UserExperienceAnalyticsAppHealthOverviewMetricValue                                                *userexperienceanalyticsapphealthoverviewmetricvalue.UserExperienceAnalyticsAppHealthOverviewMetricValueClient
	UserExperienceAnalyticsBaseline                                                                    *userexperienceanalyticsbaseline.UserExperienceAnalyticsBaselineClient
	UserExperienceAnalyticsBaselineAppHealthMetric                                                     *userexperienceanalyticsbaselineapphealthmetric.UserExperienceAnalyticsBaselineAppHealthMetricClient
	UserExperienceAnalyticsBaselineBatteryHealthMetric                                                 *userexperienceanalyticsbaselinebatteryhealthmetric.UserExperienceAnalyticsBaselineBatteryHealthMetricClient
	UserExperienceAnalyticsBaselineBestPracticesMetric                                                 *userexperienceanalyticsbaselinebestpracticesmetric.UserExperienceAnalyticsBaselineBestPracticesMetricClient
	UserExperienceAnalyticsBaselineDeviceBootPerformanceMetric                                         *userexperienceanalyticsbaselinedevicebootperformancemetric.UserExperienceAnalyticsBaselineDeviceBootPerformanceMetricClient
	UserExperienceAnalyticsBaselineResourcePerformanceMetric                                           *userexperienceanalyticsbaselineresourceperformancemetric.UserExperienceAnalyticsBaselineResourcePerformanceMetricClient
	UserExperienceAnalyticsBaselineWorkFromAnywhereMetric                                              *userexperienceanalyticsbaselineworkfromanywheremetric.UserExperienceAnalyticsBaselineWorkFromAnywhereMetricClient
	UserExperienceAnalyticsBatteryHealthAppImpact                                                      *userexperienceanalyticsbatteryhealthappimpact.UserExperienceAnalyticsBatteryHealthAppImpactClient
	UserExperienceAnalyticsBatteryHealthCapacityDetail                                                 *userexperienceanalyticsbatteryhealthcapacitydetail.UserExperienceAnalyticsBatteryHealthCapacityDetailClient
	UserExperienceAnalyticsBatteryHealthDeviceAppImpact                                                *userexperienceanalyticsbatteryhealthdeviceappimpact.UserExperienceAnalyticsBatteryHealthDeviceAppImpactClient
	UserExperienceAnalyticsBatteryHealthDevicePerformance                                              *userexperienceanalyticsbatteryhealthdeviceperformance.UserExperienceAnalyticsBatteryHealthDevicePerformanceClient
	UserExperienceAnalyticsBatteryHealthDeviceRuntimeHistory                                           *userexperienceanalyticsbatteryhealthdeviceruntimehistory.UserExperienceAnalyticsBatteryHealthDeviceRuntimeHistoryClient
	UserExperienceAnalyticsBatteryHealthModelPerformance                                               *userexperienceanalyticsbatteryhealthmodelperformance.UserExperienceAnalyticsBatteryHealthModelPerformanceClient
	UserExperienceAnalyticsBatteryHealthOsPerformance                                                  *userexperienceanalyticsbatteryhealthosperformance.UserExperienceAnalyticsBatteryHealthOsPerformanceClient
	UserExperienceAnalyticsBatteryHealthRuntimeDetail                                                  *userexperienceanalyticsbatteryhealthruntimedetail.UserExperienceAnalyticsBatteryHealthRuntimeDetailClient
	UserExperienceAnalyticsCategory                                                                    *userexperienceanalyticscategory.UserExperienceAnalyticsCategoryClient
	UserExperienceAnalyticsCategoryMetricValue                                                         *userexperienceanalyticscategorymetricvalue.UserExperienceAnalyticsCategoryMetricValueClient
	UserExperienceAnalyticsDeviceMetricHistory                                                         *userexperienceanalyticsdevicemetrichistory.UserExperienceAnalyticsDeviceMetricHistoryClient
	UserExperienceAnalyticsDevicePerformance                                                           *userexperienceanalyticsdeviceperformance.UserExperienceAnalyticsDevicePerformanceClient
	UserExperienceAnalyticsDeviceScope                                                                 *userexperienceanalyticsdevicescope.UserExperienceAnalyticsDeviceScopeClient
	UserExperienceAnalyticsDeviceScore                                                                 *userexperienceanalyticsdevicescore.UserExperienceAnalyticsDeviceScoreClient
	UserExperienceAnalyticsDeviceStartupHistory                                                        *userexperienceanalyticsdevicestartuphistory.UserExperienceAnalyticsDeviceStartupHistoryClient
	UserExperienceAnalyticsDeviceStartupProcess                                                        *userexperienceanalyticsdevicestartupprocess.UserExperienceAnalyticsDeviceStartupProcessClient
	UserExperienceAnalyticsDeviceStartupProcessPerformance                                             *userexperienceanalyticsdevicestartupprocessperformance.UserExperienceAnalyticsDeviceStartupProcessPerformanceClient
	UserExperienceAnalyticsDeviceTimelineEvent                                                         *userexperienceanalyticsdevicetimelineevent.UserExperienceAnalyticsDeviceTimelineEventClient
	UserExperienceAnalyticsDevicesWithoutCloudIdentity                                                 *userexperienceanalyticsdeviceswithoutcloudidentity.UserExperienceAnalyticsDevicesWithoutCloudIdentityClient
	UserExperienceAnalyticsImpactingProcess                                                            *userexperienceanalyticsimpactingprocess.UserExperienceAnalyticsImpactingProcessClient
	UserExperienceAnalyticsMetricHistory                                                               *userexperienceanalyticsmetrichistory.UserExperienceAnalyticsMetricHistoryClient
	UserExperienceAnalyticsModelScore                                                                  *userexperienceanalyticsmodelscore.UserExperienceAnalyticsModelScoreClient
	UserExperienceAnalyticsNotAutopilotReadyDevice                                                     *userexperienceanalyticsnotautopilotreadydevice.UserExperienceAnalyticsNotAutopilotReadyDeviceClient
	UserExperienceAnalyticsOverview                                                                    *userexperienceanalyticsoverview.UserExperienceAnalyticsOverviewClient
	UserExperienceAnalyticsRemoteConnection                                                            *userexperienceanalyticsremoteconnection.UserExperienceAnalyticsRemoteConnectionClient
	UserExperienceAnalyticsResourcePerformance                                                         *userexperienceanalyticsresourceperformance.UserExperienceAnalyticsResourcePerformanceClient
	UserExperienceAnalyticsScoreHistory                                                                *userexperienceanalyticsscorehistory.UserExperienceAnalyticsScoreHistoryClient
	UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric                                     *userexperienceanalyticsworkfromanywherehardwarereadinessmetric.UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetricClient
	UserExperienceAnalyticsWorkFromAnywhereMetric                                                      *userexperienceanalyticsworkfromanywheremetric.UserExperienceAnalyticsWorkFromAnywhereMetricClient
	UserExperienceAnalyticsWorkFromAnywhereMetricMetricDevice                                          *userexperienceanalyticsworkfromanywheremetricmetricdevice.UserExperienceAnalyticsWorkFromAnywhereMetricMetricDeviceClient
	UserExperienceAnalyticsWorkFromAnywhereModelPerformance                                            *userexperienceanalyticsworkfromanywheremodelperformance.UserExperienceAnalyticsWorkFromAnywhereModelPerformanceClient
	UserPfxCertificate                                                                                 *userpfxcertificate.UserPfxCertificateClient
	VirtualEndpoint                                                                                    *virtualendpoint.VirtualEndpointClient
	VirtualEndpointAuditEvent                                                                          *virtualendpointauditevent.VirtualEndpointAuditEventClient
	VirtualEndpointBulkAction                                                                          *virtualendpointbulkaction.VirtualEndpointBulkActionClient
	VirtualEndpointCloudPC                                                                             *virtualendpointcloudpc.VirtualEndpointCloudPCClient
	VirtualEndpointCrossCloudGovernmentOrganizationMapping                                             *virtualendpointcrosscloudgovernmentorganizationmapping.VirtualEndpointCrossCloudGovernmentOrganizationMappingClient
	VirtualEndpointDeviceImage                                                                         *virtualendpointdeviceimage.VirtualEndpointDeviceImageClient
	VirtualEndpointExternalPartnerSetting                                                              *virtualendpointexternalpartnersetting.VirtualEndpointExternalPartnerSettingClient
	VirtualEndpointFrontLineServicePlan                                                                *virtualendpointfrontlineserviceplan.VirtualEndpointFrontLineServicePlanClient
	VirtualEndpointGalleryImage                                                                        *virtualendpointgalleryimage.VirtualEndpointGalleryImageClient
	VirtualEndpointOnPremisesConnection                                                                *virtualendpointonpremisesconnection.VirtualEndpointOnPremisesConnectionClient
	VirtualEndpointOrganizationSetting                                                                 *virtualendpointorganizationsetting.VirtualEndpointOrganizationSettingClient
	VirtualEndpointProvisioningPolicy                                                                  *virtualendpointprovisioningpolicy.VirtualEndpointProvisioningPolicyClient
	VirtualEndpointProvisioningPolicyAssignment                                                        *virtualendpointprovisioningpolicyassignment.VirtualEndpointProvisioningPolicyAssignmentClient
	VirtualEndpointProvisioningPolicyAssignmentAssignedUser                                            *virtualendpointprovisioningpolicyassignmentassigneduser.VirtualEndpointProvisioningPolicyAssignmentAssignedUserClient
	VirtualEndpointProvisioningPolicyAssignmentAssignedUserMailboxSetting                              *virtualendpointprovisioningpolicyassignmentassignedusermailboxsetting.VirtualEndpointProvisioningPolicyAssignmentAssignedUserMailboxSettingClient
	VirtualEndpointProvisioningPolicyAssignmentAssignedUserServiceProvisioningError                    *virtualendpointprovisioningpolicyassignmentassigneduserserviceprovisioningerror.VirtualEndpointProvisioningPolicyAssignmentAssignedUserServiceProvisioningErrorClient
	VirtualEndpointReport                                                                              *virtualendpointreport.VirtualEndpointReportClient
	VirtualEndpointServicePlan                                                                         *virtualendpointserviceplan.VirtualEndpointServicePlanClient
	VirtualEndpointSnapshot                                                                            *virtualendpointsnapshot.VirtualEndpointSnapshotClient
	VirtualEndpointSupportedRegion                                                                     *virtualendpointsupportedregion.VirtualEndpointSupportedRegionClient
	VirtualEndpointUserSetting                                                                         *virtualendpointusersetting.VirtualEndpointUserSettingClient
	VirtualEndpointUserSettingAssignment                                                               *virtualendpointusersettingassignment.VirtualEndpointUserSettingAssignmentClient
	WindowsAutopilotDeploymentProfile                                                                  *windowsautopilotdeploymentprofile.WindowsAutopilotDeploymentProfileClient
	WindowsAutopilotDeploymentProfileAssignedDevice                                                    *windowsautopilotdeploymentprofileassigneddevice.WindowsAutopilotDeploymentProfileAssignedDeviceClient
	WindowsAutopilotDeploymentProfileAssignedDeviceDeploymentProfile                                   *windowsautopilotdeploymentprofileassigneddevicedeploymentprofile.WindowsAutopilotDeploymentProfileAssignedDeviceDeploymentProfileClient
	WindowsAutopilotDeploymentProfileAssignedDeviceIntendedDeploymentProfile                           *windowsautopilotdeploymentprofileassigneddeviceintendeddeploymentprofile.WindowsAutopilotDeploymentProfileAssignedDeviceIntendedDeploymentProfileClient
	WindowsAutopilotDeploymentProfileAssignment                                                        *windowsautopilotdeploymentprofileassignment.WindowsAutopilotDeploymentProfileAssignmentClient
	WindowsAutopilotDeviceIdentity                                                                     *windowsautopilotdeviceidentity.WindowsAutopilotDeviceIdentityClient
	WindowsAutopilotDeviceIdentityDeploymentProfile                                                    *windowsautopilotdeviceidentitydeploymentprofile.WindowsAutopilotDeviceIdentityDeploymentProfileClient
	WindowsAutopilotDeviceIdentityIntendedDeploymentProfile                                            *windowsautopilotdeviceidentityintendeddeploymentprofile.WindowsAutopilotDeviceIdentityIntendedDeploymentProfileClient
	WindowsAutopilotSetting                                                                            *windowsautopilotsetting.WindowsAutopilotSettingClient
	WindowsDriverUpdateProfile                                                                         *windowsdriverupdateprofile.WindowsDriverUpdateProfileClient
	WindowsDriverUpdateProfileAssignment                                                               *windowsdriverupdateprofileassignment.WindowsDriverUpdateProfileAssignmentClient
	WindowsDriverUpdateProfileDriverInventory                                                          *windowsdriverupdateprofiledriverinventory.WindowsDriverUpdateProfileDriverInventoryClient
	WindowsFeatureUpdateProfile                                                                        *windowsfeatureupdateprofile.WindowsFeatureUpdateProfileClient
	WindowsFeatureUpdateProfileAssignment                                                              *windowsfeatureupdateprofileassignment.WindowsFeatureUpdateProfileAssignmentClient
	WindowsInformationProtectionAppLearningSummary                                                     *windowsinformationprotectionapplearningsummary.WindowsInformationProtectionAppLearningSummaryClient
	WindowsInformationProtectionNetworkLearningSummary                                                 *windowsinformationprotectionnetworklearningsummary.WindowsInformationProtectionNetworkLearningSummaryClient
	WindowsMalwareInformation                                                                          *windowsmalwareinformation.WindowsMalwareInformationClient
	WindowsMalwareInformationDeviceMalwareState                                                        *windowsmalwareinformationdevicemalwarestate.WindowsMalwareInformationDeviceMalwareStateClient
	WindowsQualityUpdatePolicy                                                                         *windowsqualityupdatepolicy.WindowsQualityUpdatePolicyClient
	WindowsQualityUpdatePolicyAssignment                                                               *windowsqualityupdatepolicyassignment.WindowsQualityUpdatePolicyAssignmentClient
	WindowsQualityUpdateProfile                                                                        *windowsqualityupdateprofile.WindowsQualityUpdateProfileClient
	WindowsQualityUpdateProfileAssignment                                                              *windowsqualityupdateprofileassignment.WindowsQualityUpdateProfileAssignmentClient
	WindowsUpdateCatalogItem                                                                           *windowsupdatecatalogitem.WindowsUpdateCatalogItemClient
	ZebraFotaArtifact                                                                                  *zebrafotaartifact.ZebraFotaArtifactClient
	ZebraFotaConnector                                                                                 *zebrafotaconnector.ZebraFotaConnectorClient
	ZebraFotaDeployment                                                                                *zebrafotadeployment.ZebraFotaDeploymentClient
}

func NewClientWithBaseURI(sdkApi sdkEnv.Api, configureFunc func(c *msgraph.Client)) (*Client, error) {
	advancedThreatProtectionOnboardingStateSummaryAdvancedThreatProtectionOnboardingDeviceSettingStateClient, err := advancedthreatprotectiononboardingstatesummaryadvancedthreatprotectiononboardingdevicesettingstate.NewAdvancedThreatProtectionOnboardingStateSummaryAdvancedThreatProtectionOnboardingDeviceSettingStateClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building AdvancedThreatProtectionOnboardingStateSummaryAdvancedThreatProtectionOnboardingDeviceSettingState client: %+v", err)
	}
	configureFunc(advancedThreatProtectionOnboardingStateSummaryAdvancedThreatProtectionOnboardingDeviceSettingStateClient.Client)

	advancedThreatProtectionOnboardingStateSummaryClient, err := advancedthreatprotectiononboardingstatesummary.NewAdvancedThreatProtectionOnboardingStateSummaryClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building AdvancedThreatProtectionOnboardingStateSummary client: %+v", err)
	}
	configureFunc(advancedThreatProtectionOnboardingStateSummaryClient.Client)

	androidDeviceOwnerEnrollmentProfileClient, err := androiddeviceownerenrollmentprofile.NewAndroidDeviceOwnerEnrollmentProfileClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building AndroidDeviceOwnerEnrollmentProfile client: %+v", err)
	}
	configureFunc(androidDeviceOwnerEnrollmentProfileClient.Client)

	androidForWorkAppConfigurationSchemaClient, err := androidforworkappconfigurationschema.NewAndroidForWorkAppConfigurationSchemaClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building AndroidForWorkAppConfigurationSchema client: %+v", err)
	}
	configureFunc(androidForWorkAppConfigurationSchemaClient.Client)

	androidForWorkEnrollmentProfileClient, err := androidforworkenrollmentprofile.NewAndroidForWorkEnrollmentProfileClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building AndroidForWorkEnrollmentProfile client: %+v", err)
	}
	configureFunc(androidForWorkEnrollmentProfileClient.Client)

	androidForWorkSettingClient, err := androidforworksetting.NewAndroidForWorkSettingClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building AndroidForWorkSetting client: %+v", err)
	}
	configureFunc(androidForWorkSettingClient.Client)

	androidManagedStoreAccountEnterpriseSettingClient, err := androidmanagedstoreaccountenterprisesetting.NewAndroidManagedStoreAccountEnterpriseSettingClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building AndroidManagedStoreAccountEnterpriseSetting client: %+v", err)
	}
	configureFunc(androidManagedStoreAccountEnterpriseSettingClient.Client)

	androidManagedStoreAppConfigurationSchemaClient, err := androidmanagedstoreappconfigurationschema.NewAndroidManagedStoreAppConfigurationSchemaClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building AndroidManagedStoreAppConfigurationSchema client: %+v", err)
	}
	configureFunc(androidManagedStoreAppConfigurationSchemaClient.Client)

	applePushNotificationCertificateClient, err := applepushnotificationcertificate.NewApplePushNotificationCertificateClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ApplePushNotificationCertificate client: %+v", err)
	}
	configureFunc(applePushNotificationCertificateClient.Client)

	appleUserInitiatedEnrollmentProfileAssignmentClient, err := appleuserinitiatedenrollmentprofileassignment.NewAppleUserInitiatedEnrollmentProfileAssignmentClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building AppleUserInitiatedEnrollmentProfileAssignment client: %+v", err)
	}
	configureFunc(appleUserInitiatedEnrollmentProfileAssignmentClient.Client)

	appleUserInitiatedEnrollmentProfileClient, err := appleuserinitiatedenrollmentprofile.NewAppleUserInitiatedEnrollmentProfileClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building AppleUserInitiatedEnrollmentProfile client: %+v", err)
	}
	configureFunc(appleUserInitiatedEnrollmentProfileClient.Client)

	assignmentFilterClient, err := assignmentfilter.NewAssignmentFilterClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building AssignmentFilter client: %+v", err)
	}
	configureFunc(assignmentFilterClient.Client)

	auditEventClient, err := auditevent.NewAuditEventClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building AuditEvent client: %+v", err)
	}
	configureFunc(auditEventClient.Client)

	autopilotEventClient, err := autopilotevent.NewAutopilotEventClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building AutopilotEvent client: %+v", err)
	}
	configureFunc(autopilotEventClient.Client)

	autopilotEventPolicyStatusDetailClient, err := autopiloteventpolicystatusdetail.NewAutopilotEventPolicyStatusDetailClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building AutopilotEventPolicyStatusDetail client: %+v", err)
	}
	configureFunc(autopilotEventPolicyStatusDetailClient.Client)

	cartToClassAssociationClient, err := carttoclassassociation.NewCartToClassAssociationClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building CartToClassAssociation client: %+v", err)
	}
	configureFunc(cartToClassAssociationClient.Client)

	categoryClient, err := category.NewCategoryClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Category client: %+v", err)
	}
	configureFunc(categoryClient.Client)

	categorySettingDefinitionClient, err := categorysettingdefinition.NewCategorySettingDefinitionClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building CategorySettingDefinition client: %+v", err)
	}
	configureFunc(categorySettingDefinitionClient.Client)

	certificateConnectorDetailClient, err := certificateconnectordetail.NewCertificateConnectorDetailClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building CertificateConnectorDetail client: %+v", err)
	}
	configureFunc(certificateConnectorDetailClient.Client)

	chromeOSOnboardingSettingClient, err := chromeosonboardingsetting.NewChromeOSOnboardingSettingClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ChromeOSOnboardingSetting client: %+v", err)
	}
	configureFunc(chromeOSOnboardingSettingClient.Client)

	cloudPCConnectivityIssueClient, err := cloudpcconnectivityissue.NewCloudPCConnectivityIssueClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building CloudPCConnectivityIssue client: %+v", err)
	}
	configureFunc(cloudPCConnectivityIssueClient.Client)

	comanagedDeviceAssignmentFilterEvaluationStatusDetailClient, err := comanageddeviceassignmentfilterevaluationstatusdetail.NewComanagedDeviceAssignmentFilterEvaluationStatusDetailClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ComanagedDeviceAssignmentFilterEvaluationStatusDetail client: %+v", err)
	}
	configureFunc(comanagedDeviceAssignmentFilterEvaluationStatusDetailClient.Client)

	comanagedDeviceClient, err := comanageddevice.NewComanagedDeviceClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ComanagedDevice client: %+v", err)
	}
	configureFunc(comanagedDeviceClient.Client)

	comanagedDeviceDetectedAppClient, err := comanageddevicedetectedapp.NewComanagedDeviceDetectedAppClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ComanagedDeviceDetectedApp client: %+v", err)
	}
	configureFunc(comanagedDeviceDetectedAppClient.Client)

	comanagedDeviceDeviceCategoryClient, err := comanageddevicedevicecategory.NewComanagedDeviceDeviceCategoryClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ComanagedDeviceDeviceCategory client: %+v", err)
	}
	configureFunc(comanagedDeviceDeviceCategoryClient.Client)

	comanagedDeviceDeviceCompliancePolicyStateClient, err := comanageddevicedevicecompliancepolicystate.NewComanagedDeviceDeviceCompliancePolicyStateClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ComanagedDeviceDeviceCompliancePolicyState client: %+v", err)
	}
	configureFunc(comanagedDeviceDeviceCompliancePolicyStateClient.Client)

	comanagedDeviceDeviceConfigurationStateClient, err := comanageddevicedeviceconfigurationstate.NewComanagedDeviceDeviceConfigurationStateClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ComanagedDeviceDeviceConfigurationState client: %+v", err)
	}
	configureFunc(comanagedDeviceDeviceConfigurationStateClient.Client)

	comanagedDeviceDeviceHealthScriptStateClient, err := comanageddevicedevicehealthscriptstate.NewComanagedDeviceDeviceHealthScriptStateClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ComanagedDeviceDeviceHealthScriptState client: %+v", err)
	}
	configureFunc(comanagedDeviceDeviceHealthScriptStateClient.Client)

	comanagedDeviceDeviceHealthScriptStateIdIdPolicyIdPolicyIdDeviceIdDeviceIdClient, err := comanageddevicedevicehealthscriptstateididpolicyidpolicyiddeviceiddeviceid.NewComanagedDeviceDeviceHealthScriptStateIdIdPolicyIdPolicyIdDeviceIdDeviceIdClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ComanagedDeviceDeviceHealthScriptStateIdIdPolicyIdPolicyIdDeviceIdDeviceId client: %+v", err)
	}
	configureFunc(comanagedDeviceDeviceHealthScriptStateIdIdPolicyIdPolicyIdDeviceIdDeviceIdClient.Client)

	comanagedDeviceLogCollectionRequestClient, err := comanageddevicelogcollectionrequest.NewComanagedDeviceLogCollectionRequestClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ComanagedDeviceLogCollectionRequest client: %+v", err)
	}
	configureFunc(comanagedDeviceLogCollectionRequestClient.Client)

	comanagedDeviceManagedDeviceMobileAppConfigurationStateClient, err := comanageddevicemanageddevicemobileappconfigurationstate.NewComanagedDeviceManagedDeviceMobileAppConfigurationStateClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ComanagedDeviceManagedDeviceMobileAppConfigurationState client: %+v", err)
	}
	configureFunc(comanagedDeviceManagedDeviceMobileAppConfigurationStateClient.Client)

	comanagedDeviceSecurityBaselineStateClient, err := comanageddevicesecuritybaselinestate.NewComanagedDeviceSecurityBaselineStateClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ComanagedDeviceSecurityBaselineState client: %+v", err)
	}
	configureFunc(comanagedDeviceSecurityBaselineStateClient.Client)

	comanagedDeviceSecurityBaselineStateSettingStateClient, err := comanageddevicesecuritybaselinestatesettingstate.NewComanagedDeviceSecurityBaselineStateSettingStateClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ComanagedDeviceSecurityBaselineStateSettingState client: %+v", err)
	}
	configureFunc(comanagedDeviceSecurityBaselineStateSettingStateClient.Client)

	comanagedDeviceUserClient, err := comanageddeviceuser.NewComanagedDeviceUserClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ComanagedDeviceUser client: %+v", err)
	}
	configureFunc(comanagedDeviceUserClient.Client)

	comanagedDeviceWindowsProtectionStateClient, err := comanageddevicewindowsprotectionstate.NewComanagedDeviceWindowsProtectionStateClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ComanagedDeviceWindowsProtectionState client: %+v", err)
	}
	configureFunc(comanagedDeviceWindowsProtectionStateClient.Client)

	comanagedDeviceWindowsProtectionStateDetectedMalwareStateClient, err := comanageddevicewindowsprotectionstatedetectedmalwarestate.NewComanagedDeviceWindowsProtectionStateDetectedMalwareStateClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ComanagedDeviceWindowsProtectionStateDetectedMalwareState client: %+v", err)
	}
	configureFunc(comanagedDeviceWindowsProtectionStateDetectedMalwareStateClient.Client)

	comanagementEligibleDeviceClient, err := comanagementeligibledevice.NewComanagementEligibleDeviceClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ComanagementEligibleDevice client: %+v", err)
	}
	configureFunc(comanagementEligibleDeviceClient.Client)

	complianceCategoryClient, err := compliancecategory.NewComplianceCategoryClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ComplianceCategory client: %+v", err)
	}
	configureFunc(complianceCategoryClient.Client)

	complianceManagementPartnerClient, err := compliancemanagementpartner.NewComplianceManagementPartnerClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ComplianceManagementPartner client: %+v", err)
	}
	configureFunc(complianceManagementPartnerClient.Client)

	compliancePolicyAssignmentClient, err := compliancepolicyassignment.NewCompliancePolicyAssignmentClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building CompliancePolicyAssignment client: %+v", err)
	}
	configureFunc(compliancePolicyAssignmentClient.Client)

	compliancePolicyClient, err := compliancepolicy.NewCompliancePolicyClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building CompliancePolicy client: %+v", err)
	}
	configureFunc(compliancePolicyClient.Client)

	compliancePolicyScheduledActionsForRuleClient, err := compliancepolicyscheduledactionsforrule.NewCompliancePolicyScheduledActionsForRuleClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building CompliancePolicyScheduledActionsForRule client: %+v", err)
	}
	configureFunc(compliancePolicyScheduledActionsForRuleClient.Client)

	compliancePolicyScheduledActionsForRuleScheduledActionConfigurationClient, err := compliancepolicyscheduledactionsforrulescheduledactionconfiguration.NewCompliancePolicyScheduledActionsForRuleScheduledActionConfigurationClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building CompliancePolicyScheduledActionsForRuleScheduledActionConfiguration client: %+v", err)
	}
	configureFunc(compliancePolicyScheduledActionsForRuleScheduledActionConfigurationClient.Client)

	compliancePolicySettingClient, err := compliancepolicysetting.NewCompliancePolicySettingClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building CompliancePolicySetting client: %+v", err)
	}
	configureFunc(compliancePolicySettingClient.Client)

	compliancePolicySettingSettingDefinitionClient, err := compliancepolicysettingsettingdefinition.NewCompliancePolicySettingSettingDefinitionClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building CompliancePolicySettingSettingDefinition client: %+v", err)
	}
	configureFunc(compliancePolicySettingSettingDefinitionClient.Client)

	complianceSettingClient, err := compliancesetting.NewComplianceSettingClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ComplianceSetting client: %+v", err)
	}
	configureFunc(complianceSettingClient.Client)

	conditionalAccessSettingClient, err := conditionalaccesssetting.NewConditionalAccessSettingClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ConditionalAccessSetting client: %+v", err)
	}
	configureFunc(conditionalAccessSettingClient.Client)

	configManagerCollectionClient, err := configmanagercollection.NewConfigManagerCollectionClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ConfigManagerCollection client: %+v", err)
	}
	configureFunc(configManagerCollectionClient.Client)

	configurationCategoryClient, err := configurationcategory.NewConfigurationCategoryClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ConfigurationCategory client: %+v", err)
	}
	configureFunc(configurationCategoryClient.Client)

	configurationPolicyAssignmentClient, err := configurationpolicyassignment.NewConfigurationPolicyAssignmentClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ConfigurationPolicyAssignment client: %+v", err)
	}
	configureFunc(configurationPolicyAssignmentClient.Client)

	configurationPolicyClient, err := configurationpolicy.NewConfigurationPolicyClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ConfigurationPolicy client: %+v", err)
	}
	configureFunc(configurationPolicyClient.Client)

	configurationPolicySettingClient, err := configurationpolicysetting.NewConfigurationPolicySettingClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ConfigurationPolicySetting client: %+v", err)
	}
	configureFunc(configurationPolicySettingClient.Client)

	configurationPolicySettingSettingDefinitionClient, err := configurationpolicysettingsettingdefinition.NewConfigurationPolicySettingSettingDefinitionClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ConfigurationPolicySettingSettingDefinition client: %+v", err)
	}
	configureFunc(configurationPolicySettingSettingDefinitionClient.Client)

	configurationPolicyTemplateClient, err := configurationpolicytemplate.NewConfigurationPolicyTemplateClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ConfigurationPolicyTemplate client: %+v", err)
	}
	configureFunc(configurationPolicyTemplateClient.Client)

	configurationPolicyTemplateSettingTemplateClient, err := configurationpolicytemplatesettingtemplate.NewConfigurationPolicyTemplateSettingTemplateClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ConfigurationPolicyTemplateSettingTemplate client: %+v", err)
	}
	configureFunc(configurationPolicyTemplateSettingTemplateClient.Client)

	configurationPolicyTemplateSettingTemplateSettingDefinitionClient, err := configurationpolicytemplatesettingtemplatesettingdefinition.NewConfigurationPolicyTemplateSettingTemplateSettingDefinitionClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ConfigurationPolicyTemplateSettingTemplateSettingDefinition client: %+v", err)
	}
	configureFunc(configurationPolicyTemplateSettingTemplateSettingDefinitionClient.Client)

	configurationSettingClient, err := configurationsetting.NewConfigurationSettingClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ConfigurationSetting client: %+v", err)
	}
	configureFunc(configurationSettingClient.Client)

	dataSharingConsentClient, err := datasharingconsent.NewDataSharingConsentClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DataSharingConsent client: %+v", err)
	}
	configureFunc(dataSharingConsentClient.Client)

	depOnboardingSettingClient, err := deponboardingsetting.NewDepOnboardingSettingClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DepOnboardingSetting client: %+v", err)
	}
	configureFunc(depOnboardingSettingClient.Client)

	depOnboardingSettingDefaultIosEnrollmentProfileClient, err := deponboardingsettingdefaultiosenrollmentprofile.NewDepOnboardingSettingDefaultIosEnrollmentProfileClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DepOnboardingSettingDefaultIosEnrollmentProfile client: %+v", err)
	}
	configureFunc(depOnboardingSettingDefaultIosEnrollmentProfileClient.Client)

	depOnboardingSettingDefaultMacOsEnrollmentProfileClient, err := deponboardingsettingdefaultmacosenrollmentprofile.NewDepOnboardingSettingDefaultMacOsEnrollmentProfileClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DepOnboardingSettingDefaultMacOsEnrollmentProfile client: %+v", err)
	}
	configureFunc(depOnboardingSettingDefaultMacOsEnrollmentProfileClient.Client)

	depOnboardingSettingEnrollmentProfileClient, err := deponboardingsettingenrollmentprofile.NewDepOnboardingSettingEnrollmentProfileClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DepOnboardingSettingEnrollmentProfile client: %+v", err)
	}
	configureFunc(depOnboardingSettingEnrollmentProfileClient.Client)

	depOnboardingSettingImportedAppleDeviceIdentityClient, err := deponboardingsettingimportedappledeviceidentity.NewDepOnboardingSettingImportedAppleDeviceIdentityClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DepOnboardingSettingImportedAppleDeviceIdentity client: %+v", err)
	}
	configureFunc(depOnboardingSettingImportedAppleDeviceIdentityClient.Client)

	derivedCredentialClient, err := derivedcredential.NewDerivedCredentialClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DerivedCredential client: %+v", err)
	}
	configureFunc(derivedCredentialClient.Client)

	detectedAppClient, err := detectedapp.NewDetectedAppClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DetectedApp client: %+v", err)
	}
	configureFunc(detectedAppClient.Client)

	detectedAppManagedDeviceClient, err := detectedappmanageddevice.NewDetectedAppManagedDeviceClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DetectedAppManagedDevice client: %+v", err)
	}
	configureFunc(detectedAppManagedDeviceClient.Client)

	deviceCategoryClient, err := devicecategory.NewDeviceCategoryClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DeviceCategory client: %+v", err)
	}
	configureFunc(deviceCategoryClient.Client)

	deviceCompliancePolicyAssignmentClient, err := devicecompliancepolicyassignment.NewDeviceCompliancePolicyAssignmentClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DeviceCompliancePolicyAssignment client: %+v", err)
	}
	configureFunc(deviceCompliancePolicyAssignmentClient.Client)

	deviceCompliancePolicyClient, err := devicecompliancepolicy.NewDeviceCompliancePolicyClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DeviceCompliancePolicy client: %+v", err)
	}
	configureFunc(deviceCompliancePolicyClient.Client)

	deviceCompliancePolicyDeviceSettingStateSummaryClient, err := devicecompliancepolicydevicesettingstatesummary.NewDeviceCompliancePolicyDeviceSettingStateSummaryClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DeviceCompliancePolicyDeviceSettingStateSummary client: %+v", err)
	}
	configureFunc(deviceCompliancePolicyDeviceSettingStateSummaryClient.Client)

	deviceCompliancePolicyDeviceStateSummaryClient, err := devicecompliancepolicydevicestatesummary.NewDeviceCompliancePolicyDeviceStateSummaryClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DeviceCompliancePolicyDeviceStateSummary client: %+v", err)
	}
	configureFunc(deviceCompliancePolicyDeviceStateSummaryClient.Client)

	deviceCompliancePolicyDeviceStatusClient, err := devicecompliancepolicydevicestatus.NewDeviceCompliancePolicyDeviceStatusClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DeviceCompliancePolicyDeviceStatus client: %+v", err)
	}
	configureFunc(deviceCompliancePolicyDeviceStatusClient.Client)

	deviceCompliancePolicyDeviceStatusOverviewClient, err := devicecompliancepolicydevicestatusoverview.NewDeviceCompliancePolicyDeviceStatusOverviewClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DeviceCompliancePolicyDeviceStatusOverview client: %+v", err)
	}
	configureFunc(deviceCompliancePolicyDeviceStatusOverviewClient.Client)

	deviceCompliancePolicyScheduledActionsForRuleClient, err := devicecompliancepolicyscheduledactionsforrule.NewDeviceCompliancePolicyScheduledActionsForRuleClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DeviceCompliancePolicyScheduledActionsForRule client: %+v", err)
	}
	configureFunc(deviceCompliancePolicyScheduledActionsForRuleClient.Client)

	deviceCompliancePolicyScheduledActionsForRuleScheduledActionConfigurationClient, err := devicecompliancepolicyscheduledactionsforrulescheduledactionconfiguration.NewDeviceCompliancePolicyScheduledActionsForRuleScheduledActionConfigurationClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DeviceCompliancePolicyScheduledActionsForRuleScheduledActionConfiguration client: %+v", err)
	}
	configureFunc(deviceCompliancePolicyScheduledActionsForRuleScheduledActionConfigurationClient.Client)

	deviceCompliancePolicySettingStateSummaryClient, err := devicecompliancepolicysettingstatesummary.NewDeviceCompliancePolicySettingStateSummaryClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DeviceCompliancePolicySettingStateSummary client: %+v", err)
	}
	configureFunc(deviceCompliancePolicySettingStateSummaryClient.Client)

	deviceCompliancePolicySettingStateSummaryDeviceComplianceSettingStateClient, err := devicecompliancepolicysettingstatesummarydevicecompliancesettingstate.NewDeviceCompliancePolicySettingStateSummaryDeviceComplianceSettingStateClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DeviceCompliancePolicySettingStateSummaryDeviceComplianceSettingState client: %+v", err)
	}
	configureFunc(deviceCompliancePolicySettingStateSummaryDeviceComplianceSettingStateClient.Client)

	deviceCompliancePolicyUserStatusClient, err := devicecompliancepolicyuserstatus.NewDeviceCompliancePolicyUserStatusClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DeviceCompliancePolicyUserStatus client: %+v", err)
	}
	configureFunc(deviceCompliancePolicyUserStatusClient.Client)

	deviceCompliancePolicyUserStatusOverviewClient, err := devicecompliancepolicyuserstatusoverview.NewDeviceCompliancePolicyUserStatusOverviewClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DeviceCompliancePolicyUserStatusOverview client: %+v", err)
	}
	configureFunc(deviceCompliancePolicyUserStatusOverviewClient.Client)

	deviceComplianceScriptAssignmentClient, err := devicecompliancescriptassignment.NewDeviceComplianceScriptAssignmentClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DeviceComplianceScriptAssignment client: %+v", err)
	}
	configureFunc(deviceComplianceScriptAssignmentClient.Client)

	deviceComplianceScriptClient, err := devicecompliancescript.NewDeviceComplianceScriptClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DeviceComplianceScript client: %+v", err)
	}
	configureFunc(deviceComplianceScriptClient.Client)

	deviceComplianceScriptDeviceRunStateClient, err := devicecompliancescriptdevicerunstate.NewDeviceComplianceScriptDeviceRunStateClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DeviceComplianceScriptDeviceRunState client: %+v", err)
	}
	configureFunc(deviceComplianceScriptDeviceRunStateClient.Client)

	deviceComplianceScriptDeviceRunStateManagedDeviceClient, err := devicecompliancescriptdevicerunstatemanageddevice.NewDeviceComplianceScriptDeviceRunStateManagedDeviceClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DeviceComplianceScriptDeviceRunStateManagedDevice client: %+v", err)
	}
	configureFunc(deviceComplianceScriptDeviceRunStateManagedDeviceClient.Client)

	deviceComplianceScriptRunSummaryClient, err := devicecompliancescriptrunsummary.NewDeviceComplianceScriptRunSummaryClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DeviceComplianceScriptRunSummary client: %+v", err)
	}
	configureFunc(deviceComplianceScriptRunSummaryClient.Client)

	deviceConfigurationAssignmentClient, err := deviceconfigurationassignment.NewDeviceConfigurationAssignmentClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DeviceConfigurationAssignment client: %+v", err)
	}
	configureFunc(deviceConfigurationAssignmentClient.Client)

	deviceConfigurationClient, err := deviceconfiguration.NewDeviceConfigurationClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DeviceConfiguration client: %+v", err)
	}
	configureFunc(deviceConfigurationClient.Client)

	deviceConfigurationConflictSummaryClient, err := deviceconfigurationconflictsummary.NewDeviceConfigurationConflictSummaryClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DeviceConfigurationConflictSummary client: %+v", err)
	}
	configureFunc(deviceConfigurationConflictSummaryClient.Client)

	deviceConfigurationDeviceSettingStateSummaryClient, err := deviceconfigurationdevicesettingstatesummary.NewDeviceConfigurationDeviceSettingStateSummaryClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DeviceConfigurationDeviceSettingStateSummary client: %+v", err)
	}
	configureFunc(deviceConfigurationDeviceSettingStateSummaryClient.Client)

	deviceConfigurationDeviceStateSummaryClient, err := deviceconfigurationdevicestatesummary.NewDeviceConfigurationDeviceStateSummaryClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DeviceConfigurationDeviceStateSummary client: %+v", err)
	}
	configureFunc(deviceConfigurationDeviceStateSummaryClient.Client)

	deviceConfigurationDeviceStatusClient, err := deviceconfigurationdevicestatus.NewDeviceConfigurationDeviceStatusClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DeviceConfigurationDeviceStatus client: %+v", err)
	}
	configureFunc(deviceConfigurationDeviceStatusClient.Client)

	deviceConfigurationDeviceStatusOverviewClient, err := deviceconfigurationdevicestatusoverview.NewDeviceConfigurationDeviceStatusOverviewClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DeviceConfigurationDeviceStatusOverview client: %+v", err)
	}
	configureFunc(deviceConfigurationDeviceStatusOverviewClient.Client)

	deviceConfigurationGroupAssignmentClient, err := deviceconfigurationgroupassignment.NewDeviceConfigurationGroupAssignmentClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DeviceConfigurationGroupAssignment client: %+v", err)
	}
	configureFunc(deviceConfigurationGroupAssignmentClient.Client)

	deviceConfigurationGroupAssignmentDeviceConfigurationClient, err := deviceconfigurationgroupassignmentdeviceconfiguration.NewDeviceConfigurationGroupAssignmentDeviceConfigurationClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DeviceConfigurationGroupAssignmentDeviceConfiguration client: %+v", err)
	}
	configureFunc(deviceConfigurationGroupAssignmentDeviceConfigurationClient.Client)

	deviceConfigurationProfileClient, err := deviceconfigurationprofile.NewDeviceConfigurationProfileClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DeviceConfigurationProfile client: %+v", err)
	}
	configureFunc(deviceConfigurationProfileClient.Client)

	deviceConfigurationRestrictedAppsViolationClient, err := deviceconfigurationrestrictedappsviolation.NewDeviceConfigurationRestrictedAppsViolationClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DeviceConfigurationRestrictedAppsViolation client: %+v", err)
	}
	configureFunc(deviceConfigurationRestrictedAppsViolationClient.Client)

	deviceConfigurationUserStateSummaryClient, err := deviceconfigurationuserstatesummary.NewDeviceConfigurationUserStateSummaryClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DeviceConfigurationUserStateSummary client: %+v", err)
	}
	configureFunc(deviceConfigurationUserStateSummaryClient.Client)

	deviceConfigurationUserStatusClient, err := deviceconfigurationuserstatus.NewDeviceConfigurationUserStatusClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DeviceConfigurationUserStatus client: %+v", err)
	}
	configureFunc(deviceConfigurationUserStatusClient.Client)

	deviceConfigurationUserStatusOverviewClient, err := deviceconfigurationuserstatusoverview.NewDeviceConfigurationUserStatusOverviewClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DeviceConfigurationUserStatusOverview client: %+v", err)
	}
	configureFunc(deviceConfigurationUserStatusOverviewClient.Client)

	deviceConfigurationsAllManagedDeviceCertificateStateClient, err := deviceconfigurationsallmanageddevicecertificatestate.NewDeviceConfigurationsAllManagedDeviceCertificateStateClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DeviceConfigurationsAllManagedDeviceCertificateState client: %+v", err)
	}
	configureFunc(deviceConfigurationsAllManagedDeviceCertificateStateClient.Client)

	deviceCustomAttributeShellScriptAssignmentClient, err := devicecustomattributeshellscriptassignment.NewDeviceCustomAttributeShellScriptAssignmentClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DeviceCustomAttributeShellScriptAssignment client: %+v", err)
	}
	configureFunc(deviceCustomAttributeShellScriptAssignmentClient.Client)

	deviceCustomAttributeShellScriptClient, err := devicecustomattributeshellscript.NewDeviceCustomAttributeShellScriptClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DeviceCustomAttributeShellScript client: %+v", err)
	}
	configureFunc(deviceCustomAttributeShellScriptClient.Client)

	deviceCustomAttributeShellScriptDeviceRunStateClient, err := devicecustomattributeshellscriptdevicerunstate.NewDeviceCustomAttributeShellScriptDeviceRunStateClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DeviceCustomAttributeShellScriptDeviceRunState client: %+v", err)
	}
	configureFunc(deviceCustomAttributeShellScriptDeviceRunStateClient.Client)

	deviceCustomAttributeShellScriptDeviceRunStateManagedDeviceClient, err := devicecustomattributeshellscriptdevicerunstatemanageddevice.NewDeviceCustomAttributeShellScriptDeviceRunStateManagedDeviceClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DeviceCustomAttributeShellScriptDeviceRunStateManagedDevice client: %+v", err)
	}
	configureFunc(deviceCustomAttributeShellScriptDeviceRunStateManagedDeviceClient.Client)

	deviceCustomAttributeShellScriptGroupAssignmentClient, err := devicecustomattributeshellscriptgroupassignment.NewDeviceCustomAttributeShellScriptGroupAssignmentClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DeviceCustomAttributeShellScriptGroupAssignment client: %+v", err)
	}
	configureFunc(deviceCustomAttributeShellScriptGroupAssignmentClient.Client)

	deviceCustomAttributeShellScriptRunSummaryClient, err := devicecustomattributeshellscriptrunsummary.NewDeviceCustomAttributeShellScriptRunSummaryClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DeviceCustomAttributeShellScriptRunSummary client: %+v", err)
	}
	configureFunc(deviceCustomAttributeShellScriptRunSummaryClient.Client)

	deviceCustomAttributeShellScriptUserRunStateClient, err := devicecustomattributeshellscriptuserrunstate.NewDeviceCustomAttributeShellScriptUserRunStateClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DeviceCustomAttributeShellScriptUserRunState client: %+v", err)
	}
	configureFunc(deviceCustomAttributeShellScriptUserRunStateClient.Client)

	deviceCustomAttributeShellScriptUserRunStateDeviceRunStateClient, err := devicecustomattributeshellscriptuserrunstatedevicerunstate.NewDeviceCustomAttributeShellScriptUserRunStateDeviceRunStateClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DeviceCustomAttributeShellScriptUserRunStateDeviceRunState client: %+v", err)
	}
	configureFunc(deviceCustomAttributeShellScriptUserRunStateDeviceRunStateClient.Client)

	deviceCustomAttributeShellScriptUserRunStateDeviceRunStateManagedDeviceClient, err := devicecustomattributeshellscriptuserrunstatedevicerunstatemanageddevice.NewDeviceCustomAttributeShellScriptUserRunStateDeviceRunStateManagedDeviceClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DeviceCustomAttributeShellScriptUserRunStateDeviceRunStateManagedDevice client: %+v", err)
	}
	configureFunc(deviceCustomAttributeShellScriptUserRunStateDeviceRunStateManagedDeviceClient.Client)

	deviceEnrollmentConfigurationAssignmentClient, err := deviceenrollmentconfigurationassignment.NewDeviceEnrollmentConfigurationAssignmentClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DeviceEnrollmentConfigurationAssignment client: %+v", err)
	}
	configureFunc(deviceEnrollmentConfigurationAssignmentClient.Client)

	deviceEnrollmentConfigurationClient, err := deviceenrollmentconfiguration.NewDeviceEnrollmentConfigurationClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DeviceEnrollmentConfiguration client: %+v", err)
	}
	configureFunc(deviceEnrollmentConfigurationClient.Client)

	deviceHealthScriptAssignmentClient, err := devicehealthscriptassignment.NewDeviceHealthScriptAssignmentClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DeviceHealthScriptAssignment client: %+v", err)
	}
	configureFunc(deviceHealthScriptAssignmentClient.Client)

	deviceHealthScriptClient, err := devicehealthscript.NewDeviceHealthScriptClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DeviceHealthScript client: %+v", err)
	}
	configureFunc(deviceHealthScriptClient.Client)

	deviceHealthScriptDeviceRunStateClient, err := devicehealthscriptdevicerunstate.NewDeviceHealthScriptDeviceRunStateClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DeviceHealthScriptDeviceRunState client: %+v", err)
	}
	configureFunc(deviceHealthScriptDeviceRunStateClient.Client)

	deviceHealthScriptDeviceRunStateManagedDeviceClient, err := devicehealthscriptdevicerunstatemanageddevice.NewDeviceHealthScriptDeviceRunStateManagedDeviceClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DeviceHealthScriptDeviceRunStateManagedDevice client: %+v", err)
	}
	configureFunc(deviceHealthScriptDeviceRunStateManagedDeviceClient.Client)

	deviceHealthScriptRunSummaryClient, err := devicehealthscriptrunsummary.NewDeviceHealthScriptRunSummaryClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DeviceHealthScriptRunSummary client: %+v", err)
	}
	configureFunc(deviceHealthScriptRunSummaryClient.Client)

	deviceManagementClient, err := devicemanagement.NewDeviceManagementClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DeviceManagement client: %+v", err)
	}
	configureFunc(deviceManagementClient.Client)

	deviceManagementPartnerClient, err := devicemanagementpartner.NewDeviceManagementPartnerClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DeviceManagementPartner client: %+v", err)
	}
	configureFunc(deviceManagementPartnerClient.Client)

	deviceManagementScriptAssignmentClient, err := devicemanagementscriptassignment.NewDeviceManagementScriptAssignmentClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DeviceManagementScriptAssignment client: %+v", err)
	}
	configureFunc(deviceManagementScriptAssignmentClient.Client)

	deviceManagementScriptClient, err := devicemanagementscript.NewDeviceManagementScriptClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DeviceManagementScript client: %+v", err)
	}
	configureFunc(deviceManagementScriptClient.Client)

	deviceManagementScriptDeviceRunStateClient, err := devicemanagementscriptdevicerunstate.NewDeviceManagementScriptDeviceRunStateClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DeviceManagementScriptDeviceRunState client: %+v", err)
	}
	configureFunc(deviceManagementScriptDeviceRunStateClient.Client)

	deviceManagementScriptDeviceRunStateManagedDeviceClient, err := devicemanagementscriptdevicerunstatemanageddevice.NewDeviceManagementScriptDeviceRunStateManagedDeviceClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DeviceManagementScriptDeviceRunStateManagedDevice client: %+v", err)
	}
	configureFunc(deviceManagementScriptDeviceRunStateManagedDeviceClient.Client)

	deviceManagementScriptGroupAssignmentClient, err := devicemanagementscriptgroupassignment.NewDeviceManagementScriptGroupAssignmentClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DeviceManagementScriptGroupAssignment client: %+v", err)
	}
	configureFunc(deviceManagementScriptGroupAssignmentClient.Client)

	deviceManagementScriptRunSummaryClient, err := devicemanagementscriptrunsummary.NewDeviceManagementScriptRunSummaryClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DeviceManagementScriptRunSummary client: %+v", err)
	}
	configureFunc(deviceManagementScriptRunSummaryClient.Client)

	deviceManagementScriptUserRunStateClient, err := devicemanagementscriptuserrunstate.NewDeviceManagementScriptUserRunStateClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DeviceManagementScriptUserRunState client: %+v", err)
	}
	configureFunc(deviceManagementScriptUserRunStateClient.Client)

	deviceManagementScriptUserRunStateDeviceRunStateClient, err := devicemanagementscriptuserrunstatedevicerunstate.NewDeviceManagementScriptUserRunStateDeviceRunStateClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DeviceManagementScriptUserRunStateDeviceRunState client: %+v", err)
	}
	configureFunc(deviceManagementScriptUserRunStateDeviceRunStateClient.Client)

	deviceManagementScriptUserRunStateDeviceRunStateManagedDeviceClient, err := devicemanagementscriptuserrunstatedevicerunstatemanageddevice.NewDeviceManagementScriptUserRunStateDeviceRunStateManagedDeviceClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DeviceManagementScriptUserRunStateDeviceRunStateManagedDevice client: %+v", err)
	}
	configureFunc(deviceManagementScriptUserRunStateDeviceRunStateManagedDeviceClient.Client)

	deviceShellScriptAssignmentClient, err := deviceshellscriptassignment.NewDeviceShellScriptAssignmentClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DeviceShellScriptAssignment client: %+v", err)
	}
	configureFunc(deviceShellScriptAssignmentClient.Client)

	deviceShellScriptClient, err := deviceshellscript.NewDeviceShellScriptClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DeviceShellScript client: %+v", err)
	}
	configureFunc(deviceShellScriptClient.Client)

	deviceShellScriptDeviceRunStateClient, err := deviceshellscriptdevicerunstate.NewDeviceShellScriptDeviceRunStateClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DeviceShellScriptDeviceRunState client: %+v", err)
	}
	configureFunc(deviceShellScriptDeviceRunStateClient.Client)

	deviceShellScriptDeviceRunStateManagedDeviceClient, err := deviceshellscriptdevicerunstatemanageddevice.NewDeviceShellScriptDeviceRunStateManagedDeviceClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DeviceShellScriptDeviceRunStateManagedDevice client: %+v", err)
	}
	configureFunc(deviceShellScriptDeviceRunStateManagedDeviceClient.Client)

	deviceShellScriptGroupAssignmentClient, err := deviceshellscriptgroupassignment.NewDeviceShellScriptGroupAssignmentClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DeviceShellScriptGroupAssignment client: %+v", err)
	}
	configureFunc(deviceShellScriptGroupAssignmentClient.Client)

	deviceShellScriptRunSummaryClient, err := deviceshellscriptrunsummary.NewDeviceShellScriptRunSummaryClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DeviceShellScriptRunSummary client: %+v", err)
	}
	configureFunc(deviceShellScriptRunSummaryClient.Client)

	deviceShellScriptUserRunStateClient, err := deviceshellscriptuserrunstate.NewDeviceShellScriptUserRunStateClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DeviceShellScriptUserRunState client: %+v", err)
	}
	configureFunc(deviceShellScriptUserRunStateClient.Client)

	deviceShellScriptUserRunStateDeviceRunStateClient, err := deviceshellscriptuserrunstatedevicerunstate.NewDeviceShellScriptUserRunStateDeviceRunStateClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DeviceShellScriptUserRunStateDeviceRunState client: %+v", err)
	}
	configureFunc(deviceShellScriptUserRunStateDeviceRunStateClient.Client)

	deviceShellScriptUserRunStateDeviceRunStateManagedDeviceClient, err := deviceshellscriptuserrunstatedevicerunstatemanageddevice.NewDeviceShellScriptUserRunStateDeviceRunStateManagedDeviceClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DeviceShellScriptUserRunStateDeviceRunStateManagedDevice client: %+v", err)
	}
	configureFunc(deviceShellScriptUserRunStateDeviceRunStateManagedDeviceClient.Client)

	domainJoinConnectorClient, err := domainjoinconnector.NewDomainJoinConnectorClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DomainJoinConnector client: %+v", err)
	}
	configureFunc(domainJoinConnectorClient.Client)

	elevationRequestClient, err := elevationrequest.NewElevationRequestClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ElevationRequest client: %+v", err)
	}
	configureFunc(elevationRequestClient.Client)

	embeddedSIMActivationCodePoolAssignmentClient, err := embeddedsimactivationcodepoolassignment.NewEmbeddedSIMActivationCodePoolAssignmentClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building EmbeddedSIMActivationCodePoolAssignment client: %+v", err)
	}
	configureFunc(embeddedSIMActivationCodePoolAssignmentClient.Client)

	embeddedSIMActivationCodePoolClient, err := embeddedsimactivationcodepool.NewEmbeddedSIMActivationCodePoolClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building EmbeddedSIMActivationCodePool client: %+v", err)
	}
	configureFunc(embeddedSIMActivationCodePoolClient.Client)

	embeddedSIMActivationCodePoolDeviceStateClient, err := embeddedsimactivationcodepooldevicestate.NewEmbeddedSIMActivationCodePoolDeviceStateClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building EmbeddedSIMActivationCodePoolDeviceState client: %+v", err)
	}
	configureFunc(embeddedSIMActivationCodePoolDeviceStateClient.Client)

	endpointPrivilegeManagementProvisioningStatusClient, err := endpointprivilegemanagementprovisioningstatus.NewEndpointPrivilegeManagementProvisioningStatusClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building EndpointPrivilegeManagementProvisioningStatus client: %+v", err)
	}
	configureFunc(endpointPrivilegeManagementProvisioningStatusClient.Client)

	exchangeConnectorClient, err := exchangeconnector.NewExchangeConnectorClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ExchangeConnector client: %+v", err)
	}
	configureFunc(exchangeConnectorClient.Client)

	exchangeOnPremisesPolicyClient, err := exchangeonpremisespolicy.NewExchangeOnPremisesPolicyClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ExchangeOnPremisesPolicy client: %+v", err)
	}
	configureFunc(exchangeOnPremisesPolicyClient.Client)

	exchangeOnPremisesPolicyConditionalAccessSettingClient, err := exchangeonpremisespolicyconditionalaccesssetting.NewExchangeOnPremisesPolicyConditionalAccessSettingClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ExchangeOnPremisesPolicyConditionalAccessSetting client: %+v", err)
	}
	configureFunc(exchangeOnPremisesPolicyConditionalAccessSettingClient.Client)

	exportDeviceManagementReportsJobsClient, err := exportdevicemanagementreportsjobs.NewExportDeviceManagementReportsJobsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ExportDeviceManagementReportsJobs client: %+v", err)
	}
	configureFunc(exportDeviceManagementReportsJobsClient.Client)

	exportDeviceManagementVirtualEndpointReportsJobsClient, err := exportdevicemanagementvirtualendpointreportsjobs.NewExportDeviceManagementVirtualEndpointReportsJobsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ExportDeviceManagementVirtualEndpointReportsJobs client: %+v", err)
	}
	configureFunc(exportDeviceManagementVirtualEndpointReportsJobsClient.Client)

	groupPolicyCategoryChildClient, err := grouppolicycategorychild.NewGroupPolicyCategoryChildClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building GroupPolicyCategoryChild client: %+v", err)
	}
	configureFunc(groupPolicyCategoryChildClient.Client)

	groupPolicyCategoryClient, err := grouppolicycategory.NewGroupPolicyCategoryClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building GroupPolicyCategory client: %+v", err)
	}
	configureFunc(groupPolicyCategoryClient.Client)

	groupPolicyCategoryDefinitionClient, err := grouppolicycategorydefinition.NewGroupPolicyCategoryDefinitionClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building GroupPolicyCategoryDefinition client: %+v", err)
	}
	configureFunc(groupPolicyCategoryDefinitionClient.Client)

	groupPolicyCategoryDefinitionFileClient, err := grouppolicycategorydefinitionfile.NewGroupPolicyCategoryDefinitionFileClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building GroupPolicyCategoryDefinitionFile client: %+v", err)
	}
	configureFunc(groupPolicyCategoryDefinitionFileClient.Client)

	groupPolicyCategoryParentClient, err := grouppolicycategoryparent.NewGroupPolicyCategoryParentClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building GroupPolicyCategoryParent client: %+v", err)
	}
	configureFunc(groupPolicyCategoryParentClient.Client)

	groupPolicyConfigurationAssignmentClient, err := grouppolicyconfigurationassignment.NewGroupPolicyConfigurationAssignmentClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building GroupPolicyConfigurationAssignment client: %+v", err)
	}
	configureFunc(groupPolicyConfigurationAssignmentClient.Client)

	groupPolicyConfigurationClient, err := grouppolicyconfiguration.NewGroupPolicyConfigurationClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building GroupPolicyConfiguration client: %+v", err)
	}
	configureFunc(groupPolicyConfigurationClient.Client)

	groupPolicyConfigurationDefinitionValueClient, err := grouppolicyconfigurationdefinitionvalue.NewGroupPolicyConfigurationDefinitionValueClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building GroupPolicyConfigurationDefinitionValue client: %+v", err)
	}
	configureFunc(groupPolicyConfigurationDefinitionValueClient.Client)

	groupPolicyConfigurationDefinitionValueDefinitionClient, err := grouppolicyconfigurationdefinitionvaluedefinition.NewGroupPolicyConfigurationDefinitionValueDefinitionClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building GroupPolicyConfigurationDefinitionValueDefinition client: %+v", err)
	}
	configureFunc(groupPolicyConfigurationDefinitionValueDefinitionClient.Client)

	groupPolicyConfigurationDefinitionValuePresentationValueClient, err := grouppolicyconfigurationdefinitionvaluepresentationvalue.NewGroupPolicyConfigurationDefinitionValuePresentationValueClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building GroupPolicyConfigurationDefinitionValuePresentationValue client: %+v", err)
	}
	configureFunc(groupPolicyConfigurationDefinitionValuePresentationValueClient.Client)

	groupPolicyConfigurationDefinitionValuePresentationValueDefinitionValueClient, err := grouppolicyconfigurationdefinitionvaluepresentationvaluedefinitionvalue.NewGroupPolicyConfigurationDefinitionValuePresentationValueDefinitionValueClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building GroupPolicyConfigurationDefinitionValuePresentationValueDefinitionValue client: %+v", err)
	}
	configureFunc(groupPolicyConfigurationDefinitionValuePresentationValueDefinitionValueClient.Client)

	groupPolicyConfigurationDefinitionValuePresentationValuePresentationClient, err := grouppolicyconfigurationdefinitionvaluepresentationvaluepresentation.NewGroupPolicyConfigurationDefinitionValuePresentationValuePresentationClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building GroupPolicyConfigurationDefinitionValuePresentationValuePresentation client: %+v", err)
	}
	configureFunc(groupPolicyConfigurationDefinitionValuePresentationValuePresentationClient.Client)

	groupPolicyDefinitionCategoryClient, err := grouppolicydefinitioncategory.NewGroupPolicyDefinitionCategoryClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building GroupPolicyDefinitionCategory client: %+v", err)
	}
	configureFunc(groupPolicyDefinitionCategoryClient.Client)

	groupPolicyDefinitionClient, err := grouppolicydefinition.NewGroupPolicyDefinitionClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building GroupPolicyDefinition client: %+v", err)
	}
	configureFunc(groupPolicyDefinitionClient.Client)

	groupPolicyDefinitionDefinitionFileClient, err := grouppolicydefinitiondefinitionfile.NewGroupPolicyDefinitionDefinitionFileClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building GroupPolicyDefinitionDefinitionFile client: %+v", err)
	}
	configureFunc(groupPolicyDefinitionDefinitionFileClient.Client)

	groupPolicyDefinitionFileClient, err := grouppolicydefinitionfile.NewGroupPolicyDefinitionFileClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building GroupPolicyDefinitionFile client: %+v", err)
	}
	configureFunc(groupPolicyDefinitionFileClient.Client)

	groupPolicyDefinitionFileDefinitionClient, err := grouppolicydefinitionfiledefinition.NewGroupPolicyDefinitionFileDefinitionClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building GroupPolicyDefinitionFileDefinition client: %+v", err)
	}
	configureFunc(groupPolicyDefinitionFileDefinitionClient.Client)

	groupPolicyDefinitionNextVersionDefinitionCategoryClient, err := grouppolicydefinitionnextversiondefinitioncategory.NewGroupPolicyDefinitionNextVersionDefinitionCategoryClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building GroupPolicyDefinitionNextVersionDefinitionCategory client: %+v", err)
	}
	configureFunc(groupPolicyDefinitionNextVersionDefinitionCategoryClient.Client)

	groupPolicyDefinitionNextVersionDefinitionClient, err := grouppolicydefinitionnextversiondefinition.NewGroupPolicyDefinitionNextVersionDefinitionClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building GroupPolicyDefinitionNextVersionDefinition client: %+v", err)
	}
	configureFunc(groupPolicyDefinitionNextVersionDefinitionClient.Client)

	groupPolicyDefinitionNextVersionDefinitionDefinitionFileClient, err := grouppolicydefinitionnextversiondefinitiondefinitionfile.NewGroupPolicyDefinitionNextVersionDefinitionDefinitionFileClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building GroupPolicyDefinitionNextVersionDefinitionDefinitionFile client: %+v", err)
	}
	configureFunc(groupPolicyDefinitionNextVersionDefinitionDefinitionFileClient.Client)

	groupPolicyDefinitionNextVersionDefinitionPresentationClient, err := grouppolicydefinitionnextversiondefinitionpresentation.NewGroupPolicyDefinitionNextVersionDefinitionPresentationClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building GroupPolicyDefinitionNextVersionDefinitionPresentation client: %+v", err)
	}
	configureFunc(groupPolicyDefinitionNextVersionDefinitionPresentationClient.Client)

	groupPolicyDefinitionNextVersionDefinitionPresentationDefinitionClient, err := grouppolicydefinitionnextversiondefinitionpresentationdefinition.NewGroupPolicyDefinitionNextVersionDefinitionPresentationDefinitionClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building GroupPolicyDefinitionNextVersionDefinitionPresentationDefinition client: %+v", err)
	}
	configureFunc(groupPolicyDefinitionNextVersionDefinitionPresentationDefinitionClient.Client)

	groupPolicyDefinitionNextVersionDefinitionPreviousVersionDefinitionCategoryClient, err := grouppolicydefinitionnextversiondefinitionpreviousversiondefinitioncategory.NewGroupPolicyDefinitionNextVersionDefinitionPreviousVersionDefinitionCategoryClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building GroupPolicyDefinitionNextVersionDefinitionPreviousVersionDefinitionCategory client: %+v", err)
	}
	configureFunc(groupPolicyDefinitionNextVersionDefinitionPreviousVersionDefinitionCategoryClient.Client)

	groupPolicyDefinitionNextVersionDefinitionPreviousVersionDefinitionClient, err := grouppolicydefinitionnextversiondefinitionpreviousversiondefinition.NewGroupPolicyDefinitionNextVersionDefinitionPreviousVersionDefinitionClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building GroupPolicyDefinitionNextVersionDefinitionPreviousVersionDefinition client: %+v", err)
	}
	configureFunc(groupPolicyDefinitionNextVersionDefinitionPreviousVersionDefinitionClient.Client)

	groupPolicyDefinitionNextVersionDefinitionPreviousVersionDefinitionDefinitionFileClient, err := grouppolicydefinitionnextversiondefinitionpreviousversiondefinitiondefinitionfile.NewGroupPolicyDefinitionNextVersionDefinitionPreviousVersionDefinitionDefinitionFileClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building GroupPolicyDefinitionNextVersionDefinitionPreviousVersionDefinitionDefinitionFile client: %+v", err)
	}
	configureFunc(groupPolicyDefinitionNextVersionDefinitionPreviousVersionDefinitionDefinitionFileClient.Client)

	groupPolicyDefinitionNextVersionDefinitionPreviousVersionDefinitionPresentationClient, err := grouppolicydefinitionnextversiondefinitionpreviousversiondefinitionpresentation.NewGroupPolicyDefinitionNextVersionDefinitionPreviousVersionDefinitionPresentationClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building GroupPolicyDefinitionNextVersionDefinitionPreviousVersionDefinitionPresentation client: %+v", err)
	}
	configureFunc(groupPolicyDefinitionNextVersionDefinitionPreviousVersionDefinitionPresentationClient.Client)

	groupPolicyDefinitionNextVersionDefinitionPreviousVersionDefinitionPresentationDefinitionClient, err := grouppolicydefinitionnextversiondefinitionpreviousversiondefinitionpresentationdefinition.NewGroupPolicyDefinitionNextVersionDefinitionPreviousVersionDefinitionPresentationDefinitionClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building GroupPolicyDefinitionNextVersionDefinitionPreviousVersionDefinitionPresentationDefinition client: %+v", err)
	}
	configureFunc(groupPolicyDefinitionNextVersionDefinitionPreviousVersionDefinitionPresentationDefinitionClient.Client)

	groupPolicyDefinitionPresentationClient, err := grouppolicydefinitionpresentation.NewGroupPolicyDefinitionPresentationClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building GroupPolicyDefinitionPresentation client: %+v", err)
	}
	configureFunc(groupPolicyDefinitionPresentationClient.Client)

	groupPolicyDefinitionPresentationDefinitionClient, err := grouppolicydefinitionpresentationdefinition.NewGroupPolicyDefinitionPresentationDefinitionClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building GroupPolicyDefinitionPresentationDefinition client: %+v", err)
	}
	configureFunc(groupPolicyDefinitionPresentationDefinitionClient.Client)

	groupPolicyDefinitionPreviousVersionDefinitionCategoryClient, err := grouppolicydefinitionpreviousversiondefinitioncategory.NewGroupPolicyDefinitionPreviousVersionDefinitionCategoryClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building GroupPolicyDefinitionPreviousVersionDefinitionCategory client: %+v", err)
	}
	configureFunc(groupPolicyDefinitionPreviousVersionDefinitionCategoryClient.Client)

	groupPolicyDefinitionPreviousVersionDefinitionClient, err := grouppolicydefinitionpreviousversiondefinition.NewGroupPolicyDefinitionPreviousVersionDefinitionClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building GroupPolicyDefinitionPreviousVersionDefinition client: %+v", err)
	}
	configureFunc(groupPolicyDefinitionPreviousVersionDefinitionClient.Client)

	groupPolicyDefinitionPreviousVersionDefinitionDefinitionFileClient, err := grouppolicydefinitionpreviousversiondefinitiondefinitionfile.NewGroupPolicyDefinitionPreviousVersionDefinitionDefinitionFileClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building GroupPolicyDefinitionPreviousVersionDefinitionDefinitionFile client: %+v", err)
	}
	configureFunc(groupPolicyDefinitionPreviousVersionDefinitionDefinitionFileClient.Client)

	groupPolicyDefinitionPreviousVersionDefinitionNextVersionDefinitionCategoryClient, err := grouppolicydefinitionpreviousversiondefinitionnextversiondefinitioncategory.NewGroupPolicyDefinitionPreviousVersionDefinitionNextVersionDefinitionCategoryClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building GroupPolicyDefinitionPreviousVersionDefinitionNextVersionDefinitionCategory client: %+v", err)
	}
	configureFunc(groupPolicyDefinitionPreviousVersionDefinitionNextVersionDefinitionCategoryClient.Client)

	groupPolicyDefinitionPreviousVersionDefinitionNextVersionDefinitionClient, err := grouppolicydefinitionpreviousversiondefinitionnextversiondefinition.NewGroupPolicyDefinitionPreviousVersionDefinitionNextVersionDefinitionClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building GroupPolicyDefinitionPreviousVersionDefinitionNextVersionDefinition client: %+v", err)
	}
	configureFunc(groupPolicyDefinitionPreviousVersionDefinitionNextVersionDefinitionClient.Client)

	groupPolicyDefinitionPreviousVersionDefinitionNextVersionDefinitionDefinitionFileClient, err := grouppolicydefinitionpreviousversiondefinitionnextversiondefinitiondefinitionfile.NewGroupPolicyDefinitionPreviousVersionDefinitionNextVersionDefinitionDefinitionFileClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building GroupPolicyDefinitionPreviousVersionDefinitionNextVersionDefinitionDefinitionFile client: %+v", err)
	}
	configureFunc(groupPolicyDefinitionPreviousVersionDefinitionNextVersionDefinitionDefinitionFileClient.Client)

	groupPolicyDefinitionPreviousVersionDefinitionNextVersionDefinitionPresentationClient, err := grouppolicydefinitionpreviousversiondefinitionnextversiondefinitionpresentation.NewGroupPolicyDefinitionPreviousVersionDefinitionNextVersionDefinitionPresentationClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building GroupPolicyDefinitionPreviousVersionDefinitionNextVersionDefinitionPresentation client: %+v", err)
	}
	configureFunc(groupPolicyDefinitionPreviousVersionDefinitionNextVersionDefinitionPresentationClient.Client)

	groupPolicyDefinitionPreviousVersionDefinitionNextVersionDefinitionPresentationDefinitionClient, err := grouppolicydefinitionpreviousversiondefinitionnextversiondefinitionpresentationdefinition.NewGroupPolicyDefinitionPreviousVersionDefinitionNextVersionDefinitionPresentationDefinitionClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building GroupPolicyDefinitionPreviousVersionDefinitionNextVersionDefinitionPresentationDefinition client: %+v", err)
	}
	configureFunc(groupPolicyDefinitionPreviousVersionDefinitionNextVersionDefinitionPresentationDefinitionClient.Client)

	groupPolicyDefinitionPreviousVersionDefinitionPresentationClient, err := grouppolicydefinitionpreviousversiondefinitionpresentation.NewGroupPolicyDefinitionPreviousVersionDefinitionPresentationClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building GroupPolicyDefinitionPreviousVersionDefinitionPresentation client: %+v", err)
	}
	configureFunc(groupPolicyDefinitionPreviousVersionDefinitionPresentationClient.Client)

	groupPolicyDefinitionPreviousVersionDefinitionPresentationDefinitionClient, err := grouppolicydefinitionpreviousversiondefinitionpresentationdefinition.NewGroupPolicyDefinitionPreviousVersionDefinitionPresentationDefinitionClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building GroupPolicyDefinitionPreviousVersionDefinitionPresentationDefinition client: %+v", err)
	}
	configureFunc(groupPolicyDefinitionPreviousVersionDefinitionPresentationDefinitionClient.Client)

	groupPolicyMigrationReportClient, err := grouppolicymigrationreport.NewGroupPolicyMigrationReportClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building GroupPolicyMigrationReport client: %+v", err)
	}
	configureFunc(groupPolicyMigrationReportClient.Client)

	groupPolicyMigrationReportGroupPolicySettingMappingClient, err := grouppolicymigrationreportgrouppolicysettingmapping.NewGroupPolicyMigrationReportGroupPolicySettingMappingClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building GroupPolicyMigrationReportGroupPolicySettingMapping client: %+v", err)
	}
	configureFunc(groupPolicyMigrationReportGroupPolicySettingMappingClient.Client)

	groupPolicyMigrationReportUnsupportedGroupPolicyExtensionClient, err := grouppolicymigrationreportunsupportedgrouppolicyextension.NewGroupPolicyMigrationReportUnsupportedGroupPolicyExtensionClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building GroupPolicyMigrationReportUnsupportedGroupPolicyExtension client: %+v", err)
	}
	configureFunc(groupPolicyMigrationReportUnsupportedGroupPolicyExtensionClient.Client)

	groupPolicyObjectFileClient, err := grouppolicyobjectfile.NewGroupPolicyObjectFileClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building GroupPolicyObjectFile client: %+v", err)
	}
	configureFunc(groupPolicyObjectFileClient.Client)

	groupPolicyUploadedDefinitionFileClient, err := grouppolicyuploadeddefinitionfile.NewGroupPolicyUploadedDefinitionFileClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building GroupPolicyUploadedDefinitionFile client: %+v", err)
	}
	configureFunc(groupPolicyUploadedDefinitionFileClient.Client)

	groupPolicyUploadedDefinitionFileDefinitionClient, err := grouppolicyuploadeddefinitionfiledefinition.NewGroupPolicyUploadedDefinitionFileDefinitionClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building GroupPolicyUploadedDefinitionFileDefinition client: %+v", err)
	}
	configureFunc(groupPolicyUploadedDefinitionFileDefinitionClient.Client)

	groupPolicyUploadedDefinitionFileGroupPolicyOperationClient, err := grouppolicyuploadeddefinitionfilegrouppolicyoperation.NewGroupPolicyUploadedDefinitionFileGroupPolicyOperationClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building GroupPolicyUploadedDefinitionFileGroupPolicyOperation client: %+v", err)
	}
	configureFunc(groupPolicyUploadedDefinitionFileGroupPolicyOperationClient.Client)

	hardwareConfigurationAssignmentClient, err := hardwareconfigurationassignment.NewHardwareConfigurationAssignmentClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building HardwareConfigurationAssignment client: %+v", err)
	}
	configureFunc(hardwareConfigurationAssignmentClient.Client)

	hardwareConfigurationClient, err := hardwareconfiguration.NewHardwareConfigurationClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building HardwareConfiguration client: %+v", err)
	}
	configureFunc(hardwareConfigurationClient.Client)

	hardwareConfigurationDeviceRunStateClient, err := hardwareconfigurationdevicerunstate.NewHardwareConfigurationDeviceRunStateClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building HardwareConfigurationDeviceRunState client: %+v", err)
	}
	configureFunc(hardwareConfigurationDeviceRunStateClient.Client)

	hardwareConfigurationRunSummaryClient, err := hardwareconfigurationrunsummary.NewHardwareConfigurationRunSummaryClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building HardwareConfigurationRunSummary client: %+v", err)
	}
	configureFunc(hardwareConfigurationRunSummaryClient.Client)

	hardwareConfigurationUserRunStateClient, err := hardwareconfigurationuserrunstate.NewHardwareConfigurationUserRunStateClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building HardwareConfigurationUserRunState client: %+v", err)
	}
	configureFunc(hardwareConfigurationUserRunStateClient.Client)

	hardwarePasswordDetailClient, err := hardwarepassworddetail.NewHardwarePasswordDetailClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building HardwarePasswordDetail client: %+v", err)
	}
	configureFunc(hardwarePasswordDetailClient.Client)

	hardwarePasswordInfoClient, err := hardwarepasswordinfo.NewHardwarePasswordInfoClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building HardwarePasswordInfo client: %+v", err)
	}
	configureFunc(hardwarePasswordInfoClient.Client)

	importedDeviceIdentityClient, err := importeddeviceidentity.NewImportedDeviceIdentityClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ImportedDeviceIdentity client: %+v", err)
	}
	configureFunc(importedDeviceIdentityClient.Client)

	importedWindowsAutopilotDeviceIdentityClient, err := importedwindowsautopilotdeviceidentity.NewImportedWindowsAutopilotDeviceIdentityClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ImportedWindowsAutopilotDeviceIdentity client: %+v", err)
	}
	configureFunc(importedWindowsAutopilotDeviceIdentityClient.Client)

	intentAssignmentClient, err := intentassignment.NewIntentAssignmentClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building IntentAssignment client: %+v", err)
	}
	configureFunc(intentAssignmentClient.Client)

	intentCategoryClient, err := intentcategory.NewIntentCategoryClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building IntentCategory client: %+v", err)
	}
	configureFunc(intentCategoryClient.Client)

	intentCategorySettingClient, err := intentcategorysetting.NewIntentCategorySettingClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building IntentCategorySetting client: %+v", err)
	}
	configureFunc(intentCategorySettingClient.Client)

	intentCategorySettingDefinitionClient, err := intentcategorysettingdefinition.NewIntentCategorySettingDefinitionClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building IntentCategorySettingDefinition client: %+v", err)
	}
	configureFunc(intentCategorySettingDefinitionClient.Client)

	intentClient, err := intent.NewIntentClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Intent client: %+v", err)
	}
	configureFunc(intentClient.Client)

	intentDeviceSettingStateSummaryClient, err := intentdevicesettingstatesummary.NewIntentDeviceSettingStateSummaryClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building IntentDeviceSettingStateSummary client: %+v", err)
	}
	configureFunc(intentDeviceSettingStateSummaryClient.Client)

	intentDeviceStateClient, err := intentdevicestate.NewIntentDeviceStateClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building IntentDeviceState client: %+v", err)
	}
	configureFunc(intentDeviceStateClient.Client)

	intentDeviceStateSummaryClient, err := intentdevicestatesummary.NewIntentDeviceStateSummaryClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building IntentDeviceStateSummary client: %+v", err)
	}
	configureFunc(intentDeviceStateSummaryClient.Client)

	intentSettingClient, err := intentsetting.NewIntentSettingClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building IntentSetting client: %+v", err)
	}
	configureFunc(intentSettingClient.Client)

	intentUserStateClient, err := intentuserstate.NewIntentUserStateClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building IntentUserState client: %+v", err)
	}
	configureFunc(intentUserStateClient.Client)

	intentUserStateSummaryClient, err := intentuserstatesummary.NewIntentUserStateSummaryClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building IntentUserStateSummary client: %+v", err)
	}
	configureFunc(intentUserStateSummaryClient.Client)

	intuneBrandingProfileAssignmentClient, err := intunebrandingprofileassignment.NewIntuneBrandingProfileAssignmentClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building IntuneBrandingProfileAssignment client: %+v", err)
	}
	configureFunc(intuneBrandingProfileAssignmentClient.Client)

	intuneBrandingProfileClient, err := intunebrandingprofile.NewIntuneBrandingProfileClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building IntuneBrandingProfile client: %+v", err)
	}
	configureFunc(intuneBrandingProfileClient.Client)

	iosUpdateStatusClient, err := iosupdatestatus.NewIosUpdateStatusClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building IosUpdateStatus client: %+v", err)
	}
	configureFunc(iosUpdateStatusClient.Client)

	macOSSoftwareUpdateAccountSummaryCategorySummaryClient, err := macossoftwareupdateaccountsummarycategorysummary.NewMacOSSoftwareUpdateAccountSummaryCategorySummaryClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MacOSSoftwareUpdateAccountSummaryCategorySummary client: %+v", err)
	}
	configureFunc(macOSSoftwareUpdateAccountSummaryCategorySummaryClient.Client)

	macOSSoftwareUpdateAccountSummaryClient, err := macossoftwareupdateaccountsummary.NewMacOSSoftwareUpdateAccountSummaryClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MacOSSoftwareUpdateAccountSummary client: %+v", err)
	}
	configureFunc(macOSSoftwareUpdateAccountSummaryClient.Client)

	managedDeviceAssignmentFilterEvaluationStatusDetailClient, err := manageddeviceassignmentfilterevaluationstatusdetail.NewManagedDeviceAssignmentFilterEvaluationStatusDetailClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ManagedDeviceAssignmentFilterEvaluationStatusDetail client: %+v", err)
	}
	configureFunc(managedDeviceAssignmentFilterEvaluationStatusDetailClient.Client)

	managedDeviceCleanupRuleClient, err := manageddevicecleanuprule.NewManagedDeviceCleanupRuleClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ManagedDeviceCleanupRule client: %+v", err)
	}
	configureFunc(managedDeviceCleanupRuleClient.Client)

	managedDeviceClient, err := manageddevice.NewManagedDeviceClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ManagedDevice client: %+v", err)
	}
	configureFunc(managedDeviceClient.Client)

	managedDeviceDetectedAppClient, err := manageddevicedetectedapp.NewManagedDeviceDetectedAppClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ManagedDeviceDetectedApp client: %+v", err)
	}
	configureFunc(managedDeviceDetectedAppClient.Client)

	managedDeviceDeviceCategoryClient, err := manageddevicedevicecategory.NewManagedDeviceDeviceCategoryClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ManagedDeviceDeviceCategory client: %+v", err)
	}
	configureFunc(managedDeviceDeviceCategoryClient.Client)

	managedDeviceDeviceCompliancePolicyStateClient, err := manageddevicedevicecompliancepolicystate.NewManagedDeviceDeviceCompliancePolicyStateClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ManagedDeviceDeviceCompliancePolicyState client: %+v", err)
	}
	configureFunc(managedDeviceDeviceCompliancePolicyStateClient.Client)

	managedDeviceDeviceConfigurationStateClient, err := manageddevicedeviceconfigurationstate.NewManagedDeviceDeviceConfigurationStateClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ManagedDeviceDeviceConfigurationState client: %+v", err)
	}
	configureFunc(managedDeviceDeviceConfigurationStateClient.Client)

	managedDeviceDeviceHealthScriptStateClient, err := manageddevicedevicehealthscriptstate.NewManagedDeviceDeviceHealthScriptStateClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ManagedDeviceDeviceHealthScriptState client: %+v", err)
	}
	configureFunc(managedDeviceDeviceHealthScriptStateClient.Client)

	managedDeviceDeviceHealthScriptStateIdIdPolicyIdPolicyIdDeviceIdDeviceIdClient, err := manageddevicedevicehealthscriptstateididpolicyidpolicyiddeviceiddeviceid.NewManagedDeviceDeviceHealthScriptStateIdIdPolicyIdPolicyIdDeviceIdDeviceIdClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ManagedDeviceDeviceHealthScriptStateIdIdPolicyIdPolicyIdDeviceIdDeviceId client: %+v", err)
	}
	configureFunc(managedDeviceDeviceHealthScriptStateIdIdPolicyIdPolicyIdDeviceIdDeviceIdClient.Client)

	managedDeviceEncryptionStateClient, err := manageddeviceencryptionstate.NewManagedDeviceEncryptionStateClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ManagedDeviceEncryptionState client: %+v", err)
	}
	configureFunc(managedDeviceEncryptionStateClient.Client)

	managedDeviceLogCollectionRequestClient, err := manageddevicelogcollectionrequest.NewManagedDeviceLogCollectionRequestClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ManagedDeviceLogCollectionRequest client: %+v", err)
	}
	configureFunc(managedDeviceLogCollectionRequestClient.Client)

	managedDeviceManagedDeviceMobileAppConfigurationStateClient, err := manageddevicemanageddevicemobileappconfigurationstate.NewManagedDeviceManagedDeviceMobileAppConfigurationStateClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ManagedDeviceManagedDeviceMobileAppConfigurationState client: %+v", err)
	}
	configureFunc(managedDeviceManagedDeviceMobileAppConfigurationStateClient.Client)

	managedDeviceOverviewClient, err := manageddeviceoverview.NewManagedDeviceOverviewClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ManagedDeviceOverview client: %+v", err)
	}
	configureFunc(managedDeviceOverviewClient.Client)

	managedDeviceSecurityBaselineStateClient, err := manageddevicesecuritybaselinestate.NewManagedDeviceSecurityBaselineStateClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ManagedDeviceSecurityBaselineState client: %+v", err)
	}
	configureFunc(managedDeviceSecurityBaselineStateClient.Client)

	managedDeviceSecurityBaselineStateSettingStateClient, err := manageddevicesecuritybaselinestatesettingstate.NewManagedDeviceSecurityBaselineStateSettingStateClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ManagedDeviceSecurityBaselineStateSettingState client: %+v", err)
	}
	configureFunc(managedDeviceSecurityBaselineStateSettingStateClient.Client)

	managedDeviceUserClient, err := manageddeviceuser.NewManagedDeviceUserClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ManagedDeviceUser client: %+v", err)
	}
	configureFunc(managedDeviceUserClient.Client)

	managedDeviceWindowsOSImageClient, err := manageddevicewindowsosimage.NewManagedDeviceWindowsOSImageClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ManagedDeviceWindowsOSImage client: %+v", err)
	}
	configureFunc(managedDeviceWindowsOSImageClient.Client)

	managedDeviceWindowsProtectionStateClient, err := manageddevicewindowsprotectionstate.NewManagedDeviceWindowsProtectionStateClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ManagedDeviceWindowsProtectionState client: %+v", err)
	}
	configureFunc(managedDeviceWindowsProtectionStateClient.Client)

	managedDeviceWindowsProtectionStateDetectedMalwareStateClient, err := manageddevicewindowsprotectionstatedetectedmalwarestate.NewManagedDeviceWindowsProtectionStateDetectedMalwareStateClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ManagedDeviceWindowsProtectionStateDetectedMalwareState client: %+v", err)
	}
	configureFunc(managedDeviceWindowsProtectionStateDetectedMalwareStateClient.Client)

	microsoftTunnelConfigurationClient, err := microsofttunnelconfiguration.NewMicrosoftTunnelConfigurationClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MicrosoftTunnelConfiguration client: %+v", err)
	}
	configureFunc(microsoftTunnelConfigurationClient.Client)

	microsoftTunnelHealthThresholdClient, err := microsofttunnelhealththreshold.NewMicrosoftTunnelHealthThresholdClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MicrosoftTunnelHealthThreshold client: %+v", err)
	}
	configureFunc(microsoftTunnelHealthThresholdClient.Client)

	microsoftTunnelServerLogCollectionResponseClient, err := microsofttunnelserverlogcollectionresponse.NewMicrosoftTunnelServerLogCollectionResponseClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MicrosoftTunnelServerLogCollectionResponse client: %+v", err)
	}
	configureFunc(microsoftTunnelServerLogCollectionResponseClient.Client)

	microsoftTunnelSiteClient, err := microsofttunnelsite.NewMicrosoftTunnelSiteClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MicrosoftTunnelSite client: %+v", err)
	}
	configureFunc(microsoftTunnelSiteClient.Client)

	microsoftTunnelSiteMicrosoftTunnelConfigurationClient, err := microsofttunnelsitemicrosofttunnelconfiguration.NewMicrosoftTunnelSiteMicrosoftTunnelConfigurationClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MicrosoftTunnelSiteMicrosoftTunnelConfiguration client: %+v", err)
	}
	configureFunc(microsoftTunnelSiteMicrosoftTunnelConfigurationClient.Client)

	microsoftTunnelSiteMicrosoftTunnelServerClient, err := microsofttunnelsitemicrosofttunnelserver.NewMicrosoftTunnelSiteMicrosoftTunnelServerClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MicrosoftTunnelSiteMicrosoftTunnelServer client: %+v", err)
	}
	configureFunc(microsoftTunnelSiteMicrosoftTunnelServerClient.Client)

	mobileAppTroubleshootingEventAppLogCollectionRequestClient, err := mobileapptroubleshootingeventapplogcollectionrequest.NewMobileAppTroubleshootingEventAppLogCollectionRequestClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MobileAppTroubleshootingEventAppLogCollectionRequest client: %+v", err)
	}
	configureFunc(mobileAppTroubleshootingEventAppLogCollectionRequestClient.Client)

	mobileAppTroubleshootingEventClient, err := mobileapptroubleshootingevent.NewMobileAppTroubleshootingEventClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MobileAppTroubleshootingEvent client: %+v", err)
	}
	configureFunc(mobileAppTroubleshootingEventClient.Client)

	mobileThreatDefenseConnectorClient, err := mobilethreatdefenseconnector.NewMobileThreatDefenseConnectorClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MobileThreatDefenseConnector client: %+v", err)
	}
	configureFunc(mobileThreatDefenseConnectorClient.Client)

	monitoringAlertRecordClient, err := monitoringalertrecord.NewMonitoringAlertRecordClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MonitoringAlertRecord client: %+v", err)
	}
	configureFunc(monitoringAlertRecordClient.Client)

	monitoringAlertRuleClient, err := monitoringalertrule.NewMonitoringAlertRuleClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MonitoringAlertRule client: %+v", err)
	}
	configureFunc(monitoringAlertRuleClient.Client)

	monitoringClient, err := monitoring.NewMonitoringClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Monitoring client: %+v", err)
	}
	configureFunc(monitoringClient.Client)

	ndesConnectorClient, err := ndesconnector.NewNdesConnectorClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building NdesConnector client: %+v", err)
	}
	configureFunc(ndesConnectorClient.Client)

	notificationMessageTemplateClient, err := notificationmessagetemplate.NewNotificationMessageTemplateClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building NotificationMessageTemplate client: %+v", err)
	}
	configureFunc(notificationMessageTemplateClient.Client)

	notificationMessageTemplateLocalizedNotificationMessageClient, err := notificationmessagetemplatelocalizednotificationmessage.NewNotificationMessageTemplateLocalizedNotificationMessageClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building NotificationMessageTemplateLocalizedNotificationMessage client: %+v", err)
	}
	configureFunc(notificationMessageTemplateLocalizedNotificationMessageClient.Client)

	operationApprovalPolicyClient, err := operationapprovalpolicy.NewOperationApprovalPolicyClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building OperationApprovalPolicy client: %+v", err)
	}
	configureFunc(operationApprovalPolicyClient.Client)

	operationApprovalRequestClient, err := operationapprovalrequest.NewOperationApprovalRequestClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building OperationApprovalRequest client: %+v", err)
	}
	configureFunc(operationApprovalRequestClient.Client)

	privilegeManagementElevationClient, err := privilegemanagementelevation.NewPrivilegeManagementElevationClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building PrivilegeManagementElevation client: %+v", err)
	}
	configureFunc(privilegeManagementElevationClient.Client)

	rebootDeviceManagementUserExperienceAnalyticsBaselineAnalyticsMetricsClient, err := rebootdevicemanagementuserexperienceanalyticsbaselineanalyticsmetrics.NewRebootDeviceManagementUserExperienceAnalyticsBaselineAnalyticsMetricsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building RebootDeviceManagementUserExperienceAnalyticsBaselineAnalyticsMetrics client: %+v", err)
	}
	configureFunc(rebootDeviceManagementUserExperienceAnalyticsBaselineAnalyticsMetricsClient.Client)

	remoteActionAuditClient, err := remoteactionaudit.NewRemoteActionAuditClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building RemoteActionAudit client: %+v", err)
	}
	configureFunc(remoteActionAuditClient.Client)

	remoteAssistancePartnerClient, err := remoteassistancepartner.NewRemoteAssistancePartnerClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building RemoteAssistancePartner client: %+v", err)
	}
	configureFunc(remoteAssistancePartnerClient.Client)

	remoteAssistanceSettingClient, err := remoteassistancesetting.NewRemoteAssistanceSettingClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building RemoteAssistanceSetting client: %+v", err)
	}
	configureFunc(remoteAssistanceSettingClient.Client)

	reportCachedReportConfigurationClient, err := reportcachedreportconfiguration.NewReportCachedReportConfigurationClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ReportCachedReportConfiguration client: %+v", err)
	}
	configureFunc(reportCachedReportConfigurationClient.Client)

	reportClient, err := report.NewReportClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Report client: %+v", err)
	}
	configureFunc(reportClient.Client)

	resourceAccessProfileAssignmentClient, err := resourceaccessprofileassignment.NewResourceAccessProfileAssignmentClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ResourceAccessProfileAssignment client: %+v", err)
	}
	configureFunc(resourceAccessProfileAssignmentClient.Client)

	resourceAccessProfileClient, err := resourceaccessprofile.NewResourceAccessProfileClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ResourceAccessProfile client: %+v", err)
	}
	configureFunc(resourceAccessProfileClient.Client)

	resourceOperationClient, err := resourceoperation.NewResourceOperationClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ResourceOperation client: %+v", err)
	}
	configureFunc(resourceOperationClient.Client)

	reusablePolicySettingClient, err := reusablepolicysetting.NewReusablePolicySettingClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ReusablePolicySetting client: %+v", err)
	}
	configureFunc(reusablePolicySettingClient.Client)

	reusablePolicySettingReferencingConfigurationPolicyAssignmentClient, err := reusablepolicysettingreferencingconfigurationpolicyassignment.NewReusablePolicySettingReferencingConfigurationPolicyAssignmentClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ReusablePolicySettingReferencingConfigurationPolicyAssignment client: %+v", err)
	}
	configureFunc(reusablePolicySettingReferencingConfigurationPolicyAssignmentClient.Client)

	reusablePolicySettingReferencingConfigurationPolicyClient, err := reusablepolicysettingreferencingconfigurationpolicy.NewReusablePolicySettingReferencingConfigurationPolicyClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ReusablePolicySettingReferencingConfigurationPolicy client: %+v", err)
	}
	configureFunc(reusablePolicySettingReferencingConfigurationPolicyClient.Client)

	reusablePolicySettingReferencingConfigurationPolicySettingClient, err := reusablepolicysettingreferencingconfigurationpolicysetting.NewReusablePolicySettingReferencingConfigurationPolicySettingClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ReusablePolicySettingReferencingConfigurationPolicySetting client: %+v", err)
	}
	configureFunc(reusablePolicySettingReferencingConfigurationPolicySettingClient.Client)

	reusablePolicySettingReferencingConfigurationPolicySettingSettingDefinitionClient, err := reusablepolicysettingreferencingconfigurationpolicysettingsettingdefinition.NewReusablePolicySettingReferencingConfigurationPolicySettingSettingDefinitionClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ReusablePolicySettingReferencingConfigurationPolicySettingSettingDefinition client: %+v", err)
	}
	configureFunc(reusablePolicySettingReferencingConfigurationPolicySettingSettingDefinitionClient.Client)

	reusableSettingClient, err := reusablesetting.NewReusableSettingClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ReusableSetting client: %+v", err)
	}
	configureFunc(reusableSettingClient.Client)

	roleAssignmentClient, err := roleassignment.NewRoleAssignmentClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building RoleAssignment client: %+v", err)
	}
	configureFunc(roleAssignmentClient.Client)

	roleAssignmentRoleDefinitionClient, err := roleassignmentroledefinition.NewRoleAssignmentRoleDefinitionClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building RoleAssignmentRoleDefinition client: %+v", err)
	}
	configureFunc(roleAssignmentRoleDefinitionClient.Client)

	roleAssignmentRoleScopeTagClient, err := roleassignmentrolescopetag.NewRoleAssignmentRoleScopeTagClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building RoleAssignmentRoleScopeTag client: %+v", err)
	}
	configureFunc(roleAssignmentRoleScopeTagClient.Client)

	roleDefinitionClient, err := roledefinition.NewRoleDefinitionClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building RoleDefinition client: %+v", err)
	}
	configureFunc(roleDefinitionClient.Client)

	roleDefinitionRoleAssignmentClient, err := roledefinitionroleassignment.NewRoleDefinitionRoleAssignmentClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building RoleDefinitionRoleAssignment client: %+v", err)
	}
	configureFunc(roleDefinitionRoleAssignmentClient.Client)

	roleDefinitionRoleAssignmentRoleDefinitionClient, err := roledefinitionroleassignmentroledefinition.NewRoleDefinitionRoleAssignmentRoleDefinitionClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building RoleDefinitionRoleAssignmentRoleDefinition client: %+v", err)
	}
	configureFunc(roleDefinitionRoleAssignmentRoleDefinitionClient.Client)

	roleScopeTagAssignmentClient, err := rolescopetagassignment.NewRoleScopeTagAssignmentClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building RoleScopeTagAssignment client: %+v", err)
	}
	configureFunc(roleScopeTagAssignmentClient.Client)

	roleScopeTagClient, err := rolescopetag.NewRoleScopeTagClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building RoleScopeTag client: %+v", err)
	}
	configureFunc(roleScopeTagClient.Client)

	serviceNowConnectionClient, err := servicenowconnection.NewServiceNowConnectionClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ServiceNowConnection client: %+v", err)
	}
	configureFunc(serviceNowConnectionClient.Client)

	settingDefinitionClient, err := settingdefinition.NewSettingDefinitionClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building SettingDefinition client: %+v", err)
	}
	configureFunc(settingDefinitionClient.Client)

	softwareUpdateStatusSummaryClient, err := softwareupdatestatussummary.NewSoftwareUpdateStatusSummaryClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building SoftwareUpdateStatusSummary client: %+v", err)
	}
	configureFunc(softwareUpdateStatusSummaryClient.Client)

	telecomExpenseManagementPartnerClient, err := telecomexpensemanagementpartner.NewTelecomExpenseManagementPartnerClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building TelecomExpenseManagementPartner client: %+v", err)
	}
	configureFunc(telecomExpenseManagementPartnerClient.Client)

	templateCategoryClient, err := templatecategory.NewTemplateCategoryClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building TemplateCategory client: %+v", err)
	}
	configureFunc(templateCategoryClient.Client)

	templateCategoryRecommendedSettingClient, err := templatecategoryrecommendedsetting.NewTemplateCategoryRecommendedSettingClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building TemplateCategoryRecommendedSetting client: %+v", err)
	}
	configureFunc(templateCategoryRecommendedSettingClient.Client)

	templateCategorySettingDefinitionClient, err := templatecategorysettingdefinition.NewTemplateCategorySettingDefinitionClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building TemplateCategorySettingDefinition client: %+v", err)
	}
	configureFunc(templateCategorySettingDefinitionClient.Client)

	templateClient, err := template.NewTemplateClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Template client: %+v", err)
	}
	configureFunc(templateClient.Client)

	templateInsightClient, err := templateinsight.NewTemplateInsightClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building TemplateInsight client: %+v", err)
	}
	configureFunc(templateInsightClient.Client)

	templateMigratableToCategoryClient, err := templatemigratabletocategory.NewTemplateMigratableToCategoryClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building TemplateMigratableToCategory client: %+v", err)
	}
	configureFunc(templateMigratableToCategoryClient.Client)

	templateMigratableToCategoryRecommendedSettingClient, err := templatemigratabletocategoryrecommendedsetting.NewTemplateMigratableToCategoryRecommendedSettingClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building TemplateMigratableToCategoryRecommendedSetting client: %+v", err)
	}
	configureFunc(templateMigratableToCategoryRecommendedSettingClient.Client)

	templateMigratableToCategorySettingDefinitionClient, err := templatemigratabletocategorysettingdefinition.NewTemplateMigratableToCategorySettingDefinitionClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building TemplateMigratableToCategorySettingDefinition client: %+v", err)
	}
	configureFunc(templateMigratableToCategorySettingDefinitionClient.Client)

	templateMigratableToClient, err := templatemigratableto.NewTemplateMigratableToClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building TemplateMigratableTo client: %+v", err)
	}
	configureFunc(templateMigratableToClient.Client)

	templateMigratableToSettingClient, err := templatemigratabletosetting.NewTemplateMigratableToSettingClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building TemplateMigratableToSetting client: %+v", err)
	}
	configureFunc(templateMigratableToSettingClient.Client)

	templateSettingClient, err := templatesetting.NewTemplateSettingClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building TemplateSetting client: %+v", err)
	}
	configureFunc(templateSettingClient.Client)

	templateSettingSettingDefinitionClient, err := templatesettingsettingdefinition.NewTemplateSettingSettingDefinitionClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building TemplateSettingSettingDefinition client: %+v", err)
	}
	configureFunc(templateSettingSettingDefinitionClient.Client)

	tenantAttachRBACClient, err := tenantattachrbac.NewTenantAttachRBACClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building TenantAttachRBAC client: %+v", err)
	}
	configureFunc(tenantAttachRBACClient.Client)

	termsAndConditionAcceptanceStatusClient, err := termsandconditionacceptancestatus.NewTermsAndConditionAcceptanceStatusClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building TermsAndConditionAcceptanceStatus client: %+v", err)
	}
	configureFunc(termsAndConditionAcceptanceStatusClient.Client)

	termsAndConditionAcceptanceStatusTermsAndConditionClient, err := termsandconditionacceptancestatustermsandcondition.NewTermsAndConditionAcceptanceStatusTermsAndConditionClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building TermsAndConditionAcceptanceStatusTermsAndCondition client: %+v", err)
	}
	configureFunc(termsAndConditionAcceptanceStatusTermsAndConditionClient.Client)

	termsAndConditionAssignmentClient, err := termsandconditionassignment.NewTermsAndConditionAssignmentClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building TermsAndConditionAssignment client: %+v", err)
	}
	configureFunc(termsAndConditionAssignmentClient.Client)

	termsAndConditionClient, err := termsandcondition.NewTermsAndConditionClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building TermsAndCondition client: %+v", err)
	}
	configureFunc(termsAndConditionClient.Client)

	termsAndConditionGroupAssignmentClient, err := termsandconditiongroupassignment.NewTermsAndConditionGroupAssignmentClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building TermsAndConditionGroupAssignment client: %+v", err)
	}
	configureFunc(termsAndConditionGroupAssignmentClient.Client)

	termsAndConditionGroupAssignmentTermsAndConditionClient, err := termsandconditiongroupassignmenttermsandcondition.NewTermsAndConditionGroupAssignmentTermsAndConditionClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building TermsAndConditionGroupAssignmentTermsAndCondition client: %+v", err)
	}
	configureFunc(termsAndConditionGroupAssignmentTermsAndConditionClient.Client)

	troubleshootingEventClient, err := troubleshootingevent.NewTroubleshootingEventClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building TroubleshootingEvent client: %+v", err)
	}
	configureFunc(troubleshootingEventClient.Client)

	updateDeviceManagementMacOSSoftwareUpdateAccountSummaryCategorySummaryStateSummariesClient, err := updatedevicemanagementmacossoftwareupdateaccountsummarycategorysummarystatesummaries.NewUpdateDeviceManagementMacOSSoftwareUpdateAccountSummaryCategorySummaryStateSummariesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UpdateDeviceManagementMacOSSoftwareUpdateAccountSummaryCategorySummaryStateSummaries client: %+v", err)
	}
	configureFunc(updateDeviceManagementMacOSSoftwareUpdateAccountSummaryCategorySummaryStateSummariesClient.Client)

	userExperienceAnalyticsAnomalyClient, err := userexperienceanalyticsanomaly.NewUserExperienceAnalyticsAnomalyClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserExperienceAnalyticsAnomaly client: %+v", err)
	}
	configureFunc(userExperienceAnalyticsAnomalyClient.Client)

	userExperienceAnalyticsAnomalyCorrelationGroupOverviewClient, err := userexperienceanalyticsanomalycorrelationgroupoverview.NewUserExperienceAnalyticsAnomalyCorrelationGroupOverviewClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserExperienceAnalyticsAnomalyCorrelationGroupOverview client: %+v", err)
	}
	configureFunc(userExperienceAnalyticsAnomalyCorrelationGroupOverviewClient.Client)

	userExperienceAnalyticsAnomalyDeviceClient, err := userexperienceanalyticsanomalydevice.NewUserExperienceAnalyticsAnomalyDeviceClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserExperienceAnalyticsAnomalyDevice client: %+v", err)
	}
	configureFunc(userExperienceAnalyticsAnomalyDeviceClient.Client)

	userExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionClient, err := userexperienceanalyticsapphealthapplicationperformancebyappversion.NewUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersion client: %+v", err)
	}
	configureFunc(userExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionClient.Client)

	userExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDetailClient, err := userexperienceanalyticsapphealthapplicationperformancebyappversiondetail.NewUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDetailClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDetail client: %+v", err)
	}
	configureFunc(userExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDetailClient.Client)

	userExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDeviceIdClient, err := userexperienceanalyticsapphealthapplicationperformancebyappversiondeviceid.NewUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDeviceIdClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDeviceId client: %+v", err)
	}
	configureFunc(userExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDeviceIdClient.Client)

	userExperienceAnalyticsAppHealthApplicationPerformanceByOSVersionClient, err := userexperienceanalyticsapphealthapplicationperformancebyosversion.NewUserExperienceAnalyticsAppHealthApplicationPerformanceByOSVersionClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserExperienceAnalyticsAppHealthApplicationPerformanceByOSVersion client: %+v", err)
	}
	configureFunc(userExperienceAnalyticsAppHealthApplicationPerformanceByOSVersionClient.Client)

	userExperienceAnalyticsAppHealthApplicationPerformanceClient, err := userexperienceanalyticsapphealthapplicationperformance.NewUserExperienceAnalyticsAppHealthApplicationPerformanceClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserExperienceAnalyticsAppHealthApplicationPerformance client: %+v", err)
	}
	configureFunc(userExperienceAnalyticsAppHealthApplicationPerformanceClient.Client)

	userExperienceAnalyticsAppHealthDeviceModelPerformanceClient, err := userexperienceanalyticsapphealthdevicemodelperformance.NewUserExperienceAnalyticsAppHealthDeviceModelPerformanceClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserExperienceAnalyticsAppHealthDeviceModelPerformance client: %+v", err)
	}
	configureFunc(userExperienceAnalyticsAppHealthDeviceModelPerformanceClient.Client)

	userExperienceAnalyticsAppHealthDevicePerformanceClient, err := userexperienceanalyticsapphealthdeviceperformance.NewUserExperienceAnalyticsAppHealthDevicePerformanceClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserExperienceAnalyticsAppHealthDevicePerformance client: %+v", err)
	}
	configureFunc(userExperienceAnalyticsAppHealthDevicePerformanceClient.Client)

	userExperienceAnalyticsAppHealthDevicePerformanceDetailClient, err := userexperienceanalyticsapphealthdeviceperformancedetail.NewUserExperienceAnalyticsAppHealthDevicePerformanceDetailClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserExperienceAnalyticsAppHealthDevicePerformanceDetail client: %+v", err)
	}
	configureFunc(userExperienceAnalyticsAppHealthDevicePerformanceDetailClient.Client)

	userExperienceAnalyticsAppHealthOSVersionPerformanceClient, err := userexperienceanalyticsapphealthosversionperformance.NewUserExperienceAnalyticsAppHealthOSVersionPerformanceClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserExperienceAnalyticsAppHealthOSVersionPerformance client: %+v", err)
	}
	configureFunc(userExperienceAnalyticsAppHealthOSVersionPerformanceClient.Client)

	userExperienceAnalyticsAppHealthOverviewClient, err := userexperienceanalyticsapphealthoverview.NewUserExperienceAnalyticsAppHealthOverviewClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserExperienceAnalyticsAppHealthOverview client: %+v", err)
	}
	configureFunc(userExperienceAnalyticsAppHealthOverviewClient.Client)

	userExperienceAnalyticsAppHealthOverviewMetricValueClient, err := userexperienceanalyticsapphealthoverviewmetricvalue.NewUserExperienceAnalyticsAppHealthOverviewMetricValueClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserExperienceAnalyticsAppHealthOverviewMetricValue client: %+v", err)
	}
	configureFunc(userExperienceAnalyticsAppHealthOverviewMetricValueClient.Client)

	userExperienceAnalyticsBaselineAppHealthMetricClient, err := userexperienceanalyticsbaselineapphealthmetric.NewUserExperienceAnalyticsBaselineAppHealthMetricClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserExperienceAnalyticsBaselineAppHealthMetric client: %+v", err)
	}
	configureFunc(userExperienceAnalyticsBaselineAppHealthMetricClient.Client)

	userExperienceAnalyticsBaselineBatteryHealthMetricClient, err := userexperienceanalyticsbaselinebatteryhealthmetric.NewUserExperienceAnalyticsBaselineBatteryHealthMetricClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserExperienceAnalyticsBaselineBatteryHealthMetric client: %+v", err)
	}
	configureFunc(userExperienceAnalyticsBaselineBatteryHealthMetricClient.Client)

	userExperienceAnalyticsBaselineBestPracticesMetricClient, err := userexperienceanalyticsbaselinebestpracticesmetric.NewUserExperienceAnalyticsBaselineBestPracticesMetricClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserExperienceAnalyticsBaselineBestPracticesMetric client: %+v", err)
	}
	configureFunc(userExperienceAnalyticsBaselineBestPracticesMetricClient.Client)

	userExperienceAnalyticsBaselineClient, err := userexperienceanalyticsbaseline.NewUserExperienceAnalyticsBaselineClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserExperienceAnalyticsBaseline client: %+v", err)
	}
	configureFunc(userExperienceAnalyticsBaselineClient.Client)

	userExperienceAnalyticsBaselineDeviceBootPerformanceMetricClient, err := userexperienceanalyticsbaselinedevicebootperformancemetric.NewUserExperienceAnalyticsBaselineDeviceBootPerformanceMetricClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserExperienceAnalyticsBaselineDeviceBootPerformanceMetric client: %+v", err)
	}
	configureFunc(userExperienceAnalyticsBaselineDeviceBootPerformanceMetricClient.Client)

	userExperienceAnalyticsBaselineResourcePerformanceMetricClient, err := userexperienceanalyticsbaselineresourceperformancemetric.NewUserExperienceAnalyticsBaselineResourcePerformanceMetricClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserExperienceAnalyticsBaselineResourcePerformanceMetric client: %+v", err)
	}
	configureFunc(userExperienceAnalyticsBaselineResourcePerformanceMetricClient.Client)

	userExperienceAnalyticsBaselineWorkFromAnywhereMetricClient, err := userexperienceanalyticsbaselineworkfromanywheremetric.NewUserExperienceAnalyticsBaselineWorkFromAnywhereMetricClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserExperienceAnalyticsBaselineWorkFromAnywhereMetric client: %+v", err)
	}
	configureFunc(userExperienceAnalyticsBaselineWorkFromAnywhereMetricClient.Client)

	userExperienceAnalyticsBatteryHealthAppImpactClient, err := userexperienceanalyticsbatteryhealthappimpact.NewUserExperienceAnalyticsBatteryHealthAppImpactClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserExperienceAnalyticsBatteryHealthAppImpact client: %+v", err)
	}
	configureFunc(userExperienceAnalyticsBatteryHealthAppImpactClient.Client)

	userExperienceAnalyticsBatteryHealthCapacityDetailClient, err := userexperienceanalyticsbatteryhealthcapacitydetail.NewUserExperienceAnalyticsBatteryHealthCapacityDetailClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserExperienceAnalyticsBatteryHealthCapacityDetail client: %+v", err)
	}
	configureFunc(userExperienceAnalyticsBatteryHealthCapacityDetailClient.Client)

	userExperienceAnalyticsBatteryHealthDeviceAppImpactClient, err := userexperienceanalyticsbatteryhealthdeviceappimpact.NewUserExperienceAnalyticsBatteryHealthDeviceAppImpactClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserExperienceAnalyticsBatteryHealthDeviceAppImpact client: %+v", err)
	}
	configureFunc(userExperienceAnalyticsBatteryHealthDeviceAppImpactClient.Client)

	userExperienceAnalyticsBatteryHealthDevicePerformanceClient, err := userexperienceanalyticsbatteryhealthdeviceperformance.NewUserExperienceAnalyticsBatteryHealthDevicePerformanceClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserExperienceAnalyticsBatteryHealthDevicePerformance client: %+v", err)
	}
	configureFunc(userExperienceAnalyticsBatteryHealthDevicePerformanceClient.Client)

	userExperienceAnalyticsBatteryHealthDeviceRuntimeHistoryClient, err := userexperienceanalyticsbatteryhealthdeviceruntimehistory.NewUserExperienceAnalyticsBatteryHealthDeviceRuntimeHistoryClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserExperienceAnalyticsBatteryHealthDeviceRuntimeHistory client: %+v", err)
	}
	configureFunc(userExperienceAnalyticsBatteryHealthDeviceRuntimeHistoryClient.Client)

	userExperienceAnalyticsBatteryHealthModelPerformanceClient, err := userexperienceanalyticsbatteryhealthmodelperformance.NewUserExperienceAnalyticsBatteryHealthModelPerformanceClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserExperienceAnalyticsBatteryHealthModelPerformance client: %+v", err)
	}
	configureFunc(userExperienceAnalyticsBatteryHealthModelPerformanceClient.Client)

	userExperienceAnalyticsBatteryHealthOsPerformanceClient, err := userexperienceanalyticsbatteryhealthosperformance.NewUserExperienceAnalyticsBatteryHealthOsPerformanceClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserExperienceAnalyticsBatteryHealthOsPerformance client: %+v", err)
	}
	configureFunc(userExperienceAnalyticsBatteryHealthOsPerformanceClient.Client)

	userExperienceAnalyticsBatteryHealthRuntimeDetailClient, err := userexperienceanalyticsbatteryhealthruntimedetail.NewUserExperienceAnalyticsBatteryHealthRuntimeDetailClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserExperienceAnalyticsBatteryHealthRuntimeDetail client: %+v", err)
	}
	configureFunc(userExperienceAnalyticsBatteryHealthRuntimeDetailClient.Client)

	userExperienceAnalyticsCategoryClient, err := userexperienceanalyticscategory.NewUserExperienceAnalyticsCategoryClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserExperienceAnalyticsCategory client: %+v", err)
	}
	configureFunc(userExperienceAnalyticsCategoryClient.Client)

	userExperienceAnalyticsCategoryMetricValueClient, err := userexperienceanalyticscategorymetricvalue.NewUserExperienceAnalyticsCategoryMetricValueClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserExperienceAnalyticsCategoryMetricValue client: %+v", err)
	}
	configureFunc(userExperienceAnalyticsCategoryMetricValueClient.Client)

	userExperienceAnalyticsDeviceMetricHistoryClient, err := userexperienceanalyticsdevicemetrichistory.NewUserExperienceAnalyticsDeviceMetricHistoryClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserExperienceAnalyticsDeviceMetricHistory client: %+v", err)
	}
	configureFunc(userExperienceAnalyticsDeviceMetricHistoryClient.Client)

	userExperienceAnalyticsDevicePerformanceClient, err := userexperienceanalyticsdeviceperformance.NewUserExperienceAnalyticsDevicePerformanceClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserExperienceAnalyticsDevicePerformance client: %+v", err)
	}
	configureFunc(userExperienceAnalyticsDevicePerformanceClient.Client)

	userExperienceAnalyticsDeviceScopeClient, err := userexperienceanalyticsdevicescope.NewUserExperienceAnalyticsDeviceScopeClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserExperienceAnalyticsDeviceScope client: %+v", err)
	}
	configureFunc(userExperienceAnalyticsDeviceScopeClient.Client)

	userExperienceAnalyticsDeviceScoreClient, err := userexperienceanalyticsdevicescore.NewUserExperienceAnalyticsDeviceScoreClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserExperienceAnalyticsDeviceScore client: %+v", err)
	}
	configureFunc(userExperienceAnalyticsDeviceScoreClient.Client)

	userExperienceAnalyticsDeviceStartupHistoryClient, err := userexperienceanalyticsdevicestartuphistory.NewUserExperienceAnalyticsDeviceStartupHistoryClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserExperienceAnalyticsDeviceStartupHistory client: %+v", err)
	}
	configureFunc(userExperienceAnalyticsDeviceStartupHistoryClient.Client)

	userExperienceAnalyticsDeviceStartupProcessClient, err := userexperienceanalyticsdevicestartupprocess.NewUserExperienceAnalyticsDeviceStartupProcessClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserExperienceAnalyticsDeviceStartupProcess client: %+v", err)
	}
	configureFunc(userExperienceAnalyticsDeviceStartupProcessClient.Client)

	userExperienceAnalyticsDeviceStartupProcessPerformanceClient, err := userexperienceanalyticsdevicestartupprocessperformance.NewUserExperienceAnalyticsDeviceStartupProcessPerformanceClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserExperienceAnalyticsDeviceStartupProcessPerformance client: %+v", err)
	}
	configureFunc(userExperienceAnalyticsDeviceStartupProcessPerformanceClient.Client)

	userExperienceAnalyticsDeviceTimelineEventClient, err := userexperienceanalyticsdevicetimelineevent.NewUserExperienceAnalyticsDeviceTimelineEventClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserExperienceAnalyticsDeviceTimelineEvent client: %+v", err)
	}
	configureFunc(userExperienceAnalyticsDeviceTimelineEventClient.Client)

	userExperienceAnalyticsDevicesWithoutCloudIdentityClient, err := userexperienceanalyticsdeviceswithoutcloudidentity.NewUserExperienceAnalyticsDevicesWithoutCloudIdentityClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserExperienceAnalyticsDevicesWithoutCloudIdentity client: %+v", err)
	}
	configureFunc(userExperienceAnalyticsDevicesWithoutCloudIdentityClient.Client)

	userExperienceAnalyticsImpactingProcessClient, err := userexperienceanalyticsimpactingprocess.NewUserExperienceAnalyticsImpactingProcessClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserExperienceAnalyticsImpactingProcess client: %+v", err)
	}
	configureFunc(userExperienceAnalyticsImpactingProcessClient.Client)

	userExperienceAnalyticsMetricHistoryClient, err := userexperienceanalyticsmetrichistory.NewUserExperienceAnalyticsMetricHistoryClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserExperienceAnalyticsMetricHistory client: %+v", err)
	}
	configureFunc(userExperienceAnalyticsMetricHistoryClient.Client)

	userExperienceAnalyticsModelScoreClient, err := userexperienceanalyticsmodelscore.NewUserExperienceAnalyticsModelScoreClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserExperienceAnalyticsModelScore client: %+v", err)
	}
	configureFunc(userExperienceAnalyticsModelScoreClient.Client)

	userExperienceAnalyticsNotAutopilotReadyDeviceClient, err := userexperienceanalyticsnotautopilotreadydevice.NewUserExperienceAnalyticsNotAutopilotReadyDeviceClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserExperienceAnalyticsNotAutopilotReadyDevice client: %+v", err)
	}
	configureFunc(userExperienceAnalyticsNotAutopilotReadyDeviceClient.Client)

	userExperienceAnalyticsOverviewClient, err := userexperienceanalyticsoverview.NewUserExperienceAnalyticsOverviewClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserExperienceAnalyticsOverview client: %+v", err)
	}
	configureFunc(userExperienceAnalyticsOverviewClient.Client)

	userExperienceAnalyticsRemoteConnectionClient, err := userexperienceanalyticsremoteconnection.NewUserExperienceAnalyticsRemoteConnectionClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserExperienceAnalyticsRemoteConnection client: %+v", err)
	}
	configureFunc(userExperienceAnalyticsRemoteConnectionClient.Client)

	userExperienceAnalyticsResourcePerformanceClient, err := userexperienceanalyticsresourceperformance.NewUserExperienceAnalyticsResourcePerformanceClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserExperienceAnalyticsResourcePerformance client: %+v", err)
	}
	configureFunc(userExperienceAnalyticsResourcePerformanceClient.Client)

	userExperienceAnalyticsScoreHistoryClient, err := userexperienceanalyticsscorehistory.NewUserExperienceAnalyticsScoreHistoryClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserExperienceAnalyticsScoreHistory client: %+v", err)
	}
	configureFunc(userExperienceAnalyticsScoreHistoryClient.Client)

	userExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetricClient, err := userexperienceanalyticsworkfromanywherehardwarereadinessmetric.NewUserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetricClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric client: %+v", err)
	}
	configureFunc(userExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetricClient.Client)

	userExperienceAnalyticsWorkFromAnywhereMetricClient, err := userexperienceanalyticsworkfromanywheremetric.NewUserExperienceAnalyticsWorkFromAnywhereMetricClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserExperienceAnalyticsWorkFromAnywhereMetric client: %+v", err)
	}
	configureFunc(userExperienceAnalyticsWorkFromAnywhereMetricClient.Client)

	userExperienceAnalyticsWorkFromAnywhereMetricMetricDeviceClient, err := userexperienceanalyticsworkfromanywheremetricmetricdevice.NewUserExperienceAnalyticsWorkFromAnywhereMetricMetricDeviceClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserExperienceAnalyticsWorkFromAnywhereMetricMetricDevice client: %+v", err)
	}
	configureFunc(userExperienceAnalyticsWorkFromAnywhereMetricMetricDeviceClient.Client)

	userExperienceAnalyticsWorkFromAnywhereModelPerformanceClient, err := userexperienceanalyticsworkfromanywheremodelperformance.NewUserExperienceAnalyticsWorkFromAnywhereModelPerformanceClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserExperienceAnalyticsWorkFromAnywhereModelPerformance client: %+v", err)
	}
	configureFunc(userExperienceAnalyticsWorkFromAnywhereModelPerformanceClient.Client)

	userPfxCertificateClient, err := userpfxcertificate.NewUserPfxCertificateClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserPfxCertificate client: %+v", err)
	}
	configureFunc(userPfxCertificateClient.Client)

	virtualEndpointAuditEventClient, err := virtualendpointauditevent.NewVirtualEndpointAuditEventClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building VirtualEndpointAuditEvent client: %+v", err)
	}
	configureFunc(virtualEndpointAuditEventClient.Client)

	virtualEndpointBulkActionClient, err := virtualendpointbulkaction.NewVirtualEndpointBulkActionClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building VirtualEndpointBulkAction client: %+v", err)
	}
	configureFunc(virtualEndpointBulkActionClient.Client)

	virtualEndpointClient, err := virtualendpoint.NewVirtualEndpointClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building VirtualEndpoint client: %+v", err)
	}
	configureFunc(virtualEndpointClient.Client)

	virtualEndpointCloudPCClient, err := virtualendpointcloudpc.NewVirtualEndpointCloudPCClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building VirtualEndpointCloudPC client: %+v", err)
	}
	configureFunc(virtualEndpointCloudPCClient.Client)

	virtualEndpointCrossCloudGovernmentOrganizationMappingClient, err := virtualendpointcrosscloudgovernmentorganizationmapping.NewVirtualEndpointCrossCloudGovernmentOrganizationMappingClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building VirtualEndpointCrossCloudGovernmentOrganizationMapping client: %+v", err)
	}
	configureFunc(virtualEndpointCrossCloudGovernmentOrganizationMappingClient.Client)

	virtualEndpointDeviceImageClient, err := virtualendpointdeviceimage.NewVirtualEndpointDeviceImageClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building VirtualEndpointDeviceImage client: %+v", err)
	}
	configureFunc(virtualEndpointDeviceImageClient.Client)

	virtualEndpointExternalPartnerSettingClient, err := virtualendpointexternalpartnersetting.NewVirtualEndpointExternalPartnerSettingClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building VirtualEndpointExternalPartnerSetting client: %+v", err)
	}
	configureFunc(virtualEndpointExternalPartnerSettingClient.Client)

	virtualEndpointFrontLineServicePlanClient, err := virtualendpointfrontlineserviceplan.NewVirtualEndpointFrontLineServicePlanClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building VirtualEndpointFrontLineServicePlan client: %+v", err)
	}
	configureFunc(virtualEndpointFrontLineServicePlanClient.Client)

	virtualEndpointGalleryImageClient, err := virtualendpointgalleryimage.NewVirtualEndpointGalleryImageClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building VirtualEndpointGalleryImage client: %+v", err)
	}
	configureFunc(virtualEndpointGalleryImageClient.Client)

	virtualEndpointOnPremisesConnectionClient, err := virtualendpointonpremisesconnection.NewVirtualEndpointOnPremisesConnectionClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building VirtualEndpointOnPremisesConnection client: %+v", err)
	}
	configureFunc(virtualEndpointOnPremisesConnectionClient.Client)

	virtualEndpointOrganizationSettingClient, err := virtualendpointorganizationsetting.NewVirtualEndpointOrganizationSettingClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building VirtualEndpointOrganizationSetting client: %+v", err)
	}
	configureFunc(virtualEndpointOrganizationSettingClient.Client)

	virtualEndpointProvisioningPolicyAssignmentAssignedUserClient, err := virtualendpointprovisioningpolicyassignmentassigneduser.NewVirtualEndpointProvisioningPolicyAssignmentAssignedUserClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building VirtualEndpointProvisioningPolicyAssignmentAssignedUser client: %+v", err)
	}
	configureFunc(virtualEndpointProvisioningPolicyAssignmentAssignedUserClient.Client)

	virtualEndpointProvisioningPolicyAssignmentAssignedUserMailboxSettingClient, err := virtualendpointprovisioningpolicyassignmentassignedusermailboxsetting.NewVirtualEndpointProvisioningPolicyAssignmentAssignedUserMailboxSettingClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building VirtualEndpointProvisioningPolicyAssignmentAssignedUserMailboxSetting client: %+v", err)
	}
	configureFunc(virtualEndpointProvisioningPolicyAssignmentAssignedUserMailboxSettingClient.Client)

	virtualEndpointProvisioningPolicyAssignmentAssignedUserServiceProvisioningErrorClient, err := virtualendpointprovisioningpolicyassignmentassigneduserserviceprovisioningerror.NewVirtualEndpointProvisioningPolicyAssignmentAssignedUserServiceProvisioningErrorClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building VirtualEndpointProvisioningPolicyAssignmentAssignedUserServiceProvisioningError client: %+v", err)
	}
	configureFunc(virtualEndpointProvisioningPolicyAssignmentAssignedUserServiceProvisioningErrorClient.Client)

	virtualEndpointProvisioningPolicyAssignmentClient, err := virtualendpointprovisioningpolicyassignment.NewVirtualEndpointProvisioningPolicyAssignmentClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building VirtualEndpointProvisioningPolicyAssignment client: %+v", err)
	}
	configureFunc(virtualEndpointProvisioningPolicyAssignmentClient.Client)

	virtualEndpointProvisioningPolicyClient, err := virtualendpointprovisioningpolicy.NewVirtualEndpointProvisioningPolicyClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building VirtualEndpointProvisioningPolicy client: %+v", err)
	}
	configureFunc(virtualEndpointProvisioningPolicyClient.Client)

	virtualEndpointReportClient, err := virtualendpointreport.NewVirtualEndpointReportClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building VirtualEndpointReport client: %+v", err)
	}
	configureFunc(virtualEndpointReportClient.Client)

	virtualEndpointServicePlanClient, err := virtualendpointserviceplan.NewVirtualEndpointServicePlanClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building VirtualEndpointServicePlan client: %+v", err)
	}
	configureFunc(virtualEndpointServicePlanClient.Client)

	virtualEndpointSnapshotClient, err := virtualendpointsnapshot.NewVirtualEndpointSnapshotClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building VirtualEndpointSnapshot client: %+v", err)
	}
	configureFunc(virtualEndpointSnapshotClient.Client)

	virtualEndpointSupportedRegionClient, err := virtualendpointsupportedregion.NewVirtualEndpointSupportedRegionClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building VirtualEndpointSupportedRegion client: %+v", err)
	}
	configureFunc(virtualEndpointSupportedRegionClient.Client)

	virtualEndpointUserSettingAssignmentClient, err := virtualendpointusersettingassignment.NewVirtualEndpointUserSettingAssignmentClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building VirtualEndpointUserSettingAssignment client: %+v", err)
	}
	configureFunc(virtualEndpointUserSettingAssignmentClient.Client)

	virtualEndpointUserSettingClient, err := virtualendpointusersetting.NewVirtualEndpointUserSettingClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building VirtualEndpointUserSetting client: %+v", err)
	}
	configureFunc(virtualEndpointUserSettingClient.Client)

	windowsAutopilotDeploymentProfileAssignedDeviceClient, err := windowsautopilotdeploymentprofileassigneddevice.NewWindowsAutopilotDeploymentProfileAssignedDeviceClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building WindowsAutopilotDeploymentProfileAssignedDevice client: %+v", err)
	}
	configureFunc(windowsAutopilotDeploymentProfileAssignedDeviceClient.Client)

	windowsAutopilotDeploymentProfileAssignedDeviceDeploymentProfileClient, err := windowsautopilotdeploymentprofileassigneddevicedeploymentprofile.NewWindowsAutopilotDeploymentProfileAssignedDeviceDeploymentProfileClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building WindowsAutopilotDeploymentProfileAssignedDeviceDeploymentProfile client: %+v", err)
	}
	configureFunc(windowsAutopilotDeploymentProfileAssignedDeviceDeploymentProfileClient.Client)

	windowsAutopilotDeploymentProfileAssignedDeviceIntendedDeploymentProfileClient, err := windowsautopilotdeploymentprofileassigneddeviceintendeddeploymentprofile.NewWindowsAutopilotDeploymentProfileAssignedDeviceIntendedDeploymentProfileClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building WindowsAutopilotDeploymentProfileAssignedDeviceIntendedDeploymentProfile client: %+v", err)
	}
	configureFunc(windowsAutopilotDeploymentProfileAssignedDeviceIntendedDeploymentProfileClient.Client)

	windowsAutopilotDeploymentProfileAssignmentClient, err := windowsautopilotdeploymentprofileassignment.NewWindowsAutopilotDeploymentProfileAssignmentClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building WindowsAutopilotDeploymentProfileAssignment client: %+v", err)
	}
	configureFunc(windowsAutopilotDeploymentProfileAssignmentClient.Client)

	windowsAutopilotDeploymentProfileClient, err := windowsautopilotdeploymentprofile.NewWindowsAutopilotDeploymentProfileClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building WindowsAutopilotDeploymentProfile client: %+v", err)
	}
	configureFunc(windowsAutopilotDeploymentProfileClient.Client)

	windowsAutopilotDeviceIdentityClient, err := windowsautopilotdeviceidentity.NewWindowsAutopilotDeviceIdentityClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building WindowsAutopilotDeviceIdentity client: %+v", err)
	}
	configureFunc(windowsAutopilotDeviceIdentityClient.Client)

	windowsAutopilotDeviceIdentityDeploymentProfileClient, err := windowsautopilotdeviceidentitydeploymentprofile.NewWindowsAutopilotDeviceIdentityDeploymentProfileClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building WindowsAutopilotDeviceIdentityDeploymentProfile client: %+v", err)
	}
	configureFunc(windowsAutopilotDeviceIdentityDeploymentProfileClient.Client)

	windowsAutopilotDeviceIdentityIntendedDeploymentProfileClient, err := windowsautopilotdeviceidentityintendeddeploymentprofile.NewWindowsAutopilotDeviceIdentityIntendedDeploymentProfileClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building WindowsAutopilotDeviceIdentityIntendedDeploymentProfile client: %+v", err)
	}
	configureFunc(windowsAutopilotDeviceIdentityIntendedDeploymentProfileClient.Client)

	windowsAutopilotSettingClient, err := windowsautopilotsetting.NewWindowsAutopilotSettingClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building WindowsAutopilotSetting client: %+v", err)
	}
	configureFunc(windowsAutopilotSettingClient.Client)

	windowsDriverUpdateProfileAssignmentClient, err := windowsdriverupdateprofileassignment.NewWindowsDriverUpdateProfileAssignmentClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building WindowsDriverUpdateProfileAssignment client: %+v", err)
	}
	configureFunc(windowsDriverUpdateProfileAssignmentClient.Client)

	windowsDriverUpdateProfileClient, err := windowsdriverupdateprofile.NewWindowsDriverUpdateProfileClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building WindowsDriverUpdateProfile client: %+v", err)
	}
	configureFunc(windowsDriverUpdateProfileClient.Client)

	windowsDriverUpdateProfileDriverInventoryClient, err := windowsdriverupdateprofiledriverinventory.NewWindowsDriverUpdateProfileDriverInventoryClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building WindowsDriverUpdateProfileDriverInventory client: %+v", err)
	}
	configureFunc(windowsDriverUpdateProfileDriverInventoryClient.Client)

	windowsFeatureUpdateProfileAssignmentClient, err := windowsfeatureupdateprofileassignment.NewWindowsFeatureUpdateProfileAssignmentClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building WindowsFeatureUpdateProfileAssignment client: %+v", err)
	}
	configureFunc(windowsFeatureUpdateProfileAssignmentClient.Client)

	windowsFeatureUpdateProfileClient, err := windowsfeatureupdateprofile.NewWindowsFeatureUpdateProfileClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building WindowsFeatureUpdateProfile client: %+v", err)
	}
	configureFunc(windowsFeatureUpdateProfileClient.Client)

	windowsInformationProtectionAppLearningSummaryClient, err := windowsinformationprotectionapplearningsummary.NewWindowsInformationProtectionAppLearningSummaryClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building WindowsInformationProtectionAppLearningSummary client: %+v", err)
	}
	configureFunc(windowsInformationProtectionAppLearningSummaryClient.Client)

	windowsInformationProtectionNetworkLearningSummaryClient, err := windowsinformationprotectionnetworklearningsummary.NewWindowsInformationProtectionNetworkLearningSummaryClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building WindowsInformationProtectionNetworkLearningSummary client: %+v", err)
	}
	configureFunc(windowsInformationProtectionNetworkLearningSummaryClient.Client)

	windowsMalwareInformationClient, err := windowsmalwareinformation.NewWindowsMalwareInformationClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building WindowsMalwareInformation client: %+v", err)
	}
	configureFunc(windowsMalwareInformationClient.Client)

	windowsMalwareInformationDeviceMalwareStateClient, err := windowsmalwareinformationdevicemalwarestate.NewWindowsMalwareInformationDeviceMalwareStateClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building WindowsMalwareInformationDeviceMalwareState client: %+v", err)
	}
	configureFunc(windowsMalwareInformationDeviceMalwareStateClient.Client)

	windowsQualityUpdatePolicyAssignmentClient, err := windowsqualityupdatepolicyassignment.NewWindowsQualityUpdatePolicyAssignmentClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building WindowsQualityUpdatePolicyAssignment client: %+v", err)
	}
	configureFunc(windowsQualityUpdatePolicyAssignmentClient.Client)

	windowsQualityUpdatePolicyClient, err := windowsqualityupdatepolicy.NewWindowsQualityUpdatePolicyClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building WindowsQualityUpdatePolicy client: %+v", err)
	}
	configureFunc(windowsQualityUpdatePolicyClient.Client)

	windowsQualityUpdateProfileAssignmentClient, err := windowsqualityupdateprofileassignment.NewWindowsQualityUpdateProfileAssignmentClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building WindowsQualityUpdateProfileAssignment client: %+v", err)
	}
	configureFunc(windowsQualityUpdateProfileAssignmentClient.Client)

	windowsQualityUpdateProfileClient, err := windowsqualityupdateprofile.NewWindowsQualityUpdateProfileClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building WindowsQualityUpdateProfile client: %+v", err)
	}
	configureFunc(windowsQualityUpdateProfileClient.Client)

	windowsUpdateCatalogItemClient, err := windowsupdatecatalogitem.NewWindowsUpdateCatalogItemClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building WindowsUpdateCatalogItem client: %+v", err)
	}
	configureFunc(windowsUpdateCatalogItemClient.Client)

	zebraFotaArtifactClient, err := zebrafotaartifact.NewZebraFotaArtifactClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ZebraFotaArtifact client: %+v", err)
	}
	configureFunc(zebraFotaArtifactClient.Client)

	zebraFotaConnectorClient, err := zebrafotaconnector.NewZebraFotaConnectorClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ZebraFotaConnector client: %+v", err)
	}
	configureFunc(zebraFotaConnectorClient.Client)

	zebraFotaDeploymentClient, err := zebrafotadeployment.NewZebraFotaDeploymentClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ZebraFotaDeployment client: %+v", err)
	}
	configureFunc(zebraFotaDeploymentClient.Client)

	return &Client{
		AdvancedThreatProtectionOnboardingStateSummary:                                                     advancedThreatProtectionOnboardingStateSummaryClient,
		AdvancedThreatProtectionOnboardingStateSummaryAdvancedThreatProtectionOnboardingDeviceSettingState: advancedThreatProtectionOnboardingStateSummaryAdvancedThreatProtectionOnboardingDeviceSettingStateClient,
		AndroidDeviceOwnerEnrollmentProfile:                                                                androidDeviceOwnerEnrollmentProfileClient,
		AndroidForWorkAppConfigurationSchema:                                                               androidForWorkAppConfigurationSchemaClient,
		AndroidForWorkEnrollmentProfile:                                                                    androidForWorkEnrollmentProfileClient,
		AndroidForWorkSetting:                                                                              androidForWorkSettingClient,
		AndroidManagedStoreAccountEnterpriseSetting:                                                        androidManagedStoreAccountEnterpriseSettingClient,
		AndroidManagedStoreAppConfigurationSchema:                                                          androidManagedStoreAppConfigurationSchemaClient,
		ApplePushNotificationCertificate:                                                                   applePushNotificationCertificateClient,
		AppleUserInitiatedEnrollmentProfile:                                                                appleUserInitiatedEnrollmentProfileClient,
		AppleUserInitiatedEnrollmentProfileAssignment:                                                      appleUserInitiatedEnrollmentProfileAssignmentClient,
		AssignmentFilter:                                                                                   assignmentFilterClient,
		AuditEvent:                                                                                         auditEventClient,
		AutopilotEvent:                                                                                     autopilotEventClient,
		AutopilotEventPolicyStatusDetail:                                                                   autopilotEventPolicyStatusDetailClient,
		CartToClassAssociation:                                                                             cartToClassAssociationClient,
		Category:                                                                                           categoryClient,
		CategorySettingDefinition:                                                                          categorySettingDefinitionClient,
		CertificateConnectorDetail:                                                                         certificateConnectorDetailClient,
		ChromeOSOnboardingSetting:                                                                          chromeOSOnboardingSettingClient,
		CloudPCConnectivityIssue:                                                                           cloudPCConnectivityIssueClient,
		ComanagedDevice:                                                                                    comanagedDeviceClient,
		ComanagedDeviceAssignmentFilterEvaluationStatusDetail:                                              comanagedDeviceAssignmentFilterEvaluationStatusDetailClient,
		ComanagedDeviceDetectedApp:                                                                         comanagedDeviceDetectedAppClient,
		ComanagedDeviceDeviceCategory:                                                                      comanagedDeviceDeviceCategoryClient,
		ComanagedDeviceDeviceCompliancePolicyState:                                                         comanagedDeviceDeviceCompliancePolicyStateClient,
		ComanagedDeviceDeviceConfigurationState:                                                            comanagedDeviceDeviceConfigurationStateClient,
		ComanagedDeviceDeviceHealthScriptState:                                                             comanagedDeviceDeviceHealthScriptStateClient,
		ComanagedDeviceDeviceHealthScriptStateIdIdPolicyIdPolicyIdDeviceIdDeviceId: comanagedDeviceDeviceHealthScriptStateIdIdPolicyIdPolicyIdDeviceIdDeviceIdClient,
		ComanagedDeviceLogCollectionRequest:                                        comanagedDeviceLogCollectionRequestClient,
		ComanagedDeviceManagedDeviceMobileAppConfigurationState:                    comanagedDeviceManagedDeviceMobileAppConfigurationStateClient,
		ComanagedDeviceSecurityBaselineState:                                       comanagedDeviceSecurityBaselineStateClient,
		ComanagedDeviceSecurityBaselineStateSettingState:                           comanagedDeviceSecurityBaselineStateSettingStateClient,
		ComanagedDeviceUser:                                                 comanagedDeviceUserClient,
		ComanagedDeviceWindowsProtectionState:                               comanagedDeviceWindowsProtectionStateClient,
		ComanagedDeviceWindowsProtectionStateDetectedMalwareState:           comanagedDeviceWindowsProtectionStateDetectedMalwareStateClient,
		ComanagementEligibleDevice:                                          comanagementEligibleDeviceClient,
		ComplianceCategory:                                                  complianceCategoryClient,
		ComplianceManagementPartner:                                         complianceManagementPartnerClient,
		CompliancePolicy:                                                    compliancePolicyClient,
		CompliancePolicyAssignment:                                          compliancePolicyAssignmentClient,
		CompliancePolicyScheduledActionsForRule:                             compliancePolicyScheduledActionsForRuleClient,
		CompliancePolicyScheduledActionsForRuleScheduledActionConfiguration: compliancePolicyScheduledActionsForRuleScheduledActionConfigurationClient,
		CompliancePolicySetting:                                             compliancePolicySettingClient,
		CompliancePolicySettingSettingDefinition:                            compliancePolicySettingSettingDefinitionClient,
		ComplianceSetting:                                                   complianceSettingClient,
		ConditionalAccessSetting:                                            conditionalAccessSettingClient,
		ConfigManagerCollection:                                             configManagerCollectionClient,
		ConfigurationCategory:                                               configurationCategoryClient,
		ConfigurationPolicy:                                                 configurationPolicyClient,
		ConfigurationPolicyAssignment:                                       configurationPolicyAssignmentClient,
		ConfigurationPolicySetting:                                          configurationPolicySettingClient,
		ConfigurationPolicySettingSettingDefinition:                         configurationPolicySettingSettingDefinitionClient,
		ConfigurationPolicyTemplate:                                         configurationPolicyTemplateClient,
		ConfigurationPolicyTemplateSettingTemplate:                          configurationPolicyTemplateSettingTemplateClient,
		ConfigurationPolicyTemplateSettingTemplateSettingDefinition:         configurationPolicyTemplateSettingTemplateSettingDefinitionClient,
		ConfigurationSetting:                                                configurationSettingClient,
		DataSharingConsent:                                                  dataSharingConsentClient,
		DepOnboardingSetting:                                                depOnboardingSettingClient,
		DepOnboardingSettingDefaultIosEnrollmentProfile:                     depOnboardingSettingDefaultIosEnrollmentProfileClient,
		DepOnboardingSettingDefaultMacOsEnrollmentProfile:                   depOnboardingSettingDefaultMacOsEnrollmentProfileClient,
		DepOnboardingSettingEnrollmentProfile:                               depOnboardingSettingEnrollmentProfileClient,
		DepOnboardingSettingImportedAppleDeviceIdentity:                     depOnboardingSettingImportedAppleDeviceIdentityClient,
		DerivedCredential:                                                   derivedCredentialClient,
		DetectedApp:                                                         detectedAppClient,
		DetectedAppManagedDevice:                                            detectedAppManagedDeviceClient,
		DeviceCategory:                                                      deviceCategoryClient,
		DeviceCompliancePolicy:                                              deviceCompliancePolicyClient,
		DeviceCompliancePolicyAssignment:                                    deviceCompliancePolicyAssignmentClient,
		DeviceCompliancePolicyDeviceSettingStateSummary:                     deviceCompliancePolicyDeviceSettingStateSummaryClient,
		DeviceCompliancePolicyDeviceStateSummary:                            deviceCompliancePolicyDeviceStateSummaryClient,
		DeviceCompliancePolicyDeviceStatus:                                  deviceCompliancePolicyDeviceStatusClient,
		DeviceCompliancePolicyDeviceStatusOverview:                          deviceCompliancePolicyDeviceStatusOverviewClient,
		DeviceCompliancePolicyScheduledActionsForRule:                       deviceCompliancePolicyScheduledActionsForRuleClient,
		DeviceCompliancePolicyScheduledActionsForRuleScheduledActionConfiguration:                 deviceCompliancePolicyScheduledActionsForRuleScheduledActionConfigurationClient,
		DeviceCompliancePolicySettingStateSummary:                                                 deviceCompliancePolicySettingStateSummaryClient,
		DeviceCompliancePolicySettingStateSummaryDeviceComplianceSettingState:                     deviceCompliancePolicySettingStateSummaryDeviceComplianceSettingStateClient,
		DeviceCompliancePolicyUserStatus:                                                          deviceCompliancePolicyUserStatusClient,
		DeviceCompliancePolicyUserStatusOverview:                                                  deviceCompliancePolicyUserStatusOverviewClient,
		DeviceComplianceScript:                                                                    deviceComplianceScriptClient,
		DeviceComplianceScriptAssignment:                                                          deviceComplianceScriptAssignmentClient,
		DeviceComplianceScriptDeviceRunState:                                                      deviceComplianceScriptDeviceRunStateClient,
		DeviceComplianceScriptDeviceRunStateManagedDevice:                                         deviceComplianceScriptDeviceRunStateManagedDeviceClient,
		DeviceComplianceScriptRunSummary:                                                          deviceComplianceScriptRunSummaryClient,
		DeviceConfiguration:                                                                       deviceConfigurationClient,
		DeviceConfigurationAssignment:                                                             deviceConfigurationAssignmentClient,
		DeviceConfigurationConflictSummary:                                                        deviceConfigurationConflictSummaryClient,
		DeviceConfigurationDeviceSettingStateSummary:                                              deviceConfigurationDeviceSettingStateSummaryClient,
		DeviceConfigurationDeviceStateSummary:                                                     deviceConfigurationDeviceStateSummaryClient,
		DeviceConfigurationDeviceStatus:                                                           deviceConfigurationDeviceStatusClient,
		DeviceConfigurationDeviceStatusOverview:                                                   deviceConfigurationDeviceStatusOverviewClient,
		DeviceConfigurationGroupAssignment:                                                        deviceConfigurationGroupAssignmentClient,
		DeviceConfigurationGroupAssignmentDeviceConfiguration:                                     deviceConfigurationGroupAssignmentDeviceConfigurationClient,
		DeviceConfigurationProfile:                                                                deviceConfigurationProfileClient,
		DeviceConfigurationRestrictedAppsViolation:                                                deviceConfigurationRestrictedAppsViolationClient,
		DeviceConfigurationUserStateSummary:                                                       deviceConfigurationUserStateSummaryClient,
		DeviceConfigurationUserStatus:                                                             deviceConfigurationUserStatusClient,
		DeviceConfigurationUserStatusOverview:                                                     deviceConfigurationUserStatusOverviewClient,
		DeviceConfigurationsAllManagedDeviceCertificateState:                                      deviceConfigurationsAllManagedDeviceCertificateStateClient,
		DeviceCustomAttributeShellScript:                                                          deviceCustomAttributeShellScriptClient,
		DeviceCustomAttributeShellScriptAssignment:                                                deviceCustomAttributeShellScriptAssignmentClient,
		DeviceCustomAttributeShellScriptDeviceRunState:                                            deviceCustomAttributeShellScriptDeviceRunStateClient,
		DeviceCustomAttributeShellScriptDeviceRunStateManagedDevice:                               deviceCustomAttributeShellScriptDeviceRunStateManagedDeviceClient,
		DeviceCustomAttributeShellScriptGroupAssignment:                                           deviceCustomAttributeShellScriptGroupAssignmentClient,
		DeviceCustomAttributeShellScriptRunSummary:                                                deviceCustomAttributeShellScriptRunSummaryClient,
		DeviceCustomAttributeShellScriptUserRunState:                                              deviceCustomAttributeShellScriptUserRunStateClient,
		DeviceCustomAttributeShellScriptUserRunStateDeviceRunState:                                deviceCustomAttributeShellScriptUserRunStateDeviceRunStateClient,
		DeviceCustomAttributeShellScriptUserRunStateDeviceRunStateManagedDevice:                   deviceCustomAttributeShellScriptUserRunStateDeviceRunStateManagedDeviceClient,
		DeviceEnrollmentConfiguration:                                                             deviceEnrollmentConfigurationClient,
		DeviceEnrollmentConfigurationAssignment:                                                   deviceEnrollmentConfigurationAssignmentClient,
		DeviceHealthScript:                                                                        deviceHealthScriptClient,
		DeviceHealthScriptAssignment:                                                              deviceHealthScriptAssignmentClient,
		DeviceHealthScriptDeviceRunState:                                                          deviceHealthScriptDeviceRunStateClient,
		DeviceHealthScriptDeviceRunStateManagedDevice:                                             deviceHealthScriptDeviceRunStateManagedDeviceClient,
		DeviceHealthScriptRunSummary:                                                              deviceHealthScriptRunSummaryClient,
		DeviceManagement:                                                                          deviceManagementClient,
		DeviceManagementPartner:                                                                   deviceManagementPartnerClient,
		DeviceManagementScript:                                                                    deviceManagementScriptClient,
		DeviceManagementScriptAssignment:                                                          deviceManagementScriptAssignmentClient,
		DeviceManagementScriptDeviceRunState:                                                      deviceManagementScriptDeviceRunStateClient,
		DeviceManagementScriptDeviceRunStateManagedDevice:                                         deviceManagementScriptDeviceRunStateManagedDeviceClient,
		DeviceManagementScriptGroupAssignment:                                                     deviceManagementScriptGroupAssignmentClient,
		DeviceManagementScriptRunSummary:                                                          deviceManagementScriptRunSummaryClient,
		DeviceManagementScriptUserRunState:                                                        deviceManagementScriptUserRunStateClient,
		DeviceManagementScriptUserRunStateDeviceRunState:                                          deviceManagementScriptUserRunStateDeviceRunStateClient,
		DeviceManagementScriptUserRunStateDeviceRunStateManagedDevice:                             deviceManagementScriptUserRunStateDeviceRunStateManagedDeviceClient,
		DeviceShellScript:                                                                         deviceShellScriptClient,
		DeviceShellScriptAssignment:                                                               deviceShellScriptAssignmentClient,
		DeviceShellScriptDeviceRunState:                                                           deviceShellScriptDeviceRunStateClient,
		DeviceShellScriptDeviceRunStateManagedDevice:                                              deviceShellScriptDeviceRunStateManagedDeviceClient,
		DeviceShellScriptGroupAssignment:                                                          deviceShellScriptGroupAssignmentClient,
		DeviceShellScriptRunSummary:                                                               deviceShellScriptRunSummaryClient,
		DeviceShellScriptUserRunState:                                                             deviceShellScriptUserRunStateClient,
		DeviceShellScriptUserRunStateDeviceRunState:                                               deviceShellScriptUserRunStateDeviceRunStateClient,
		DeviceShellScriptUserRunStateDeviceRunStateManagedDevice:                                  deviceShellScriptUserRunStateDeviceRunStateManagedDeviceClient,
		DomainJoinConnector:                                                                       domainJoinConnectorClient,
		ElevationRequest:                                                                          elevationRequestClient,
		EmbeddedSIMActivationCodePool:                                                             embeddedSIMActivationCodePoolClient,
		EmbeddedSIMActivationCodePoolAssignment:                                                   embeddedSIMActivationCodePoolAssignmentClient,
		EmbeddedSIMActivationCodePoolDeviceState:                                                  embeddedSIMActivationCodePoolDeviceStateClient,
		EndpointPrivilegeManagementProvisioningStatus:                                             endpointPrivilegeManagementProvisioningStatusClient,
		ExchangeConnector:                                                                         exchangeConnectorClient,
		ExchangeOnPremisesPolicy:                                                                  exchangeOnPremisesPolicyClient,
		ExchangeOnPremisesPolicyConditionalAccessSetting:                                          exchangeOnPremisesPolicyConditionalAccessSettingClient,
		ExportDeviceManagementReportsJobs:                                                         exportDeviceManagementReportsJobsClient,
		ExportDeviceManagementVirtualEndpointReportsJobs:                                          exportDeviceManagementVirtualEndpointReportsJobsClient,
		GroupPolicyCategory:                                                                       groupPolicyCategoryClient,
		GroupPolicyCategoryChild:                                                                  groupPolicyCategoryChildClient,
		GroupPolicyCategoryDefinition:                                                             groupPolicyCategoryDefinitionClient,
		GroupPolicyCategoryDefinitionFile:                                                         groupPolicyCategoryDefinitionFileClient,
		GroupPolicyCategoryParent:                                                                 groupPolicyCategoryParentClient,
		GroupPolicyConfiguration:                                                                  groupPolicyConfigurationClient,
		GroupPolicyConfigurationAssignment:                                                        groupPolicyConfigurationAssignmentClient,
		GroupPolicyConfigurationDefinitionValue:                                                   groupPolicyConfigurationDefinitionValueClient,
		GroupPolicyConfigurationDefinitionValueDefinition:                                         groupPolicyConfigurationDefinitionValueDefinitionClient,
		GroupPolicyConfigurationDefinitionValuePresentationValue:                                  groupPolicyConfigurationDefinitionValuePresentationValueClient,
		GroupPolicyConfigurationDefinitionValuePresentationValueDefinitionValue:                   groupPolicyConfigurationDefinitionValuePresentationValueDefinitionValueClient,
		GroupPolicyConfigurationDefinitionValuePresentationValuePresentation:                      groupPolicyConfigurationDefinitionValuePresentationValuePresentationClient,
		GroupPolicyDefinition:                                                                     groupPolicyDefinitionClient,
		GroupPolicyDefinitionCategory:                                                             groupPolicyDefinitionCategoryClient,
		GroupPolicyDefinitionDefinitionFile:                                                       groupPolicyDefinitionDefinitionFileClient,
		GroupPolicyDefinitionFile:                                                                 groupPolicyDefinitionFileClient,
		GroupPolicyDefinitionFileDefinition:                                                       groupPolicyDefinitionFileDefinitionClient,
		GroupPolicyDefinitionNextVersionDefinition:                                                groupPolicyDefinitionNextVersionDefinitionClient,
		GroupPolicyDefinitionNextVersionDefinitionCategory:                                        groupPolicyDefinitionNextVersionDefinitionCategoryClient,
		GroupPolicyDefinitionNextVersionDefinitionDefinitionFile:                                  groupPolicyDefinitionNextVersionDefinitionDefinitionFileClient,
		GroupPolicyDefinitionNextVersionDefinitionPresentation:                                    groupPolicyDefinitionNextVersionDefinitionPresentationClient,
		GroupPolicyDefinitionNextVersionDefinitionPresentationDefinition:                          groupPolicyDefinitionNextVersionDefinitionPresentationDefinitionClient,
		GroupPolicyDefinitionNextVersionDefinitionPreviousVersionDefinition:                       groupPolicyDefinitionNextVersionDefinitionPreviousVersionDefinitionClient,
		GroupPolicyDefinitionNextVersionDefinitionPreviousVersionDefinitionCategory:               groupPolicyDefinitionNextVersionDefinitionPreviousVersionDefinitionCategoryClient,
		GroupPolicyDefinitionNextVersionDefinitionPreviousVersionDefinitionDefinitionFile:         groupPolicyDefinitionNextVersionDefinitionPreviousVersionDefinitionDefinitionFileClient,
		GroupPolicyDefinitionNextVersionDefinitionPreviousVersionDefinitionPresentation:           groupPolicyDefinitionNextVersionDefinitionPreviousVersionDefinitionPresentationClient,
		GroupPolicyDefinitionNextVersionDefinitionPreviousVersionDefinitionPresentationDefinition: groupPolicyDefinitionNextVersionDefinitionPreviousVersionDefinitionPresentationDefinitionClient,
		GroupPolicyDefinitionPresentation:                                                         groupPolicyDefinitionPresentationClient,
		GroupPolicyDefinitionPresentationDefinition:                                               groupPolicyDefinitionPresentationDefinitionClient,
		GroupPolicyDefinitionPreviousVersionDefinition:                                            groupPolicyDefinitionPreviousVersionDefinitionClient,
		GroupPolicyDefinitionPreviousVersionDefinitionCategory:                                    groupPolicyDefinitionPreviousVersionDefinitionCategoryClient,
		GroupPolicyDefinitionPreviousVersionDefinitionDefinitionFile:                              groupPolicyDefinitionPreviousVersionDefinitionDefinitionFileClient,
		GroupPolicyDefinitionPreviousVersionDefinitionNextVersionDefinition:                       groupPolicyDefinitionPreviousVersionDefinitionNextVersionDefinitionClient,
		GroupPolicyDefinitionPreviousVersionDefinitionNextVersionDefinitionCategory:               groupPolicyDefinitionPreviousVersionDefinitionNextVersionDefinitionCategoryClient,
		GroupPolicyDefinitionPreviousVersionDefinitionNextVersionDefinitionDefinitionFile:         groupPolicyDefinitionPreviousVersionDefinitionNextVersionDefinitionDefinitionFileClient,
		GroupPolicyDefinitionPreviousVersionDefinitionNextVersionDefinitionPresentation:           groupPolicyDefinitionPreviousVersionDefinitionNextVersionDefinitionPresentationClient,
		GroupPolicyDefinitionPreviousVersionDefinitionNextVersionDefinitionPresentationDefinition: groupPolicyDefinitionPreviousVersionDefinitionNextVersionDefinitionPresentationDefinitionClient,
		GroupPolicyDefinitionPreviousVersionDefinitionPresentation:                                groupPolicyDefinitionPreviousVersionDefinitionPresentationClient,
		GroupPolicyDefinitionPreviousVersionDefinitionPresentationDefinition:                      groupPolicyDefinitionPreviousVersionDefinitionPresentationDefinitionClient,
		GroupPolicyMigrationReport:                                                                groupPolicyMigrationReportClient,
		GroupPolicyMigrationReportGroupPolicySettingMapping:                                       groupPolicyMigrationReportGroupPolicySettingMappingClient,
		GroupPolicyMigrationReportUnsupportedGroupPolicyExtension:                                 groupPolicyMigrationReportUnsupportedGroupPolicyExtensionClient,
		GroupPolicyObjectFile:                                                                     groupPolicyObjectFileClient,
		GroupPolicyUploadedDefinitionFile:                                                         groupPolicyUploadedDefinitionFileClient,
		GroupPolicyUploadedDefinitionFileDefinition:                                               groupPolicyUploadedDefinitionFileDefinitionClient,
		GroupPolicyUploadedDefinitionFileGroupPolicyOperation:                                     groupPolicyUploadedDefinitionFileGroupPolicyOperationClient,
		HardwareConfiguration:                                                                     hardwareConfigurationClient,
		HardwareConfigurationAssignment:                                                           hardwareConfigurationAssignmentClient,
		HardwareConfigurationDeviceRunState:                                                       hardwareConfigurationDeviceRunStateClient,
		HardwareConfigurationRunSummary:                                                           hardwareConfigurationRunSummaryClient,
		HardwareConfigurationUserRunState:                                                         hardwareConfigurationUserRunStateClient,
		HardwarePasswordDetail:                                                                    hardwarePasswordDetailClient,
		HardwarePasswordInfo:                                                                      hardwarePasswordInfoClient,
		ImportedDeviceIdentity:                                                                    importedDeviceIdentityClient,
		ImportedWindowsAutopilotDeviceIdentity:                                                    importedWindowsAutopilotDeviceIdentityClient,
		Intent:                                                                                    intentClient,
		IntentAssignment:                                                                          intentAssignmentClient,
		IntentCategory:                                                                            intentCategoryClient,
		IntentCategorySetting:                                                                     intentCategorySettingClient,
		IntentCategorySettingDefinition:                                                           intentCategorySettingDefinitionClient,
		IntentDeviceSettingStateSummary:                                                           intentDeviceSettingStateSummaryClient,
		IntentDeviceState:                                                                         intentDeviceStateClient,
		IntentDeviceStateSummary:                                                                  intentDeviceStateSummaryClient,
		IntentSetting:                                                                             intentSettingClient,
		IntentUserState:                                                                           intentUserStateClient,
		IntentUserStateSummary:                                                                    intentUserStateSummaryClient,
		IntuneBrandingProfile:                                                                     intuneBrandingProfileClient,
		IntuneBrandingProfileAssignment:                                                           intuneBrandingProfileAssignmentClient,
		IosUpdateStatus:                                                                           iosUpdateStatusClient,
		MacOSSoftwareUpdateAccountSummary:                                                         macOSSoftwareUpdateAccountSummaryClient,
		MacOSSoftwareUpdateAccountSummaryCategorySummary:                                          macOSSoftwareUpdateAccountSummaryCategorySummaryClient,
		ManagedDevice: managedDeviceClient,
		ManagedDeviceAssignmentFilterEvaluationStatusDetail:                      managedDeviceAssignmentFilterEvaluationStatusDetailClient,
		ManagedDeviceCleanupRule:                                                 managedDeviceCleanupRuleClient,
		ManagedDeviceDetectedApp:                                                 managedDeviceDetectedAppClient,
		ManagedDeviceDeviceCategory:                                              managedDeviceDeviceCategoryClient,
		ManagedDeviceDeviceCompliancePolicyState:                                 managedDeviceDeviceCompliancePolicyStateClient,
		ManagedDeviceDeviceConfigurationState:                                    managedDeviceDeviceConfigurationStateClient,
		ManagedDeviceDeviceHealthScriptState:                                     managedDeviceDeviceHealthScriptStateClient,
		ManagedDeviceDeviceHealthScriptStateIdIdPolicyIdPolicyIdDeviceIdDeviceId: managedDeviceDeviceHealthScriptStateIdIdPolicyIdPolicyIdDeviceIdDeviceIdClient,
		ManagedDeviceEncryptionState:                                             managedDeviceEncryptionStateClient,
		ManagedDeviceLogCollectionRequest:                                        managedDeviceLogCollectionRequestClient,
		ManagedDeviceManagedDeviceMobileAppConfigurationState:                    managedDeviceManagedDeviceMobileAppConfigurationStateClient,
		ManagedDeviceOverview:                                                    managedDeviceOverviewClient,
		ManagedDeviceSecurityBaselineState:                                       managedDeviceSecurityBaselineStateClient,
		ManagedDeviceSecurityBaselineStateSettingState:                           managedDeviceSecurityBaselineStateSettingStateClient,
		ManagedDeviceUser:                                                        managedDeviceUserClient,
		ManagedDeviceWindowsOSImage:                                              managedDeviceWindowsOSImageClient,
		ManagedDeviceWindowsProtectionState:                                      managedDeviceWindowsProtectionStateClient,
		ManagedDeviceWindowsProtectionStateDetectedMalwareState:                  managedDeviceWindowsProtectionStateDetectedMalwareStateClient,
		MicrosoftTunnelConfiguration:                                             microsoftTunnelConfigurationClient,
		MicrosoftTunnelHealthThreshold:                                           microsoftTunnelHealthThresholdClient,
		MicrosoftTunnelServerLogCollectionResponse:                               microsoftTunnelServerLogCollectionResponseClient,
		MicrosoftTunnelSite:                                                      microsoftTunnelSiteClient,
		MicrosoftTunnelSiteMicrosoftTunnelConfiguration:                          microsoftTunnelSiteMicrosoftTunnelConfigurationClient,
		MicrosoftTunnelSiteMicrosoftTunnelServer:                                 microsoftTunnelSiteMicrosoftTunnelServerClient,
		MobileAppTroubleshootingEvent:                                            mobileAppTroubleshootingEventClient,
		MobileAppTroubleshootingEventAppLogCollectionRequest:                     mobileAppTroubleshootingEventAppLogCollectionRequestClient,
		MobileThreatDefenseConnector:                                             mobileThreatDefenseConnectorClient,
		Monitoring:                                                               monitoringClient,
		MonitoringAlertRecord:                                                    monitoringAlertRecordClient,
		MonitoringAlertRule:                                                      monitoringAlertRuleClient,
		NdesConnector:                                                            ndesConnectorClient,
		NotificationMessageTemplate:                                              notificationMessageTemplateClient,
		NotificationMessageTemplateLocalizedNotificationMessage:                  notificationMessageTemplateLocalizedNotificationMessageClient,
		OperationApprovalPolicy:                                                  operationApprovalPolicyClient,
		OperationApprovalRequest:                                                 operationApprovalRequestClient,
		PrivilegeManagementElevation:                                             privilegeManagementElevationClient,
		RebootDeviceManagementUserExperienceAnalyticsBaselineAnalyticsMetrics:    rebootDeviceManagementUserExperienceAnalyticsBaselineAnalyticsMetricsClient,
		RemoteActionAudit:                                                        remoteActionAuditClient,
		RemoteAssistancePartner:                                                  remoteAssistancePartnerClient,
		RemoteAssistanceSetting:                                                  remoteAssistanceSettingClient,
		Report:                                                                   reportClient,
		ReportCachedReportConfiguration:                                          reportCachedReportConfigurationClient,
		ResourceAccessProfile:                                                    resourceAccessProfileClient,
		ResourceAccessProfileAssignment:                                          resourceAccessProfileAssignmentClient,
		ResourceOperation:                                                        resourceOperationClient,
		ReusablePolicySetting:                                                    reusablePolicySettingClient,
		ReusablePolicySettingReferencingConfigurationPolicy:                      reusablePolicySettingReferencingConfigurationPolicyClient,
		ReusablePolicySettingReferencingConfigurationPolicyAssignment:            reusablePolicySettingReferencingConfigurationPolicyAssignmentClient,
		ReusablePolicySettingReferencingConfigurationPolicySetting:               reusablePolicySettingReferencingConfigurationPolicySettingClient,
		ReusablePolicySettingReferencingConfigurationPolicySettingSettingDefinition: reusablePolicySettingReferencingConfigurationPolicySettingSettingDefinitionClient,
		ReusableSetting:                                    reusableSettingClient,
		RoleAssignment:                                     roleAssignmentClient,
		RoleAssignmentRoleDefinition:                       roleAssignmentRoleDefinitionClient,
		RoleAssignmentRoleScopeTag:                         roleAssignmentRoleScopeTagClient,
		RoleDefinition:                                     roleDefinitionClient,
		RoleDefinitionRoleAssignment:                       roleDefinitionRoleAssignmentClient,
		RoleDefinitionRoleAssignmentRoleDefinition:         roleDefinitionRoleAssignmentRoleDefinitionClient,
		RoleScopeTag:                                       roleScopeTagClient,
		RoleScopeTagAssignment:                             roleScopeTagAssignmentClient,
		ServiceNowConnection:                               serviceNowConnectionClient,
		SettingDefinition:                                  settingDefinitionClient,
		SoftwareUpdateStatusSummary:                        softwareUpdateStatusSummaryClient,
		TelecomExpenseManagementPartner:                    telecomExpenseManagementPartnerClient,
		Template:                                           templateClient,
		TemplateCategory:                                   templateCategoryClient,
		TemplateCategoryRecommendedSetting:                 templateCategoryRecommendedSettingClient,
		TemplateCategorySettingDefinition:                  templateCategorySettingDefinitionClient,
		TemplateInsight:                                    templateInsightClient,
		TemplateMigratableTo:                               templateMigratableToClient,
		TemplateMigratableToCategory:                       templateMigratableToCategoryClient,
		TemplateMigratableToCategoryRecommendedSetting:     templateMigratableToCategoryRecommendedSettingClient,
		TemplateMigratableToCategorySettingDefinition:      templateMigratableToCategorySettingDefinitionClient,
		TemplateMigratableToSetting:                        templateMigratableToSettingClient,
		TemplateSetting:                                    templateSettingClient,
		TemplateSettingSettingDefinition:                   templateSettingSettingDefinitionClient,
		TenantAttachRBAC:                                   tenantAttachRBACClient,
		TermsAndCondition:                                  termsAndConditionClient,
		TermsAndConditionAcceptanceStatus:                  termsAndConditionAcceptanceStatusClient,
		TermsAndConditionAcceptanceStatusTermsAndCondition: termsAndConditionAcceptanceStatusTermsAndConditionClient,
		TermsAndConditionAssignment:                        termsAndConditionAssignmentClient,
		TermsAndConditionGroupAssignment:                   termsAndConditionGroupAssignmentClient,
		TermsAndConditionGroupAssignmentTermsAndCondition:  termsAndConditionGroupAssignmentTermsAndConditionClient,
		TroubleshootingEvent:                               troubleshootingEventClient,
		UpdateDeviceManagementMacOSSoftwareUpdateAccountSummaryCategorySummaryStateSummaries: updateDeviceManagementMacOSSoftwareUpdateAccountSummaryCategorySummaryStateSummariesClient,
		UserExperienceAnalyticsAnomaly:                                             userExperienceAnalyticsAnomalyClient,
		UserExperienceAnalyticsAnomalyCorrelationGroupOverview:                     userExperienceAnalyticsAnomalyCorrelationGroupOverviewClient,
		UserExperienceAnalyticsAnomalyDevice:                                       userExperienceAnalyticsAnomalyDeviceClient,
		UserExperienceAnalyticsAppHealthApplicationPerformance:                     userExperienceAnalyticsAppHealthApplicationPerformanceClient,
		UserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersion:         userExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionClient,
		UserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDetail:   userExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDetailClient,
		UserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDeviceId: userExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDeviceIdClient,
		UserExperienceAnalyticsAppHealthApplicationPerformanceByOSVersion:          userExperienceAnalyticsAppHealthApplicationPerformanceByOSVersionClient,
		UserExperienceAnalyticsAppHealthDeviceModelPerformance:                     userExperienceAnalyticsAppHealthDeviceModelPerformanceClient,
		UserExperienceAnalyticsAppHealthDevicePerformance:                          userExperienceAnalyticsAppHealthDevicePerformanceClient,
		UserExperienceAnalyticsAppHealthDevicePerformanceDetail:                    userExperienceAnalyticsAppHealthDevicePerformanceDetailClient,
		UserExperienceAnalyticsAppHealthOSVersionPerformance:                       userExperienceAnalyticsAppHealthOSVersionPerformanceClient,
		UserExperienceAnalyticsAppHealthOverview:                                   userExperienceAnalyticsAppHealthOverviewClient,
		UserExperienceAnalyticsAppHealthOverviewMetricValue:                        userExperienceAnalyticsAppHealthOverviewMetricValueClient,
		UserExperienceAnalyticsBaseline:                                            userExperienceAnalyticsBaselineClient,
		UserExperienceAnalyticsBaselineAppHealthMetric:                             userExperienceAnalyticsBaselineAppHealthMetricClient,
		UserExperienceAnalyticsBaselineBatteryHealthMetric:                         userExperienceAnalyticsBaselineBatteryHealthMetricClient,
		UserExperienceAnalyticsBaselineBestPracticesMetric:                         userExperienceAnalyticsBaselineBestPracticesMetricClient,
		UserExperienceAnalyticsBaselineDeviceBootPerformanceMetric:                 userExperienceAnalyticsBaselineDeviceBootPerformanceMetricClient,
		UserExperienceAnalyticsBaselineResourcePerformanceMetric:                   userExperienceAnalyticsBaselineResourcePerformanceMetricClient,
		UserExperienceAnalyticsBaselineWorkFromAnywhereMetric:                      userExperienceAnalyticsBaselineWorkFromAnywhereMetricClient,
		UserExperienceAnalyticsBatteryHealthAppImpact:                              userExperienceAnalyticsBatteryHealthAppImpactClient,
		UserExperienceAnalyticsBatteryHealthCapacityDetail:                         userExperienceAnalyticsBatteryHealthCapacityDetailClient,
		UserExperienceAnalyticsBatteryHealthDeviceAppImpact:                        userExperienceAnalyticsBatteryHealthDeviceAppImpactClient,
		UserExperienceAnalyticsBatteryHealthDevicePerformance:                      userExperienceAnalyticsBatteryHealthDevicePerformanceClient,
		UserExperienceAnalyticsBatteryHealthDeviceRuntimeHistory:                   userExperienceAnalyticsBatteryHealthDeviceRuntimeHistoryClient,
		UserExperienceAnalyticsBatteryHealthModelPerformance:                       userExperienceAnalyticsBatteryHealthModelPerformanceClient,
		UserExperienceAnalyticsBatteryHealthOsPerformance:                          userExperienceAnalyticsBatteryHealthOsPerformanceClient,
		UserExperienceAnalyticsBatteryHealthRuntimeDetail:                          userExperienceAnalyticsBatteryHealthRuntimeDetailClient,
		UserExperienceAnalyticsCategory:                                            userExperienceAnalyticsCategoryClient,
		UserExperienceAnalyticsCategoryMetricValue:                                 userExperienceAnalyticsCategoryMetricValueClient,
		UserExperienceAnalyticsDeviceMetricHistory:                                 userExperienceAnalyticsDeviceMetricHistoryClient,
		UserExperienceAnalyticsDevicePerformance:                                   userExperienceAnalyticsDevicePerformanceClient,
		UserExperienceAnalyticsDeviceScope:                                         userExperienceAnalyticsDeviceScopeClient,
		UserExperienceAnalyticsDeviceScore:                                         userExperienceAnalyticsDeviceScoreClient,
		UserExperienceAnalyticsDeviceStartupHistory:                                userExperienceAnalyticsDeviceStartupHistoryClient,
		UserExperienceAnalyticsDeviceStartupProcess:                                userExperienceAnalyticsDeviceStartupProcessClient,
		UserExperienceAnalyticsDeviceStartupProcessPerformance:                     userExperienceAnalyticsDeviceStartupProcessPerformanceClient,
		UserExperienceAnalyticsDeviceTimelineEvent:                                 userExperienceAnalyticsDeviceTimelineEventClient,
		UserExperienceAnalyticsDevicesWithoutCloudIdentity:                         userExperienceAnalyticsDevicesWithoutCloudIdentityClient,
		UserExperienceAnalyticsImpactingProcess:                                    userExperienceAnalyticsImpactingProcessClient,
		UserExperienceAnalyticsMetricHistory:                                       userExperienceAnalyticsMetricHistoryClient,
		UserExperienceAnalyticsModelScore:                                          userExperienceAnalyticsModelScoreClient,
		UserExperienceAnalyticsNotAutopilotReadyDevice:                             userExperienceAnalyticsNotAutopilotReadyDeviceClient,
		UserExperienceAnalyticsOverview:                                            userExperienceAnalyticsOverviewClient,
		UserExperienceAnalyticsRemoteConnection:                                    userExperienceAnalyticsRemoteConnectionClient,
		UserExperienceAnalyticsResourcePerformance:                                 userExperienceAnalyticsResourcePerformanceClient,
		UserExperienceAnalyticsScoreHistory:                                        userExperienceAnalyticsScoreHistoryClient,
		UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric:             userExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetricClient,
		UserExperienceAnalyticsWorkFromAnywhereMetric:                              userExperienceAnalyticsWorkFromAnywhereMetricClient,
		UserExperienceAnalyticsWorkFromAnywhereMetricMetricDevice:                  userExperienceAnalyticsWorkFromAnywhereMetricMetricDeviceClient,
		UserExperienceAnalyticsWorkFromAnywhereModelPerformance:                    userExperienceAnalyticsWorkFromAnywhereModelPerformanceClient,
		UserPfxCertificate:        userPfxCertificateClient,
		VirtualEndpoint:           virtualEndpointClient,
		VirtualEndpointAuditEvent: virtualEndpointAuditEventClient,
		VirtualEndpointBulkAction: virtualEndpointBulkActionClient,
		VirtualEndpointCloudPC:    virtualEndpointCloudPCClient,
		VirtualEndpointCrossCloudGovernmentOrganizationMapping:                          virtualEndpointCrossCloudGovernmentOrganizationMappingClient,
		VirtualEndpointDeviceImage:                                                      virtualEndpointDeviceImageClient,
		VirtualEndpointExternalPartnerSetting:                                           virtualEndpointExternalPartnerSettingClient,
		VirtualEndpointFrontLineServicePlan:                                             virtualEndpointFrontLineServicePlanClient,
		VirtualEndpointGalleryImage:                                                     virtualEndpointGalleryImageClient,
		VirtualEndpointOnPremisesConnection:                                             virtualEndpointOnPremisesConnectionClient,
		VirtualEndpointOrganizationSetting:                                              virtualEndpointOrganizationSettingClient,
		VirtualEndpointProvisioningPolicy:                                               virtualEndpointProvisioningPolicyClient,
		VirtualEndpointProvisioningPolicyAssignment:                                     virtualEndpointProvisioningPolicyAssignmentClient,
		VirtualEndpointProvisioningPolicyAssignmentAssignedUser:                         virtualEndpointProvisioningPolicyAssignmentAssignedUserClient,
		VirtualEndpointProvisioningPolicyAssignmentAssignedUserMailboxSetting:           virtualEndpointProvisioningPolicyAssignmentAssignedUserMailboxSettingClient,
		VirtualEndpointProvisioningPolicyAssignmentAssignedUserServiceProvisioningError: virtualEndpointProvisioningPolicyAssignmentAssignedUserServiceProvisioningErrorClient,
		VirtualEndpointReport:                                                           virtualEndpointReportClient,
		VirtualEndpointServicePlan:                                                      virtualEndpointServicePlanClient,
		VirtualEndpointSnapshot:                                                         virtualEndpointSnapshotClient,
		VirtualEndpointSupportedRegion:                                                  virtualEndpointSupportedRegionClient,
		VirtualEndpointUserSetting:                                                      virtualEndpointUserSettingClient,
		VirtualEndpointUserSettingAssignment:                                            virtualEndpointUserSettingAssignmentClient,
		WindowsAutopilotDeploymentProfile:                                               windowsAutopilotDeploymentProfileClient,
		WindowsAutopilotDeploymentProfileAssignedDevice:                                 windowsAutopilotDeploymentProfileAssignedDeviceClient,
		WindowsAutopilotDeploymentProfileAssignedDeviceDeploymentProfile:                windowsAutopilotDeploymentProfileAssignedDeviceDeploymentProfileClient,
		WindowsAutopilotDeploymentProfileAssignedDeviceIntendedDeploymentProfile:        windowsAutopilotDeploymentProfileAssignedDeviceIntendedDeploymentProfileClient,
		WindowsAutopilotDeploymentProfileAssignment:                                     windowsAutopilotDeploymentProfileAssignmentClient,
		WindowsAutopilotDeviceIdentity:                                                  windowsAutopilotDeviceIdentityClient,
		WindowsAutopilotDeviceIdentityDeploymentProfile:                                 windowsAutopilotDeviceIdentityDeploymentProfileClient,
		WindowsAutopilotDeviceIdentityIntendedDeploymentProfile:                         windowsAutopilotDeviceIdentityIntendedDeploymentProfileClient,
		WindowsAutopilotSetting:                                                         windowsAutopilotSettingClient,
		WindowsDriverUpdateProfile:                                                      windowsDriverUpdateProfileClient,
		WindowsDriverUpdateProfileAssignment:                                            windowsDriverUpdateProfileAssignmentClient,
		WindowsDriverUpdateProfileDriverInventory:                                       windowsDriverUpdateProfileDriverInventoryClient,
		WindowsFeatureUpdateProfile:                                                     windowsFeatureUpdateProfileClient,
		WindowsFeatureUpdateProfileAssignment:                                           windowsFeatureUpdateProfileAssignmentClient,
		WindowsInformationProtectionAppLearningSummary:                                  windowsInformationProtectionAppLearningSummaryClient,
		WindowsInformationProtectionNetworkLearningSummary:                              windowsInformationProtectionNetworkLearningSummaryClient,
		WindowsMalwareInformation:                                                       windowsMalwareInformationClient,
		WindowsMalwareInformationDeviceMalwareState:                                     windowsMalwareInformationDeviceMalwareStateClient,
		WindowsQualityUpdatePolicy:                                                      windowsQualityUpdatePolicyClient,
		WindowsQualityUpdatePolicyAssignment:                                            windowsQualityUpdatePolicyAssignmentClient,
		WindowsQualityUpdateProfile:                                                     windowsQualityUpdateProfileClient,
		WindowsQualityUpdateProfileAssignment:                                           windowsQualityUpdateProfileAssignmentClient,
		WindowsUpdateCatalogItem:                                                        windowsUpdateCatalogItemClient,
		ZebraFotaArtifact:                                                               zebraFotaArtifactClient,
		ZebraFotaConnector:                                                              zebraFotaConnectorClient,
		ZebraFotaDeployment:                                                             zebraFotaDeploymentClient,
	}, nil
}
