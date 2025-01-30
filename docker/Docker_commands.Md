## Container logs
to see the docker container logs proper way 

# docker exec -it container_id/name /bin/bash


```sh
cd /var/lib/docker/containers # To see Container list, (Host machine where container is running)
# we can see directory list name with container Id.
# inside container id directory we will see container_id.log . this is the log file where it stored
# cat container_id.log | jq (it will show the logs in json format that will be more readable)
```

# Docker Commands

```sh
docker info
```