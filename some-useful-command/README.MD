### mongodb connections string
k8s: 
```sh
mongodb://root:Password@mongodb-sts-0.eticker-mongodb.default.svc.cluster.local:27017/eticket?retryWrites=true&w=majority&authSource=admin
```

## mysql connection 

```sh
mysql -h http://172.17.18.176 -P 3306 -u nweudb -p
```
here 
    p = password (will ask pass after enter) and
    P = port


if mysql not installed:
```sh
sudo apt install mysql-server
```


## postgresql connection
```sh
psql -h 172.17.18.176 -U mah007 -d pgdb -p 5432
```
here -p = port , -d = detabase name

if psql is not installed:
```sh
sudo apt install postgresql-client
```


## mongo connection
installations:
```sh
#import public key
sudo apt-get install gnupg curl

curl -fsSL https://www.mongodb.org/static/pgp/server-7.0.asc | \
   sudo gpg -o /usr/share/keyrings/mongodb-server-7.0.gpg \
   --dearmor

#create the list file

#Ubuntu 22.04 jammy
echo "deb [ arch=amd64,arm64 signed-by=/usr/share/keyrings/mongodb-server-7.0.gpg ] https://repo.mongodb.org/apt/ubuntu jammy/mongodb-org/7.0 multiverse" | sudo tee /etc/apt/sources.list.d/mongodb-org-7.0.list

sudo apt-get update
sudo apt-get install -y mongodb-org
sudo systemctl start mongod
sudo systemctl daemon-reload
sudo systemctl status mongod
sudo systemctl enable mongod
#sudo systemctl stop mongod
#sudo systemctl restart mongod


mongosh "mongodb://mah007:123456789Ma@172.17.18.176:27017/testdb?retryWrites=true&w=majority&authSource=admin"
```

## Mysql push data script
some useful mysql command 
```sh
    show databases;
    show tables;
    select * from table_name;
    use table_name;
```
```sh

#!/bin/bash

# Variables
DB_HOST="167.172.5.224" # MySQL host (change if needed)
DB_PORT="3306"
DB_USER="ahsan"               # MySQL user (replace with your user)
DB_PASS="12345678Ma"      # MySQL password (replace with your password)
DB_NAME="test"         # Database name
TABLE_NAME="pushdata"         # Table name

# Number of rows to insert (You can adjust this number to control the size of data)
ROWS=100000

# Function to insert data using a loop
insert_data() {
    echo "Inserting $ROWS rows into $DB_NAME.$TABLE_NAME..."
    
    for ((i=1;i<=ROWS;i++)); do
        # Generate dummy data for the row
        VALUE1="Data_$i"
        VALUE2="This is a sample text for row $i"
        
        # Insert data into MySQL
        mysql -h "$DB_HOST" -P "${DB_PORT}" -u "$DB_USER" -p"$DB_PASS" -e "
        USE $DB_NAME;
        INSERT INTO $TABLE_NAME (column1, column2) VALUES ('$VALUE1', '$VALUE2');
        "
        
        # Print a message every 10,000 rows
        if (( i % 10000 == 0 )); then
            echo "Inserted $i rows..."
        fi
    done
    
    echo "Data insertion complete."
}

# Call the insert_data function
insert_data

```
