// Copyright Amazon.com Inc. or its affiliates. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License"). You may
// not use this file except in compliance with the License. A copy of the
// License is located at
//
//     http://aws.amazon.com/apache2.0/
//
// or in the "license" file accompanying this file. This file is distributed
// on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either
// express or implied. See the License for the specific language governing
// permissions and limitations under the License.

// Code generated by ack-generate. DO NOT EDIT.

package v1alpha1

import (
	ackv1alpha1 "github.com/aws-controllers-k8s/runtime/apis/core/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// DBParameterGroupSpec defines the desired state of DBParameterGroup.
//
// Contains the details of an Amazon RDS DB parameter group.
//
// This data type is used as a response element in the DescribeDBParameterGroups
// action.
type DBParameterGroupSpec struct {
	// The description for the DB parameter group.
	// +kubebuilder:validation:Required
	Description *string `json:"description"`
	// The DB parameter group family name. A DB parameter group can be associated
	// with one and only one DB parameter group family, and can be applied only
	// to a DB instance running a database engine and engine version compatible
	// with that DB parameter group family.
	//
	// To list all of the available parameter group families for a DB engine, use
	// the following command:
	//
	// aws rds describe-db-engine-versions --query "DBEngineVersions[].DBParameterGroupFamily"
	// --engine <engine>
	//
	// For example, to list all of the available parameter group families for the
	// MySQL DB engine, use the following command:
	//
	// aws rds describe-db-engine-versions --query "DBEngineVersions[].DBParameterGroupFamily"
	// --engine mysql
	//
	// The output contains duplicates.
	//
	// The following are the valid DB engine values:
	//
	//    * aurora (for MySQL 5.6-compatible Aurora)
	//
	//    * aurora-mysql (for MySQL 5.7-compatible Aurora)
	//
	//    * aurora-postgresql
	//
	//    * mariadb
	//
	//    * mysql
	//
	//    * oracle-ee
	//
	//    * oracle-ee-cdb
	//
	//    * oracle-se2
	//
	//    * oracle-se2-cdb
	//
	//    * postgres
	//
	//    * sqlserver-ee
	//
	//    * sqlserver-se
	//
	//    * sqlserver-ex
	//
	//    * sqlserver-web
	// +kubebuilder:validation:Required
	Family *string `json:"family"`
	// The name of the DB parameter group.
	//
	// Constraints:
	//
	//    * Must be 1 to 255 letters, numbers, or hyphens.
	//
	//    * First character must be a letter
	//
	//    * Can't end with a hyphen or contain two consecutive hyphens
	//
	// This value is stored as a lowercase string.
	// +kubebuilder:validation:Required
	Name *string `json:"name"`
	// An array of parameter names, values, and the application methods for the
	// parameter update. At least one parameter name, value, and application method
	// method must be supplied; later arguments are optional. A maximum of 20 parameters
	// can be modified in a single request.
	//
	// Valid Values (for the application method): immediate | pending-reboot
	//
	// You can use the immediate value with dynamic parameters only. You can use
	// the pending-reboot value for both dynamic and static parameters.
	//
	// When the application method is immediate, changes to dynamic parameters are
	// applied immediately to the DB instances associated with the parameter group.
	// When the application method is pending-reboot, changes to dynamic and static
	// parameters are applied after a reboot without failover to the DB instances
	// associated with the parameter group.
	Parameters []*Parameter `json:"parameters,omitempty"`
	// Tags to assign to the DB parameter group.
	Tags []*Tag `json:"tags,omitempty"`
}

// DBParameterGroupStatus defines the observed state of DBParameterGroup
type DBParameterGroupStatus struct {
	// All CRs managed by ACK have a common `Status.ACKResourceMetadata` member
	// that is used to contain resource sync state, account ownership,
	// constructed ARN for the resource
	// +kubebuilder:validation:Optional
	ACKResourceMetadata *ackv1alpha1.ResourceMetadata `json:"ackResourceMetadata"`
	// All CRS managed by ACK have a common `Status.Conditions` member that
	// contains a collection of `ackv1alpha1.Condition` objects that describe
	// the various terminal states of the CR and its backend AWS service API
	// resource
	// +kubebuilder:validation:Optional
	Conditions []*ackv1alpha1.Condition `json:"conditions"`
}

// DBParameterGroup is the Schema for the DBParameterGroups API
// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
type DBParameterGroup struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DBParameterGroupSpec   `json:"spec,omitempty"`
	Status            DBParameterGroupStatus `json:"status,omitempty"`
}

// DBParameterGroupList contains a list of DBParameterGroup
// +kubebuilder:object:root=true
type DBParameterGroupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DBParameterGroup `json:"items"`
}

func init() {
	SchemeBuilder.Register(&DBParameterGroup{}, &DBParameterGroupList{})
}
