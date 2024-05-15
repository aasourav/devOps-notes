    sudo apt update
    sudo apt install nfs-kernel-server -y
    sudo mkdir -p /var/k8-nfs/data
    sudo chown -R nobody:nogroup /var/k8-nfs/data
    sudo chmod 777 /var/k8-nfs/data  #this directory will be act as nfs server

    cat <<EOF | sudo tee /etc/exports
    /var/k8-nfs/data *(rw,sync,no_subtree_check,no_root_squash,no_all_squash)
    EOF

    sudo exportfs -avr
    sudo systemctl restart nfs-kernel-server./
    sudo systemctl status nfs-kernel-server