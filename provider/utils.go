package provider

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"log"
)

func AddAttributeValueToResourceDataAndProcessValueWithHandleError(getValueOperation func() (interface{}, error), processValue func(interface{}) (interface{}, error), d *schema.ResourceData, attributeName string, attributeDescription string, diags diag.Diagnostics) diag.Diagnostics {
	log.Printf("[Attribute: %s]: Entering AddAttributeValueToResourceDataAndProcessValueWithHandleError()", attributeName)
	defer log.Printf("[Attribute: %s]: Exiting AddAttributeValueToResourceDataAndProcessValueWithHandleError()", attributeName)
	log.Printf("[Attribute: %s]: Running function to get initial value for processing", attributeName)
	value, err := getValueOperation()
	if err != nil {
		return append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  err.Error(),
		})
	} else {
		log.Printf("[Attribute: %s]: Retrieved value %v", attributeName, value)
		return AddAttributeValueToResourceDataAndHandleError(func() (interface{}, error) { return processValue(value) }, d, attributeName, attributeDescription, diags)
	}
}

func AddAttributeValueToResourceDataAndHandleError(operation func() (interface{}, error), d *schema.ResourceData, attributeName string, attributeDescription string, diags diag.Diagnostics) diag.Diagnostics {
	log.Printf("[Attribute: %s]: Entering AddAttributeValueToResourceDataAndHandleError()", attributeName)
	defer log.Printf("[Attribute: %s]: Exiting AddAttributeValueToResourceDataAndHandleError()", attributeName)
	log.Printf("[Attribute: %s]: Running function to get value", attributeName)
	if attributeValue, attributeError := operation(); attributeError != nil {
		log.Printf("[Attribute: %s]: Failed to get %s with error %v", attributeName, attributeDescription, attributeError)
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  attributeError.Error(),
		})
	} else {
		diags = AddAttributeValueToResourceData(func() interface{} { return attributeValue }, d, attributeName, diags)
	}
	return diags
}

func AddAttributeValueToResourceData(operation func() interface{}, d *schema.ResourceData, attributeName string, diags diag.Diagnostics) diag.Diagnostics {
	log.Printf("[Attribute: %s]: Entering AddAttributeValueToResourceData()", attributeName)
	defer log.Printf("[Attribute: %s]: Exiting AddAttributeValueToResourceData()", attributeName)
	log.Printf("[Attribute: %s]: Running function to get value", attributeName)
	attributeValue := operation()
	log.Printf("[Attribute: %s]: Adding value %v", attributeName, attributeValue)
	if err := d.Set(attributeName, attributeValue); err != nil {
		log.Printf("[Attribute: %s]: Setting value failed with error %v", attributeName, err)
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  err.Error(),
		})
	}
	return diags
}
