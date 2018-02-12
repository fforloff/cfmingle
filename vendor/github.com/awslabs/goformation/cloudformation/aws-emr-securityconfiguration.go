package cloudformation

import (
	"encoding/json"
	"errors"
	"fmt"
)

// AWSEMRSecurityConfiguration AWS CloudFormation Resource (AWS::EMR::SecurityConfiguration)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-emr-securityconfiguration.html
type AWSEMRSecurityConfiguration struct {

	// Name AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-emr-securityconfiguration.html#cfn-emr-securityconfiguration-name
	Name string `json:"Name,omitempty"`

	// SecurityConfiguration AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-emr-securityconfiguration.html#cfn-emr-securityconfiguration-securityconfiguration
	SecurityConfiguration interface{} `json:"SecurityConfiguration,omitempty"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSEMRSecurityConfiguration) AWSCloudFormationType() string {
	return "AWS::EMR::SecurityConfiguration"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSEMRSecurityConfiguration) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}

// MarshalJSON is a custom JSON marshalling hook that embeds this object into
// an AWS CloudFormation JSON resource's 'Properties' field and adds a 'Type'.
func (r *AWSEMRSecurityConfiguration) MarshalJSON() ([]byte, error) {
	type Properties AWSEMRSecurityConfiguration
	return json.Marshal(&struct {
		Type       string
		Properties Properties
	}{
		Type:       r.AWSCloudFormationType(),
		Properties: (Properties)(*r),
	})
}

// UnmarshalJSON is a custom JSON unmarshalling hook that strips the outer
// AWS CloudFormation resource object, and just keeps the 'Properties' field.
func (r *AWSEMRSecurityConfiguration) UnmarshalJSON(b []byte) error {
	type Properties AWSEMRSecurityConfiguration
	res := &struct {
		Type       string
		Properties *Properties
	}{}
	if err := json.Unmarshal(b, &res); err != nil {
		fmt.Printf("ERROR: %s\n", err)
		return err
	}
	*r = AWSEMRSecurityConfiguration(*res.Properties)
	return nil
}

// GetAllAWSEMRSecurityConfigurationResources retrieves all AWSEMRSecurityConfiguration items from an AWS CloudFormation template
func (t *Template) GetAllAWSEMRSecurityConfigurationResources() map[string]AWSEMRSecurityConfiguration {
	results := map[string]AWSEMRSecurityConfiguration{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case AWSEMRSecurityConfiguration:
			// We found a strongly typed resource of the correct type; use it
			results[name] = resource
		case map[string]interface{}:
			// We found an untyped resource (likely from JSON) which *might* be
			// the correct type, but we need to check it's 'Type' field
			if resType, ok := resource["Type"]; ok {
				if resType == "AWS::EMR::SecurityConfiguration" {
					// The resource is correct, unmarshal it into the results
					if b, err := json.Marshal(resource); err == nil {
						var result AWSEMRSecurityConfiguration
						if err := json.Unmarshal(b, &result); err == nil {
							results[name] = result
						}
					}
				}
			}
		}
	}
	return results
}

// GetAWSEMRSecurityConfigurationWithName retrieves all AWSEMRSecurityConfiguration items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSEMRSecurityConfigurationWithName(name string) (AWSEMRSecurityConfiguration, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case AWSEMRSecurityConfiguration:
			// We found a strongly typed resource of the correct type; use it
			return resource, nil
		case map[string]interface{}:
			// We found an untyped resource (likely from JSON) which *might* be
			// the correct type, but we need to check it's 'Type' field
			if resType, ok := resource["Type"]; ok {
				if resType == "AWS::EMR::SecurityConfiguration" {
					// The resource is correct, unmarshal it into the results
					if b, err := json.Marshal(resource); err == nil {
						var result AWSEMRSecurityConfiguration
						if err := json.Unmarshal(b, &result); err == nil {
							return result, nil
						}
					}
				}
			}
		}
	}
	return AWSEMRSecurityConfiguration{}, errors.New("resource not found")
}
