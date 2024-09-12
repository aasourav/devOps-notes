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

