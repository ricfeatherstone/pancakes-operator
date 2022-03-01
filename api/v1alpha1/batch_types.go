/*
Copyright 2022 ricfeatherstone.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// Flour defines the flour to use.
type Flour struct {
	// Type the type of flour
	// +kubebuilder:validation:Enum=Plain;SelfRaising
	Type string `json:"type"`
	// Amount the amount of flour
	Amount int `json:"amount"`
	// Weight the weight of flour
	// +kubebuilder:validation:Enum=Grams;Kilograms
	Weight string `json:"weight"`
}

// Eggs defines the eggs to use.
type Eggs struct {
	// Number the number of eggs to use.
	Number int `json:"number"`
	// Size the size of the eggs to use
	// +kubebuilder:validation:Enum=Small;Medium;Large
	Size string `json:"size"`
}

// Milk defines the milk to use.
type Milk struct {
	// Amount the amount of milk
	Amount int `json:"amount"`
	// Unit the measuring unit
	// +kubebuilder:validation:Enum=MilliLitre;Litre
	Unit string `json:"unit"`
}

// Oil defines the oil to use.
type Oil struct {
	// Type the type of oil
	// +kubebuilder:validation:Enum=Sunflower;Vegetable
	Type string `json:"type"`
	// Amount the amount of oil
	Amount int `json:"amount"`
	// Unit the measuring unit
	// +kubebuilder:validation:Enum=TeaSpoon;TableSpoon
	Unit string `json:"unit"`
}

// LemonWedges defines the lemon wedges to use.
type LemonWedges struct {
	Number int `json:"number"`
}

// Sugar defines the sugar to use.
type Sugar struct {
	// Amount the amount of sugar
	Amount int `json:"amount"`
	// Unit the measuring unit
	// +kubebuilder:validation:Enum=Grams;Kilograms
	Unit string `json:"unit"`
}

// BatchSpec defines the desired state of Batch
type BatchSpec struct {
	// Flour the flour to use.
	Flour `json:"flour"`
	// Eggs the eggs to use.
	Eggs `json:"eggs"`
	// Oil the oil to use.
	Milk `json:"milk"`
	// Milk the milk to use.
	Oil `json:"oil"`
	// LemonWedges the lemon wedges to use.
	LemonWedges `json:"lemonWedges,omitempty"`
	// Sugar the sugar to use.
	Sugar `json:"sugar,omitempty"`
}

// BatchStatus defines the observed state of Batch
type BatchStatus struct {
	// Conditions represent the latest available observations of a batch's state
	Conditions []metav1.Condition `json:"conditions"`
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// Batch is the Schema for the batches API
type Batch struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   BatchSpec   `json:"spec,omitempty"`
	Status BatchStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// BatchList contains a list of Batch
type BatchList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Batch `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Batch{}, &BatchList{})
}
