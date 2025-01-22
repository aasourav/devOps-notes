In this task the requirement was write a cronjob where the job is TURNCAT database table every 12 hours

Steps:
- Collect root username , password (requirement was loging as root)
- Collect host and port
- Create a table. the targeted table name is `automation_sessions`
```sh
CREATE TABLE automation_sessions (
    id SERIAL PRIMARY KEY,
    session_name VARCHAR(100) NOT NULL,
    user_id INT NOT NULL,
    start_time TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    end_time TIMESTAMP,
    status VARCHAR(50)
);
```
- Insert data on it 
```sh
INSERT INTO automation_sessions (session_name, user_id, end_time, status)
VALUES
    ('Session 1', 1, '2025-01-16 12:00:00', 'completed'),
    ('Session 2', 2, '2025-01-16 13:00:00', 'active'),
    ('Session 3', 3, '2025-01-16 14:00:00', 'completed');
```
- Verify data
```sh
SELECT * FROM automation_sessions;
```
- Write cronjob manifest:
```yaml
apiVersion: batch/v1
kind: CronJob
metadata:
  name: truncate-automation-sessions
  namespace: aas-ns
spec:
  schedule: "* * * * *" # every minutes for testing purpose
  jobTemplate:
    spec:
      template:
        spec:
          containers:
          - name: truncate-table
            image: postgres:15
            env:
              - name: POSTGRES_USER
                value: adfaes
              - name: POSTGRES_PASSWORD
                value: asdfedd
              - name: POSTGRES_DB
                value: aesdb
              - name: POSTGRES_HOST
                value: "172.19.255.201"
              - name: POSTGRES_PORT
                value: "3306"
            command: ["sh", "-c"]
            args:
              - |
                PGPASSWORD=$POSTGRES_PASSWORD psql -h $POSTGRES_HOST -p $POSTGRES_PORT -U $POSTGRES_USER -d $POSTGRES_DB -c "TRUNCATE TABLE automation_sessions;"
          restartPolicy: OnFailure

```