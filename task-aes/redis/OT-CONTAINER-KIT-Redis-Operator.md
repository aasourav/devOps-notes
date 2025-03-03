Update `tcp-keep-alive` field:
ref: https://ot-redis-operator.netlify.app/docs/configuration/redis/

![alt text](image.png)

update `timeout` field:

this has not dedicated field. we have to configure it using `additionalRedisConfig` field:

```yaml
apiVersion: redis.redis.opstreelabs.in/v1beta2
kind: Redis
metadata:
  name: redis-standalone
spec:
  replicas: 1
  redisConfig:
    additionalRedisConfig: |
      timeout 300
  kubernetesConfig:
    image: "redis:6.0.9"
    imagePullPolicy: IfNotPresent
    resources:
      requests:
.....

```