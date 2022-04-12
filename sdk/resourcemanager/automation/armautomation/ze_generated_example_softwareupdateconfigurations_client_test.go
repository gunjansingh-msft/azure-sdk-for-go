//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armautomation_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/automation/armautomation"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/automation/resource-manager/Microsoft.Automation/stable/2019-06-01/examples/softwareUpdateConfiguration/createSoftwareUpdateConfiguration.json
func ExampleSoftwareUpdateConfigurationsClient_Create() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
		return
	}
	ctx := context.Background()
	client, err := armautomation.NewSoftwareUpdateConfigurationsClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
		return
	}
	res, err := client.Create(ctx,
		"<resource-group-name>",
		"<automation-account-name>",
		"<software-update-configuration-name>",
		armautomation.SoftwareUpdateConfiguration{
			Properties: &armautomation.SoftwareUpdateConfigurationProperties{
				ScheduleInfo: &armautomation.SUCScheduleProperties{
					AdvancedSchedule: &armautomation.AdvancedSchedule{
						WeekDays: []*string{
							to.Ptr("Monday"),
							to.Ptr("Thursday")},
					},
					ExpiryTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-11-09T11:22:57+00:00"); return t }()),
					Frequency:  to.Ptr(armautomation.ScheduleFrequencyHour),
					Interval:   to.Ptr[int64](1),
					StartTime:  to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-10-19T12:22:57+00:00"); return t }()),
					TimeZone:   to.Ptr("<time-zone>"),
				},
				Tasks: &armautomation.SoftwareUpdateConfigurationTasks{
					PostTask: &armautomation.TaskProperties{
						Source: to.Ptr("<source>"),
					},
					PreTask: &armautomation.TaskProperties{
						Parameters: map[string]*string{
							"COMPUTERNAME": to.Ptr("Computer1"),
						},
						Source: to.Ptr("<source>"),
					},
				},
				UpdateConfiguration: &armautomation.UpdateConfiguration{
					AzureVirtualMachines: []*string{
						to.Ptr("/subscriptions/5ae68d89-69a4-454f-b5ce-e443cc4e0067/resourceGroups/myresources/providers/Microsoft.Compute/virtualMachines/vm-01"),
						to.Ptr("/subscriptions/5ae68d89-69a4-454f-b5ce-e443cc4e0067/resourceGroups/myresources/providers/Microsoft.Compute/virtualMachines/vm-02"),
						to.Ptr("/subscriptions/5ae68d89-69a4-454f-b5ce-e443cc4e0067/resourceGroups/myresources/providers/Microsoft.Compute/virtualMachines/vm-03")},
					Duration: to.Ptr("<duration>"),
					NonAzureComputerNames: []*string{
						to.Ptr("box1.contoso.com"),
						to.Ptr("box2.contoso.com")},
					OperatingSystem: to.Ptr(armautomation.OperatingSystemTypeWindows),
					Targets: &armautomation.TargetProperties{
						AzureQueries: []*armautomation.AzureQueryProperties{
							{
								Locations: []*string{
									to.Ptr("Japan East"),
									to.Ptr("UK South")},
								Scope: []*string{
									to.Ptr("/subscriptions/5ae68d89-69a4-454f-b5ce-e443cc4e0067/resourceGroups/myresources"),
									to.Ptr("/subscriptions/5ae68d89-69a4-454f-b5ce-e443cc4e0067")},
								TagSettings: &armautomation.TagSettingsProperties{
									FilterOperator: to.Ptr(armautomation.TagOperatorsAll),
									Tags: map[string][]*string{
										"tag1": {
											to.Ptr("tag1Value1"),
											to.Ptr("tag1Value2"),
											to.Ptr("tag1Value3")},
										"tag2": {
											to.Ptr("tag2Value1"),
											to.Ptr("tag2Value2"),
											to.Ptr("tag2Value3")},
									},
								},
							}},
						NonAzureQueries: []*armautomation.NonAzureQueryProperties{
							{
								FunctionAlias: to.Ptr("<function-alias>"),
								WorkspaceID:   to.Ptr("<workspace-id>"),
							},
							{
								FunctionAlias: to.Ptr("<function-alias>"),
								WorkspaceID:   to.Ptr("<workspace-id>"),
							}},
					},
					Windows: &armautomation.WindowsProperties{
						ExcludedKbNumbers: []*string{
							to.Ptr("168934"),
							to.Ptr("168973")},
						IncludedUpdateClassifications: to.Ptr(armautomation.WindowsUpdateClassesCritical),
						RebootSetting:                 to.Ptr("<reboot-setting>"),
					},
				},
			},
		},
		&armautomation.SoftwareUpdateConfigurationsClientCreateOptions{ClientRequestID: nil})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
		return
	}
	// TODO: use response item
	_ = res
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/automation/resource-manager/Microsoft.Automation/stable/2019-06-01/examples/softwareUpdateConfiguration/getSoftwareUpdateConfigurationByName.json
func ExampleSoftwareUpdateConfigurationsClient_GetByName() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
		return
	}
	ctx := context.Background()
	client, err := armautomation.NewSoftwareUpdateConfigurationsClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
		return
	}
	res, err := client.GetByName(ctx,
		"<resource-group-name>",
		"<automation-account-name>",
		"<software-update-configuration-name>",
		&armautomation.SoftwareUpdateConfigurationsClientGetByNameOptions{ClientRequestID: nil})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
		return
	}
	// TODO: use response item
	_ = res
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/automation/resource-manager/Microsoft.Automation/stable/2019-06-01/examples/softwareUpdateConfiguration/deleteSoftwareUpdateConfiguration.json
func ExampleSoftwareUpdateConfigurationsClient_Delete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
		return
	}
	ctx := context.Background()
	client, err := armautomation.NewSoftwareUpdateConfigurationsClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
		return
	}
	_, err = client.Delete(ctx,
		"<resource-group-name>",
		"<automation-account-name>",
		"<software-update-configuration-name>",
		&armautomation.SoftwareUpdateConfigurationsClientDeleteOptions{ClientRequestID: nil})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
		return
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/automation/resource-manager/Microsoft.Automation/stable/2019-06-01/examples/softwareUpdateConfiguration/listSoftwareUpdateConfigurations.json
func ExampleSoftwareUpdateConfigurationsClient_List() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
		return
	}
	ctx := context.Background()
	client, err := armautomation.NewSoftwareUpdateConfigurationsClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
		return
	}
	res, err := client.List(ctx,
		"<resource-group-name>",
		"<automation-account-name>",
		&armautomation.SoftwareUpdateConfigurationsClientListOptions{ClientRequestID: nil,
			Filter: nil,
		})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
		return
	}
	// TODO: use response item
	_ = res
}
