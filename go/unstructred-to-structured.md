```go
deployment, err := l.dynamicClient.Resource(gvr).Namespace("superadmin-dev-ns").Get(context.TODO(), inp.Name+"-kafka", metav1.GetOptions{})

deploymentData := deployment.UnstructuredContent()

status := deploymentData["status"].(map[string]interface{})
pods := status["pods"].(int64)
readyPods := status["readyPods"].(int64)
```


#### Unstractured with structured (for kubernetes go client)
```go

import (
	"k8s.io/apimachinery/pkg/runtime"
)

deploymentData := deployment.UnstructuredContent()
selector := deploymentData["spec"].(map[string]interface{})["selector"].(map[string]interface{})

var labelSelector metav1.LabelSelector
runtime.DefaultUnstructuredConverter.FromUnstructured(selector, &labelSelector)

fmt.Println(labelSelector)

```