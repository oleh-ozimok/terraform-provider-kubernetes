package kubernetes

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func rbacAggregationRuleSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"cluster_role_selector": {
			Type:        schema.TypeList,
			Description: "",
			Optional:    true,
			MinItems:    1,
			Elem: &schema.Resource{
				Schema: labelSelectorFields(),
			},
		},
	}
}

func rbacPolicyRuleSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"api_groups": {
			Type:        schema.TypeSet,
			Elem:        &schema.Schema{Type: schema.TypeString},
			Set:         schema.HashString,
			Description: "",
			Optional:    true,
		},
		"resources": {
			Type:        schema.TypeSet,
			Elem:        &schema.Schema{Type: schema.TypeString},
			Set:         schema.HashString,
			Description: "",
			Optional:    true,
		},
		"resource_names": {
			Type:        schema.TypeSet,
			Elem:        &schema.Schema{Type: schema.TypeString},
			Set:         schema.HashString,
			Description: "",
			Optional:    true,
		},
		"verbs": {
			Type:        schema.TypeSet,
			Elem:        &schema.Schema{Type: schema.TypeString},
			Set:         schema.HashString,
			Description: "",
			Optional:    true,
		},
		"non_resource_urls": {
			Type:        schema.TypeSet,
			Elem:        &schema.Schema{Type: schema.TypeString},
			Set:         schema.HashString,
			Description: "",
			Optional:    true,
		},
	}
}

func rbacRoleRefSchema(kind string) map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"api_group": {
			Type:        schema.TypeString,
			ForceNew:    true,
			Optional:    true,
			Description: "The API group of the user. Always `rbac.authorization.k8s.io`",
			Default:     "rbac.authorization.k8s.io",
		},
		"kind": {
			Type:        schema.TypeString,
			ForceNew:    true,
			Description: "The kind of resource.",
			Default:     kind,
			Optional:    true,
		},
		"name": {
			Type:        schema.TypeString,
			ForceNew:    true,
			Description: "The name of the User to bind to.",
			Required:    true,
		},
	}
}

func rbacSubjectSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"api_group": {
			Type:        schema.TypeString,
			Description: "The API group of the user. Always `rbac.authorization.k8s.io`",
			Optional:    true,
			Default:     "",
		},
		"kind": {
			Type:        schema.TypeString,
			Description: "The kind of resource.",
			Required:    true,
		},
		"name": {
			Type:        schema.TypeString,
			Description: "The name of the resource to bind to.",
			Required:    true,
		},
		"namespace": {
			Type:        schema.TypeString,
			Description: "The Namespace of the ServiceAccount",
			Optional:    true,
			Default:     "default",
		},
	}
}
