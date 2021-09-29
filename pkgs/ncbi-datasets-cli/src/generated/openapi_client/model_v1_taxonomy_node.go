/*
NCBI Datasets API

### NCBI Datasets is a resource that lets you easily gather data from NCBI. The Datasets API is still in alpha, and we're updating it often to add new functionality, iron out bugs and enhance usability. For some larger downloads, you may want to download a [dehydrated bag](https://www.ncbi.nlm.nih.gov/datasets/docs/rehydrate/), and retrieve the individual data files at a later time. 

API version: v1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package datasets

import (
	"encoding/json"
)

// V1TaxonomyNode struct for V1TaxonomyNode
type V1TaxonomyNode struct {
	TaxId *int32 `json:"tax_id,omitempty"`
	OrganismName *string `json:"organism_name,omitempty"`
	CommonName *string `json:"common_name,omitempty"`
	Lineage *[]int32 `json:"lineage,omitempty"`
	Children *[]int32 `json:"children,omitempty"`
	DescendentWithDescribedSpeciesNamesCount *int32 `json:"descendent_with_described_species_names_count,omitempty"`
	Rank *V1OrganismRankType `json:"rank,omitempty"`
	HasDescribedSpeciesName *bool `json:"has_described_species_name,omitempty"`
	Counts *[]V1TaxonomyNodeCountByType `json:"counts,omitempty"`
	MinOrd *int32 `json:"min_ord,omitempty"`
	MaxOrd *int32 `json:"max_ord,omitempty"`
	IndexedNames *[]string `json:"_indexed_names,omitempty"`
}

// NewV1TaxonomyNode instantiates a new V1TaxonomyNode object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV1TaxonomyNode() *V1TaxonomyNode {
	this := V1TaxonomyNode{}
	var rank V1OrganismRankType = V1ORGANISMRANKTYPE_NO_RANK
	this.Rank = &rank
	return &this
}

// NewV1TaxonomyNodeWithDefaults instantiates a new V1TaxonomyNode object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV1TaxonomyNodeWithDefaults() *V1TaxonomyNode {
	this := V1TaxonomyNode{}
	var rank V1OrganismRankType = V1ORGANISMRANKTYPE_NO_RANK
	this.Rank = &rank
	return &this
}

// GetTaxId returns the TaxId field value if set, zero value otherwise.
func (o *V1TaxonomyNode) GetTaxId() int32 {
	if o == nil || o.TaxId == nil {
		var ret int32
		return ret
	}
	return *o.TaxId
}

// GetTaxIdOk returns a tuple with the TaxId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1TaxonomyNode) GetTaxIdOk() (*int32, bool) {
	if o == nil || o.TaxId == nil {
		return nil, false
	}
	return o.TaxId, true
}

// HasTaxId returns a boolean if a field has been set.
func (o *V1TaxonomyNode) HasTaxId() bool {
	if o != nil && o.TaxId != nil {
		return true
	}

	return false
}

// SetTaxId gets a reference to the given int32 and assigns it to the TaxId field.
func (o *V1TaxonomyNode) SetTaxId(v int32) {
	o.TaxId = &v
}

// GetOrganismName returns the OrganismName field value if set, zero value otherwise.
func (o *V1TaxonomyNode) GetOrganismName() string {
	if o == nil || o.OrganismName == nil {
		var ret string
		return ret
	}
	return *o.OrganismName
}

// GetOrganismNameOk returns a tuple with the OrganismName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1TaxonomyNode) GetOrganismNameOk() (*string, bool) {
	if o == nil || o.OrganismName == nil {
		return nil, false
	}
	return o.OrganismName, true
}

// HasOrganismName returns a boolean if a field has been set.
func (o *V1TaxonomyNode) HasOrganismName() bool {
	if o != nil && o.OrganismName != nil {
		return true
	}

	return false
}

// SetOrganismName gets a reference to the given string and assigns it to the OrganismName field.
func (o *V1TaxonomyNode) SetOrganismName(v string) {
	o.OrganismName = &v
}

// GetCommonName returns the CommonName field value if set, zero value otherwise.
func (o *V1TaxonomyNode) GetCommonName() string {
	if o == nil || o.CommonName == nil {
		var ret string
		return ret
	}
	return *o.CommonName
}

// GetCommonNameOk returns a tuple with the CommonName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1TaxonomyNode) GetCommonNameOk() (*string, bool) {
	if o == nil || o.CommonName == nil {
		return nil, false
	}
	return o.CommonName, true
}

// HasCommonName returns a boolean if a field has been set.
func (o *V1TaxonomyNode) HasCommonName() bool {
	if o != nil && o.CommonName != nil {
		return true
	}

	return false
}

// SetCommonName gets a reference to the given string and assigns it to the CommonName field.
func (o *V1TaxonomyNode) SetCommonName(v string) {
	o.CommonName = &v
}

// GetLineage returns the Lineage field value if set, zero value otherwise.
func (o *V1TaxonomyNode) GetLineage() []int32 {
	if o == nil || o.Lineage == nil {
		var ret []int32
		return ret
	}
	return *o.Lineage
}

// GetLineageOk returns a tuple with the Lineage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1TaxonomyNode) GetLineageOk() (*[]int32, bool) {
	if o == nil || o.Lineage == nil {
		return nil, false
	}
	return o.Lineage, true
}

// HasLineage returns a boolean if a field has been set.
func (o *V1TaxonomyNode) HasLineage() bool {
	if o != nil && o.Lineage != nil {
		return true
	}

	return false
}

// SetLineage gets a reference to the given []int32 and assigns it to the Lineage field.
func (o *V1TaxonomyNode) SetLineage(v []int32) {
	o.Lineage = &v
}

// GetChildren returns the Children field value if set, zero value otherwise.
func (o *V1TaxonomyNode) GetChildren() []int32 {
	if o == nil || o.Children == nil {
		var ret []int32
		return ret
	}
	return *o.Children
}

// GetChildrenOk returns a tuple with the Children field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1TaxonomyNode) GetChildrenOk() (*[]int32, bool) {
	if o == nil || o.Children == nil {
		return nil, false
	}
	return o.Children, true
}

// HasChildren returns a boolean if a field has been set.
func (o *V1TaxonomyNode) HasChildren() bool {
	if o != nil && o.Children != nil {
		return true
	}

	return false
}

// SetChildren gets a reference to the given []int32 and assigns it to the Children field.
func (o *V1TaxonomyNode) SetChildren(v []int32) {
	o.Children = &v
}

// GetDescendentWithDescribedSpeciesNamesCount returns the DescendentWithDescribedSpeciesNamesCount field value if set, zero value otherwise.
func (o *V1TaxonomyNode) GetDescendentWithDescribedSpeciesNamesCount() int32 {
	if o == nil || o.DescendentWithDescribedSpeciesNamesCount == nil {
		var ret int32
		return ret
	}
	return *o.DescendentWithDescribedSpeciesNamesCount
}

// GetDescendentWithDescribedSpeciesNamesCountOk returns a tuple with the DescendentWithDescribedSpeciesNamesCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1TaxonomyNode) GetDescendentWithDescribedSpeciesNamesCountOk() (*int32, bool) {
	if o == nil || o.DescendentWithDescribedSpeciesNamesCount == nil {
		return nil, false
	}
	return o.DescendentWithDescribedSpeciesNamesCount, true
}

// HasDescendentWithDescribedSpeciesNamesCount returns a boolean if a field has been set.
func (o *V1TaxonomyNode) HasDescendentWithDescribedSpeciesNamesCount() bool {
	if o != nil && o.DescendentWithDescribedSpeciesNamesCount != nil {
		return true
	}

	return false
}

// SetDescendentWithDescribedSpeciesNamesCount gets a reference to the given int32 and assigns it to the DescendentWithDescribedSpeciesNamesCount field.
func (o *V1TaxonomyNode) SetDescendentWithDescribedSpeciesNamesCount(v int32) {
	o.DescendentWithDescribedSpeciesNamesCount = &v
}

// GetRank returns the Rank field value if set, zero value otherwise.
func (o *V1TaxonomyNode) GetRank() V1OrganismRankType {
	if o == nil || o.Rank == nil {
		var ret V1OrganismRankType
		return ret
	}
	return *o.Rank
}

// GetRankOk returns a tuple with the Rank field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1TaxonomyNode) GetRankOk() (*V1OrganismRankType, bool) {
	if o == nil || o.Rank == nil {
		return nil, false
	}
	return o.Rank, true
}

// HasRank returns a boolean if a field has been set.
func (o *V1TaxonomyNode) HasRank() bool {
	if o != nil && o.Rank != nil {
		return true
	}

	return false
}

// SetRank gets a reference to the given V1OrganismRankType and assigns it to the Rank field.
func (o *V1TaxonomyNode) SetRank(v V1OrganismRankType) {
	o.Rank = &v
}

// GetHasDescribedSpeciesName returns the HasDescribedSpeciesName field value if set, zero value otherwise.
func (o *V1TaxonomyNode) GetHasDescribedSpeciesName() bool {
	if o == nil || o.HasDescribedSpeciesName == nil {
		var ret bool
		return ret
	}
	return *o.HasDescribedSpeciesName
}

// GetHasDescribedSpeciesNameOk returns a tuple with the HasDescribedSpeciesName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1TaxonomyNode) GetHasDescribedSpeciesNameOk() (*bool, bool) {
	if o == nil || o.HasDescribedSpeciesName == nil {
		return nil, false
	}
	return o.HasDescribedSpeciesName, true
}

// HasHasDescribedSpeciesName returns a boolean if a field has been set.
func (o *V1TaxonomyNode) HasHasDescribedSpeciesName() bool {
	if o != nil && o.HasDescribedSpeciesName != nil {
		return true
	}

	return false
}

// SetHasDescribedSpeciesName gets a reference to the given bool and assigns it to the HasDescribedSpeciesName field.
func (o *V1TaxonomyNode) SetHasDescribedSpeciesName(v bool) {
	o.HasDescribedSpeciesName = &v
}

// GetCounts returns the Counts field value if set, zero value otherwise.
func (o *V1TaxonomyNode) GetCounts() []V1TaxonomyNodeCountByType {
	if o == nil || o.Counts == nil {
		var ret []V1TaxonomyNodeCountByType
		return ret
	}
	return *o.Counts
}

// GetCountsOk returns a tuple with the Counts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1TaxonomyNode) GetCountsOk() (*[]V1TaxonomyNodeCountByType, bool) {
	if o == nil || o.Counts == nil {
		return nil, false
	}
	return o.Counts, true
}

// HasCounts returns a boolean if a field has been set.
func (o *V1TaxonomyNode) HasCounts() bool {
	if o != nil && o.Counts != nil {
		return true
	}

	return false
}

// SetCounts gets a reference to the given []V1TaxonomyNodeCountByType and assigns it to the Counts field.
func (o *V1TaxonomyNode) SetCounts(v []V1TaxonomyNodeCountByType) {
	o.Counts = &v
}

// GetMinOrd returns the MinOrd field value if set, zero value otherwise.
func (o *V1TaxonomyNode) GetMinOrd() int32 {
	if o == nil || o.MinOrd == nil {
		var ret int32
		return ret
	}
	return *o.MinOrd
}

// GetMinOrdOk returns a tuple with the MinOrd field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1TaxonomyNode) GetMinOrdOk() (*int32, bool) {
	if o == nil || o.MinOrd == nil {
		return nil, false
	}
	return o.MinOrd, true
}

// HasMinOrd returns a boolean if a field has been set.
func (o *V1TaxonomyNode) HasMinOrd() bool {
	if o != nil && o.MinOrd != nil {
		return true
	}

	return false
}

// SetMinOrd gets a reference to the given int32 and assigns it to the MinOrd field.
func (o *V1TaxonomyNode) SetMinOrd(v int32) {
	o.MinOrd = &v
}

// GetMaxOrd returns the MaxOrd field value if set, zero value otherwise.
func (o *V1TaxonomyNode) GetMaxOrd() int32 {
	if o == nil || o.MaxOrd == nil {
		var ret int32
		return ret
	}
	return *o.MaxOrd
}

// GetMaxOrdOk returns a tuple with the MaxOrd field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1TaxonomyNode) GetMaxOrdOk() (*int32, bool) {
	if o == nil || o.MaxOrd == nil {
		return nil, false
	}
	return o.MaxOrd, true
}

// HasMaxOrd returns a boolean if a field has been set.
func (o *V1TaxonomyNode) HasMaxOrd() bool {
	if o != nil && o.MaxOrd != nil {
		return true
	}

	return false
}

// SetMaxOrd gets a reference to the given int32 and assigns it to the MaxOrd field.
func (o *V1TaxonomyNode) SetMaxOrd(v int32) {
	o.MaxOrd = &v
}

// GetIndexedNames returns the IndexedNames field value if set, zero value otherwise.
func (o *V1TaxonomyNode) GetIndexedNames() []string {
	if o == nil || o.IndexedNames == nil {
		var ret []string
		return ret
	}
	return *o.IndexedNames
}

// GetIndexedNamesOk returns a tuple with the IndexedNames field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1TaxonomyNode) GetIndexedNamesOk() (*[]string, bool) {
	if o == nil || o.IndexedNames == nil {
		return nil, false
	}
	return o.IndexedNames, true
}

// HasIndexedNames returns a boolean if a field has been set.
func (o *V1TaxonomyNode) HasIndexedNames() bool {
	if o != nil && o.IndexedNames != nil {
		return true
	}

	return false
}

// SetIndexedNames gets a reference to the given []string and assigns it to the IndexedNames field.
func (o *V1TaxonomyNode) SetIndexedNames(v []string) {
	o.IndexedNames = &v
}

func (o V1TaxonomyNode) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.TaxId != nil  {
		toSerialize["tax_id"] = o.TaxId
	}
	if o.OrganismName != nil  {
		toSerialize["organism_name"] = o.OrganismName
	}
	if o.CommonName != nil  {
		toSerialize["common_name"] = o.CommonName
	}
	if o.Lineage != nil && len(o.GetLineage()) > 0  {
		toSerialize["lineage"] = o.Lineage
	}
	if o.Children != nil && len(o.GetChildren()) > 0  {
		toSerialize["children"] = o.Children
	}
	if o.DescendentWithDescribedSpeciesNamesCount != nil  {
		toSerialize["descendent_with_described_species_names_count"] = o.DescendentWithDescribedSpeciesNamesCount
	}
	if o.Rank != nil  {
		toSerialize["rank"] = o.Rank
	}
	if o.HasDescribedSpeciesName != nil  {
		toSerialize["has_described_species_name"] = o.HasDescribedSpeciesName
	}
	if o.Counts != nil && len(o.GetCounts()) > 0  {
		toSerialize["counts"] = o.Counts
	}
	if o.MinOrd != nil  {
		toSerialize["min_ord"] = o.MinOrd
	}
	if o.MaxOrd != nil  {
		toSerialize["max_ord"] = o.MaxOrd
	}
	if o.IndexedNames != nil && len(o.GetIndexedNames()) > 0  {
		toSerialize["_indexed_names"] = o.IndexedNames
	}
	return json.Marshal(toSerialize)
}

type NullableV1TaxonomyNode struct {
	value *V1TaxonomyNode
	isSet bool
}

func (v NullableV1TaxonomyNode) Get() *V1TaxonomyNode {
	return v.value
}

func (v *NullableV1TaxonomyNode) Set(val *V1TaxonomyNode) {
	v.value = val
	v.isSet = true
}

func (v NullableV1TaxonomyNode) IsSet() bool {
	return v.isSet
}

func (v *NullableV1TaxonomyNode) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV1TaxonomyNode(val *V1TaxonomyNode) *NullableV1TaxonomyNode {
	return &NullableV1TaxonomyNode{value: val, isSet: true}
}

func (v NullableV1TaxonomyNode) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV1TaxonomyNode) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


