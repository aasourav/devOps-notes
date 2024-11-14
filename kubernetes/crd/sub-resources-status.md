# Understanding the `status` Subresource in Kubernetes CRDs

The `status` subresource in Kubernetes Custom Resource Definitions (CRDs) can be a bit confusing at first, but it's an important concept when designing custom resources. Let's break down what it does and why it's useful.

## What is the `status` Subresource?

When you define a Kubernetes custom resource (CR), you often want to differentiate between **desired state** and **current state**.

- **Spec**: Represents the desired state of the resource.
- **Status**: Represents the actual state of the resource.

The **`/status` subresource** is a special endpoint for updating only the `status` field of your custom resource. When you enable this subresource, updates to the `status` field are handled separately from updates to the rest of the custom resource (`spec`, `metadata`, etc.). This ensures that the `status` field can only be modified through the `/status` endpoint, which helps maintain a clear distinction between user-driven configuration (`spec`) and system-driven status reporting (`status`).

## Why Use the `status` Subresource?

1. **Separation of Concerns**:
   - Users or automation tools can update the `spec` (desired state) without worrying about accidentally modifying the `status`.
   - Controllers (like a Kubernetes operator) can update the `status` field to reflect the current state of the resource without affecting the `spec`.

2. **Enhanced Security**:
   - By enabling the `status` subresource, Kubernetes RBAC (Role-Based Access Control) can differentiate between permissions for updating the `spec` and the `status`. For example, you could allow a controller to update the `status` while restricting users to only update the `spec`.

3. **Reduced Race Conditions**:
   - Updates to the `spec` and `status` are done independently, which can help prevent race conditions between user-driven changes and controller updates.

## Example: Enabling the `status` Subresource

Here's your example with the `+kubebuilder:subresource:status` marker:

```go
// +kubebuilder:subresource:status
type Toy struct {
    metav1.TypeMeta   `json:",inline"`
    metav1.ObjectMeta `json:"metadata,omitempty"`

    Spec   ToySpec   `json:"spec,omitempty"`
    Status ToyStatus `json:"status,omitempty"`
}

### How It Works in Practice
When you enable the /status subresource, Kubernetes generates two separate endpoints for your CR:

Main resource endpoint:
/apis/<group>/<version>/namespaces/<namespace>/toys/<name>

You can use this to update the spec and other fields, excluding the status.
Status subresource endpoint:
/apis/<group>/<version>/namespaces/<namespace>/toys/<name>/status

You can use this to update only the status field.
Example Usage with kubectl


Updating the spec:
```bash
kubectl patch toy <toy-name> --type=merge -p '{"spec":{"color":"blue"}}'
```

Updating the status:

```bash
kubectl patch toy <toy-name> --type=merge -p '{"status":{"phase":"Running"}}' --subresource=status
```

### Controller Usage
If you are writing a Kubernetes controller or operator using the controller-runtime library, you typically use the client.Status().Update() method to update the status field:

```go
toy := &Toy{}
if err := r.Get(ctx, req.NamespacedName, toy); err != nil {
    return ctrl.Result{}, err
}

// Update the status
toy.Status.Phase = "Running"
if err := r.Status().Update(ctx, toy); err != nil {
    return ctrl.Result{}, err
}

```
Summary
The status subresource is a separate endpoint dedicated to updating the status field of your custom resource.
It helps separate the user-defined configuration (spec) from the system-reported state (status).
It enables fine-grained RBAC control and reduces the risk of race conditions.
To enable it, use the +kubebuilder:subresource:status marker above your CRD definition.
This feature is particularly useful when building operators or controllers that need to update the status of resources based on observed conditions.