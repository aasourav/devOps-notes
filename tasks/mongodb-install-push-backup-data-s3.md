Step 1: Create IAM role for ec2 instance to access s3 resource.
    IAM -> Roles -> Create role -> Select "AWS service" -> Choose usecase 'EC2' -> Click 'Next' -> Add Permission (S3fullaccess) -> Click 'Next' -> Role Name -> Click 'Create role'

Step 2: install mongodb
```sh
wget -qO- https://gist.githubusercontent.com/aasourav/a55f099bd809db22e8214a014e87eddd/raw/mongo.sh | bash
#or
curl -s https://gist.githubusercontent.com/aasourav/a55f099bd809db22e8214a014e87eddd/raw/mongo.sh | bash
```

Step 3: configure mongodb.conf (update 'net' section)
```sh
    sudo vim /etc/mongo.conf 
    # update bind ip 127.0.0.1 to 0.0.0.0
    # I will add some productoin grade configureation later
    # mongosh connection
    mongosh --host 3.110.219.44 --port 27772 -u admin -p --authenticationDatabase admin

    use admin
    db.grantRolesToUser("admin", [{ role: "root", db: "admin" }]) # provide all access
```

Or use docker compse:
```yaml
services:
  mongo:
    image: mongo:6.0.21-jammy # Use the desired MongoDB version
    container_name: mongo
    restart: always
    environment:
      MONGO_INITDB_ROOT_USERNAME: root # Set MongoDB root user
      MONGO_INITDB_ROOT_PASSWORD: menToripdB # Set MongoDB root password
      MONGO_INITDB_DATABASE: myapp # Set the initial database (optional)
    volumes:
      - /home/ubuntu/mongo-data-backup/:/data/db # Persist data to a named volume
      #- ./mongo/init:/docker-entrypoint-initdb.d # Custom initialization scripts
    ports:
      - "27772:27017" # Expose MongoDB port
    networks:
      - mongo_network
    healthcheck:
      test: ["CMD", "mongosh", "--eval", "db.runCommand('ping')"]
      interval: 10s
      retries: 5
      timeout: 5s
      start_period: 10s
    logging: #if you wanna disable logs
      driver: none
networks:
  mongo_network:
    driver: bridge
```
`docker compose -f docker-compose.yaml up -d`

Step 3.5: install aws cli
```sh
sudo apt install unzip
curl "https://awscli.amazonaws.com/awscli-exe-linux-x86_64.zip" -o "awscliv2.zip"
unzip awscliv2.zip
sudo ./aws/install

#or
curl -s https://gist.githubusercontent.com/aasourav/46312b3b08efacb3fa214888d0584caa/raw/mongo.sh | bash
```

Step 4: Create backup script ( before doing this first create a s3 bucket) . give this executable permission
```sh
chmod +x filename.sh
```
```sh
#!/bin/bash

# Variables
TIMESTAMP=$(date +"%Y-%m-%d_%H-%M-%S")
BUCKET_NAME="mongo-db-backups-yourproject"
BACKUP_DIR="/tmp/mongobackup"
ARCHIVE_NAME="mongo_backup_$TIMESTAMP.tar.gz"

# Create backup
mkdir -p $BACKUP_DIR
mongodump --out $BACKUP_DIR/dump_$TIMESTAMP

mongodump --host 3.110.219.44 --port 27772 \
  --username admin \
  --password "mentsor!pAdmin" \
  --authenticationDatabase admin \
  --db dbName \
  --out $BACKUP_DIR/dump_$TIMESTAMP
# 'mongodump --uri="mongodb+srv://hrm:BpcQUaseeFLFYcadfIIvaUNu@cluster0.pisbwzh.mongodb.net/mentor_ip?retryWrites=true&w=majority" --out=./dump'

# Compress it
tar -czf /tmp/$ARCHIVE_NAME -C $BACKUP_DIR .

# Upload to S3
aws s3 cp /tmp/$ARCHIVE_NAME s3://$BUCKET_NAME/$ARCHIVE_NAME

# Cleanup
rm -rf $BACKUP_DIR
rm -f /tmp/$ARCHIVE_NAME
```

Step 5: Create cronjob

```sh
crontab -e
# then update the file
* * * * * /usr/local/bin/mongo_backup.sh >> /home/ubuntu/mongo_backup.log #output log file maybe create first by manually
# or if  you dont wanna log
* * * * * /home/ubuntu/mong-back.sh >/dev/null 2>&1
crontab -l # see cronjobs
grep CRON /var/log/syslog # see cronjobs executions

```


Step 6: Restore mongodb
```sh
mongorestore --uri="mongodb://root:menToripdB@3.110.219.44:27772/myapp?authSource=admin" --drop /home/aes-sourav/Downloads/mongo_backup_2025-04-08_06-25-16/myapp # you have to tell where your prelude.json is located
```
