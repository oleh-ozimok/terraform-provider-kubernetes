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

func resourceKubernetesRole() *schema.Resource {
	return &schema.Resource{
		Create: resourceKubernetesRoleCreate,
		Read:   resourceKubernetesRoleRead,
		Exists: resourceKubernetesRoleExists,
		Update: resourceKubernetesRoleUpdate,
		Delete: resourceKubernetesRoleDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"metadata": namespacedMetadataSchema("role", false),
			"rule": {
				Type:        schema.TypeList,
				Description: "",
				Required:    true,
				MinItems:    1,
				Elem: &schema.Resource{
					Schema: rbacPolicyRuleSchema(),
				},
			},
		},
	}
}

func resourceKubernetesRoleCreate(d *schema.ResourceData, meta interface{}) error {
	conn := meta.(*kubernetes.Clientset)

	metadata := expandMetadata(d.Get("metadata").([]interface{}))
	role := &api.Role{
		ObjectMeta: metadata,
		Rules:      expandRBACPolicyRules(d.Get("rule").([]interface{})),
	}
	log.Printf("[INFO] Creating new Role: %#v", role)

	role, err := conn.RbacV1().Roles(metadata.Namespace).Create(role)
	if err != nil {
		return err
	}

	log.Printf("[INFO] Submitted new Role: %#v", role)

	d.SetId(buildId(metadata))

	return resourceKubernetesRoleRead(d, meta)
}

func resourceKubernetesRoleRead(d *schema.ResourceData, meta interface{}) error {
	conn := meta.(*kubernetes.Clientset)

	namespace, name, err := idParts(d.Id())
	if err != nil {
		return err
	}

	log.Printf("[INFO] Reading Role %s", name)
	role, err := conn.RbacV1().Roles(namespace).Get(name, meta_v1.GetOptions{})
	if err != nil {
		log.Printf("[DEBUG] Received error: %#v", err)
		return err
	}

	log.Printf("[INFO] Received ClusterRole: %#v", role)
	if err = d.Set("metadata", flattenMetadata(role.ObjectMeta)); err != nil {
		return err
	}

	flattenedRules := flattenRBACPolicyRules(role.Rules)
	log.Printf("[DEBUG] Flattened Role rules: %#v", flattenedRules)
	if err = d.Set("rule", flattenedRules); err != nil {
		return err
	}

	return nil
}

func resourceKubernetesRoleUpdate(d *schema.ResourceData, meta interface{}) error {
	conn := meta.(*kubernetes.Clientset)

	namespace, name, err := idParts(d.Id())
	if err != nil {
		return err
	}

	ops := patchMetadata("metadata.0.", "/metadata/", d)
	if d.HasChange("rule") {
		diffOps := patchPolicyRules(d)
		ops = append(ops, diffOps...)
	}
	data, err := ops.MarshalJSON()
	if err != nil {
		return fmt.Errorf("Failed to marshal update operations: %s", err)
	}
	log.Printf("[INFO] Updating Role %q: %v", name, string(data))
	out, err := conn.RbacV1().Roles(namespace).Patch(name, pkgApi.JSONPatchType, data)
	if err != nil {
		return fmt.Errorf("Failed to update Role: %s", err)
	}
	log.Printf("[INFO] Submitted updated Role: %#v", out)
	d.SetId(buildId(out.ObjectMeta))

	return resourceKubernetesRoleRead(d, meta)
}

func resourceKubernetesRoleDelete(d *schema.ResourceData, meta interface{}) error {
	conn := meta.(*kubernetes.Clientset)

	namespace, name, err := idParts(d.Id())
	if err != nil {
		return err
	}

	log.Printf("[INFO] Deleting Role: %#v", name)
	err = conn.RbacV1().Roles(namespace).Delete(name, &meta_v1.DeleteOptions{})
	if err != nil {
		return err
	}
	log.Printf("[INFO] Role %s deleted", name)

	d.SetId("")
	return nil
}

func resourceKubernetesRoleExists(d *schema.ResourceData, meta interface{}) (bool, error) {
	conn := meta.(*kubernetes.Clientset)

	namespace, name, err := idParts(d.Id())
	if err != nil {
		return false, err
	}

	log.Printf("[INFO] Checking Role %s", name)
	_, err = conn.RbacV1().Roles(namespace).Get(name, meta_v1.GetOptions{})
	if err != nil {
		if statusErr, ok := err.(*errors.StatusError); ok && statusErr.ErrStatus.Code == 404 {
			return false, nil
		}
		log.Printf("[DEBUG] Received error: %#v", err)
	}
	return true, err
}
