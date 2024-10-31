```go
deployment, err := l.dynamicClient.Resource(gvr).Namespace("superadmin-dev-ns").Get(context.TODO(), inp.Name+"-kafka", metav1.GetOptions{})

deploymentData := deployment.UnstructuredContent()

status := deploymentData["status"].(map[string]interface{})
pods := status["pods"].(int64)
readyPods := status["readyPods"].(int64)
```