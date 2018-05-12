package kubernetes

import (
	"strconv"

	"github.com/hashicorp/terraform/helper/schema"
	api "k8s.io/api/rbac/v1"
)

func expandRBACAggregationRule(in []interface{}) *api.AggregationRule {
	if len(in) == 0 || in[0] == nil {
		return nil
	}

	rule := &api.AggregationRule{}

	m := in[0].(map[string]interface{})

	if selectors, ok := m["cluster_role_selector"].([]interface{}); ok {
		for _, selector := range selectors {
			rule.ClusterRoleSelectors = append(rule.ClusterRoleSelectors, *expandLabelSelector([]interface{}{selector}))
		}
	}

	return rule
}

func flattenRBACAggregationRule(in *api.AggregationRule) []interface{} {
	out := make([]interface{}, 1)

	var selectors []interface{}

	for _, v := range in.ClusterRoleSelectors {
		selectors = append(selectors, flattenLabelSelector(&v)[0])
	}

	out[0] = map[string]interface{}{
		"cluster_role_selector": selectors,
	}

	return out
}

func expandRBACPolicyRules(in []interface{}) []api.PolicyRule {
	if len(in) == 0 || in[0] == nil {
		return []api.PolicyRule{}
	}

	var rules []api.PolicyRule

	for i := range in {
		rule := api.PolicyRule{}
		m := in[i].(map[string]interface{})

		if v, ok := m["api_groups"].(*schema.Set); ok && v.Len() > 0 {
			rule.APIGroups = sliceOfString(v.List())
		}

		if v, ok := m["non_resource_urls"].(*schema.Set); ok && v.Len() > 0 {
			rule.NonResourceURLs = sliceOfString(v.List())
		}

		if v, ok := m["resource_names"].(*schema.Set); ok && v.Len() > 0 {
			rule.ResourceNames = sliceOfString(v.List())
		}

		if v, ok := m["resources"].(*schema.Set); ok && v.Len() > 0 {
			rule.Resources = sliceOfString(v.List())
		}

		if v, ok := m["verbs"].(*schema.Set); ok && v.Len() > 0 {
			rule.Verbs = sliceOfString(v.List())
		}

		rules = append(rules, rule)
	}

	return rules
}

func flattenRBACPolicyRules(in []api.PolicyRule) []interface{} {
	att := make([]interface{}, len(in), len(in))
	for i, n := range in {
		m := make(map[string]interface{})

		if len(n.APIGroups) > 0 {
			m["api_groups"] = newStringSet(schema.HashString, n.APIGroups)
		}

		if len(n.NonResourceURLs) > 0 {
			m["non_resource_urls"] = newStringSet(schema.HashString, n.NonResourceURLs)
		}

		if len(n.ResourceNames) > 0 {
			m["resource_names"] = newStringSet(schema.HashString, n.ResourceNames)
		}

		if len(n.Resources) > 0 {
			m["resources"] = newStringSet(schema.HashString, n.Resources)
		}

		if len(n.Verbs) > 0 {
			m["verbs"] = newStringSet(schema.HashString, n.Verbs)
		}

		att[i] = m
	}
	return att
}

func patchPolicyRules(d *schema.ResourceData) PatchOperations {
	ops := make([]PatchOperation, 0, 0)

	rules := expandRBACPolicyRules(d.Get("rule").([]interface{}))
	for i, v := range rules {
		ops = append(ops, &ReplaceOperation{
			Path:  "/rules/" + strconv.Itoa(i),
			Value: v,
		})
	}

	return ops
}

func expandRBACRoleRef(in []interface{}) api.RoleRef {
	if len(in) == 0 || in[0] == nil {
		return api.RoleRef{}
	}

	ref := api.RoleRef{}
	m := in[0].(map[string]interface{})

	if v, ok := m["api_group"]; ok {
		ref.APIGroup = v.(string)
	}
	if v, ok := m["kind"].(string); ok {
		ref.Kind = string(v)
	}
	if v, ok := m["name"]; ok {
		ref.Name = v.(string)
	}

	return ref
}

func expandRBACSubjects(in []interface{}) []api.Subject {
	if len(in) == 0 || in[0] == nil {
		return []api.Subject{}
	}
	var subjects []api.Subject
	for i := range in {
		subject := api.Subject{}
		m := in[i].(map[string]interface{})
		if v, ok := m["api_group"]; ok {
			subject.APIGroup = v.(string)
		}
		if v, ok := m["kind"].(string); ok {
			subject.Kind = string(v)
		}
		if v, ok := m["name"]; ok {
			subject.Name = v.(string)
		}
		if v, ok := m["namespace"]; ok {
			subject.Namespace = v.(string)
		}
		subjects = append(subjects, subject)
	}
	return subjects
}

func flattenRBACRoleRef(in api.RoleRef) []interface{} {
	att := make(map[string]interface{}, 0)

	if in.APIGroup != "" {
		att["api_group"] = in.APIGroup
	}
	att["kind"] = in.Kind
	att["name"] = in.Name
	return []interface{}{att}
}

func flattenRBACSubjects(in []api.Subject) []interface{} {
	att := make([]interface{}, len(in), len(in))
	for i, n := range in {
		m := make(map[string]interface{})
		if n.APIGroup != "" {
			m["api_group"] = n.APIGroup
		}
		m["kind"] = n.Kind
		m["name"] = n.Name
		if n.Namespace != "" {
			m["namespace"] = n.Namespace
		}

		att[i] = m
	}
	return att
}

func patchRbacSubject(d *schema.ResourceData) PatchOperations {
	ops := make([]PatchOperation, 0, 0)

	subjects := expandRBACSubjects(d.Get("subject").([]interface{}))
	for i, v := range subjects {
		ops = append(ops, &ReplaceOperation{
			Path:  "/subjects/" + strconv.Itoa(i),
			Value: v,
		})
	}

	return ops
}
