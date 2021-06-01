/*
 * NCBI Datasets API
 *
 * NCBI service to query and download biological sequence data across all domains of life from NCBI databases.
 *
 * API version: v1alpha
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package datasets
// V1alpha1AssemblyDatasetAvailability struct for V1alpha1AssemblyDatasetAvailability
type V1alpha1AssemblyDatasetAvailability struct {
	InvalidAssemblies []string `json:"invalid_assemblies,omitempty"`
	Reason string `json:"reason,omitempty"`
	ValidAssemblies []string `json:"valid_assemblies,omitempty"`
}
