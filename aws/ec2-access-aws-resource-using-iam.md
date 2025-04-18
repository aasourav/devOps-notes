Step 1: Deploy EC2
Step 2: Create `Role` 
    - Create role
    - Choose AWS service 
    - Use case (choose ec2) 
    - add policy (if you want to create your own then create policy first) 
    - Role name 
    - finally submit (create role)


Step 3: Attach this IAM role into ec2
    - Go go desired instance
    - Stop it first
    - Wait instance state become stop
    - go to `Actions` > Security > Modify IAM role
    - Choose the desired role and Click `Update IAM role`
    - Turn on ec2 and try aws cli on command list
    - Now you can access aws service by the help of IAM role without using aws configure(without providing token, secret)
