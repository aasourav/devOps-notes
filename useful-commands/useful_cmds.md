    //aes
    make generate manifest  -> aesl api-resource build
    sudo hostname new_host_name -> change hostname

    //kubernetes
    kubens -> switch namespace, it's also inside in kubectx
    kubectx -> tools to switch contexts(cluster) on kubectl faster
    kubectl port-forward service/ingress-nginx-controller 8080:80 -n core-ns --address 0.0.0.0
    mongosh "mongodb://ratul:1234@172.18.255.202:3342/admin?retryWrites=true&w=majority"
    mongodb://databaseAdmin:databaseAdmin123456@host:port/admin?retryWrites=true&w=majority // this is for docker


    //linux

    sudo vim /etc/hostname -> type your new hostname
    sudo vim /etc/hosts -> add hosts with ip address so that you don't need type ip address for ssh. then you can log in just ssh hostname
    // authorised keys, ssh public,private key (rsa ) stays in ~/.ssh dicrectory

    //permit ssh linux as root -> (ubuntu 22 tested)

    step 1: 
        `sudo vim /etc/ssh/sshd_config`
    step 2: 
        FROM:
            #PermitRootLogin prohibit-password
        TO:
        PermitRootLogin yes

        shortcut way: `sudo sed -i 's/#PermitRootLogin prohibit-password/PermitRootLogin yes/' /etc/ssh/sshd_config`

    step 3:
        `sudo systemctl restart ssh`
    
    step 4:
        `sudo passwd`
        [sudo] password for linuxconfig: 
        Enter new UNIX password: 
        Retype new UNIX password: 
        passwd: password updated successfully
    
    step 5:
        `sudo ufw allow ssh`