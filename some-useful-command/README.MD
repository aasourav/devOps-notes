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
