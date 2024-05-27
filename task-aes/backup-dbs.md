```sh
mongo backdup :
	"/usr/bin/mongodump -u $USER -p $PASSWORD -o /tmp/backup -h " + dbpod.Name + "-replica && tar cvzf mongodb-backup-$(date +\\%Y\\%m\\%d_\\%H).tar.gz /tmp/backup  && apt update && apt install curl && curl -o mc https://dl.min.io/client/mc/release/linux-amd64/mc && chmod +x mc && mv mc /usr/local/bin/ && mc alias set myminio http://minio.minio-dev.svc.cluster.local:9000 minioadmin minioadmin && mc put mongodb-backup-$(date +\\%Y\\%m\\%d_\\%H).tar.gz myminio/" + dbpod.Name + "-" + dbpod.Namespace + "-" + "mongodb",
```
```sh	
mysql backup : 
	"mysqldump --triggers --routines --events --no-tablespaces -u$MYSQL_USERNAME -p$MYSQL_PASSWORD -h" + dbpod.Name + "-mysql-master." + dbpod.Namespace + ".svc.cluster.local -P3306 --default-character-set=utf8 --databases " + dbpod.Spec.DbName + " > /tmp/$(date +\\%Y\\%m\\%d_\\%H).sql && curl -o mc https://dl.min.io/client/mc/release/linux-amd64/mc && chmod +x mc && mv mc /usr/local/bin/ && mc alias set myminio http://minio.minio-dev.svc.cluster.local:9000 minioadmin minioadmin && mc put /tmp/$(date +\\%Y\\%m\\%d_\\%H).sql myminio/" + dbpod.Name + "-" + dbpod.Namespace + "-" + "mysql",
```	
```sh
postgresql backup :
	"echo $PGPASS > /root/.pgpass && chmod 600 /root/.pgpass && pg_dump -U " + userName + " -h " + dbpod.Name + "-cluster " + dbpod.Spec.DbName + " > /var/backups/backup-$(date +\\%Y\\%m\\%d_\\%H).sql  && curl -o mc https://dl.min.io/client/mc/release/linux-amd64/mc && chmod +x mc && mv mc /usr/local/bin/ && mc alias set myminio http://minio.minio-dev.svc.cluster.local:9000 puNPwHUvuNbzap8JWW7u llSVWiTyBK3Cv7p4R6QJvbYDZ44HrW78237iZ5bN && mc put /var/backups/backup-$(date +\\%Y\\%m\\%d_\\%H).sql myminio/postgress",
```