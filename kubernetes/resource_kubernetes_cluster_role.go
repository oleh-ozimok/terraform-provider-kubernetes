package kubernetes

import (
	"fmt"
	"log"

	"github.com/hashicorp/terraform/helper/schema"
	api "k8s.io/api/rbac/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	pkgApi "k8s.io/apimachinery/pkg/types"
	"k8s.io/client-go/kubernetes"
)

func resourceKubernetesClusterRole() *schema.Resource {
	return &schema.Resource{
		Create: resourceKubernetesClusterRoleCreate,
		Read:   resourceKubernetesClusterRoleRead,
		Exists: resourceKubernetesClusterRoleExists,
		Update: resourceKubernetesClusterRoleUpdate,
		Delete: resourceKubernetesClusterRoleDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"metadata": metadataSchema("cluster role", false),
			"aggregation_rule": {
				Type:        schema.TypeList,
				Description: "",
				MaxItems:    1,
				Optional:    true,
				Elem: &schema.Resource{
					Schema: rbacAggregationRuleSchema(),
				},
			},
			"rule": {
				Type:          schema.TypeList,
				Description:   "",
				Optional:      true,
				ConflictsWith: []string{"aggregation_rule"},
				MinItems:      1,
				Elem: &schema.Resource{
					Schema: rbacPolicyRuleSchema(),
				},
			},
		},
	}
}

func resourceKubernetesClusterRoleCreate(d *schema.ResourceData, meta interface{}) error {
	conn := meta.(*kubernetes.Clientset)

	metadata := expandMetadata(d.Get("metadata").([]interface{}))
	clusterRole := &api.ClusterRole{
		ObjectMeta: metadata,
	}

	aggregationRule := expandRBACAggregationRule(d.Get("aggregation_rule").([]interface{}))

	if aggregationRule != nil {
		clusterRole.AggregationRule = aggregationRule
	} else {
		clusterRole.Rules = expandRBACPolicyRules(d.Get("rule").([]interface{}))
	}

	log.Printf("[INFO] Creating new ClusterRole: %#v", clusterRole)

	clusterRole, err := conn.RbacV1().ClusterRoles().Create(clusterRole)
	if err != nil {
		return err
	}

	log.Printf("[INFO] Submitted new ClusterRole: %#v", clusterRole)

	d.SetId(metadata.Name)

	return resourceKubernetesClusterRoleRead(d, meta)
}

func resourceKubernetesClusterRoleRead(d *schema.ResourceData, meta interface{}) error {
	conn := meta.(*kubernetes.Clientset)

	name := d.Id()
	log.Printf("[INFO] Reading ClusterRole %s", name)
	clusterRole, err := conn.RbacV1().ClusterRoles().Get(name, meta_v1.GetOptions{})
	if err != nil {
		log.Printf("[DEBUG] Received error: %#v", err)
		return err
	}

	log.Printf("[INFO] Received ClusterRole: %#v", clusterRole)
	if err = d.Set("metadata", flattenMetadata(clusterRole.ObjectMeta)); err != nil {
		return err
	}

	if clusterRole.AggregationRule != nil {
		flattenedAggregationRule := flattenRBACAggregationRule(clusterRole.AggregationRule)
		log.Printf("[DEBUG] Flattened ClusterRole aggregation_rule: %#v", flattenedAggregationRule)
		if err = d.Set("aggregation_rule", flattenedAggregationRule); err != nil {
			return err
		}
	} else {
		flattenedRules := flattenRBACPolicyRules(clusterRole.Rules)
		log.Printf("[DEBUG] Flattened ClusterRole rules: %#v", flattenedRules)
		if err = d.Set("rule", flattenedRules); err != nil {
			return err
		}
	}

	return nil
}

func resourceKubernetesClusterRoleUpdate(d *schema.ResourceData, meta interface{}) error {
	conn := meta.(*kubernetes.Clientset)

	name := d.Id()

	ops := patchMetadata("metadata.0.", "/metadata/", d)

	if d.HasChange("rule") {
		diffOps := patchPolicyRules(d)
		ops = append(ops, diffOps...)
	}

	if d.HasChange("aggregation_rule") {
		ops = append(ops, &ReplaceOperation{
			Path:  "/aggregationRule",
			Value: expandRBACAggregationRule(d.Get("aggregation_rule").([]interface{})),
		})
	}

	data, err := ops.MarshalJSON()
	if err != nil {
		return fmt.Errorf("Failed to marshal update operations: %s", err)
	}

	log.Printf("[INFO] Updating ClusterRole %q: %v", name, string(data))

	out, err := conn.RbacV1().ClusterRoles().Patch(name, pkgApi.JSONPatchType, data)
	if err != nil {
		return fmt.Errorf("Failed to update ClusterRole: %s", err)
	}

	log.Printf("[INFO] Submitted updated ClusterRole: %#v", out)

	d.SetId(out.ObjectMeta.Name)

	return resourceKubernetesClusterRoleRead(d, meta)
}

func resourceKubernetesClusterRoleDelete(d *schema.ResourceData, meta interface{}) error {
	conn := meta.(*kubernetes.Clientset)

	name := d.Id()
	log.Printf("[INFO] Deleting ClusterRole: %#v", name)
	err := conn.RbacV1().ClusterRoles().Delete(name, &meta_v1.DeleteOptions{})
	if err != nil {
		return err
	}
	log.Printf("[INFO] ClusterRole %s deleted", name)

	d.SetId("")
	return nil
}

func resourceKubernetesClusterRoleExists(d *schema.ResourceData, meta interface{}) (bool, error) {
	conn := meta.(*kubernetes.Clientset)

	name := d.Id()
	log.Printf("[INFO] Checking ClusterRole %s", name)
	_, err := conn.RbacV1().ClusterRoles().Get(name, meta_v1.GetOptions{})
	if err != nil {
		if statusErr, ok := err.(*errors.StatusError); ok && statusErr.ErrStatus.Code == 404 {
			return false, nil
		}
		log.Printf("[DEBUG] Received error: %#v", err)
	}
	return true, err
}
