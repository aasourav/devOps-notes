setup minio client (mc)

```sh
    curl -o mc https://dl.min.io/client/mc/release/linux-amd64/mc
    chmod +x mc
    sudo mv mc /usr/local/bin/
    mc alias set myminio http://minio.minio-dev.svc.cluster.local:9000 ACCESS_KEY SECRET_KEY  #add remote minio server
    mc mb myminio/test // create bucket
    mc put aliasname-see-line-7/bucket_name // put file on bucket
```

