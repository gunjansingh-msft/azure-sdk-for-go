// +build go1.9

// Copyright 2021 Microsoft Corporation
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// This code was auto-generated by:
// github.com/Azure/azure-sdk-for-go/tools/profileBuilder

package workloadmonitor

import (
	"context"

	original "github.com/Azure/azure-sdk-for-go/services/preview/workloadmonitor/mgmt/2020-01-13-preview/workloadmonitor"
)

const (
	DefaultBaseURI = original.DefaultBaseURI
)

type HealthState = original.HealthState

const (
	Critical HealthState = original.Critical
	Disabled HealthState = original.Disabled
	Healthy  HealthState = original.Healthy
	None     HealthState = original.None
	Unknown  HealthState = original.Unknown
	Warning  HealthState = original.Warning
)

type BaseClient = original.BaseClient
type ErrorDetails = original.ErrorDetails
type ErrorResponse = original.ErrorResponse
type ErrorResponseError = original.ErrorResponseError
type HealthMonitor = original.HealthMonitor
type HealthMonitorList = original.HealthMonitorList
type HealthMonitorListIterator = original.HealthMonitorListIterator
type HealthMonitorListPage = original.HealthMonitorListPage
type HealthMonitorProperties = original.HealthMonitorProperties
type HealthMonitorStateChange = original.HealthMonitorStateChange
type HealthMonitorStateChangeList = original.HealthMonitorStateChangeList
type HealthMonitorStateChangeListIterator = original.HealthMonitorStateChangeListIterator
type HealthMonitorStateChangeListPage = original.HealthMonitorStateChangeListPage
type HealthMonitorStateChangeProperties = original.HealthMonitorStateChangeProperties
type HealthMonitorsClient = original.HealthMonitorsClient
type Operation = original.Operation
type OperationDisplay = original.OperationDisplay
type OperationList = original.OperationList
type OperationListIterator = original.OperationListIterator
type OperationListPage = original.OperationListPage
type OperationsClient = original.OperationsClient
type Resource = original.Resource

func New() BaseClient {
	return original.New()
}
func NewHealthMonitorListIterator(page HealthMonitorListPage) HealthMonitorListIterator {
	return original.NewHealthMonitorListIterator(page)
}
func NewHealthMonitorListPage(cur HealthMonitorList, getNextPage func(context.Context, HealthMonitorList) (HealthMonitorList, error)) HealthMonitorListPage {
	return original.NewHealthMonitorListPage(cur, getNextPage)
}
func NewHealthMonitorStateChangeListIterator(page HealthMonitorStateChangeListPage) HealthMonitorStateChangeListIterator {
	return original.NewHealthMonitorStateChangeListIterator(page)
}
func NewHealthMonitorStateChangeListPage(cur HealthMonitorStateChangeList, getNextPage func(context.Context, HealthMonitorStateChangeList) (HealthMonitorStateChangeList, error)) HealthMonitorStateChangeListPage {
	return original.NewHealthMonitorStateChangeListPage(cur, getNextPage)
}
func NewHealthMonitorsClient() HealthMonitorsClient {
	return original.NewHealthMonitorsClient()
}
func NewHealthMonitorsClientWithBaseURI(baseURI string) HealthMonitorsClient {
	return original.NewHealthMonitorsClientWithBaseURI(baseURI)
}
func NewOperationListIterator(page OperationListPage) OperationListIterator {
	return original.NewOperationListIterator(page)
}
func NewOperationListPage(cur OperationList, getNextPage func(context.Context, OperationList) (OperationList, error)) OperationListPage {
	return original.NewOperationListPage(cur, getNextPage)
}
func NewOperationsClient() OperationsClient {
	return original.NewOperationsClient()
}
func NewOperationsClientWithBaseURI(baseURI string) OperationsClient {
	return original.NewOperationsClientWithBaseURI(baseURI)
}
func NewWithBaseURI(baseURI string) BaseClient {
	return original.NewWithBaseURI(baseURI)
}
func PossibleHealthStateValues() []HealthState {
	return original.PossibleHealthStateValues()
}
func UserAgent() string {
	return original.UserAgent() + " profiles/preview"
}
func Version() string {
	return original.Version()
}
