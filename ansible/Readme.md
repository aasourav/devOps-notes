tutorial link: https://www.youtube.com/watch?v=y2TSR7p3N0M&t
ansible: it automation tool 

## dry run : test without running
## start at task: name of a task the playbook should start at. it will skip all the task before specified task
ex: ansible-playbook playbook.yaml --start-at-task "task name"

## tag: tags (option) playbook or task and tag them
```yaml
- name: install httpd
  tags: install and start
  hosts: all
  tasks:
  - yum:
     name: httpd
     state: installed
     tags: install
```
ex: ansible-playbook playbook.yaml --tags "install"
ex: ansible-playbook playbook.yaml --skip-tags "install"

ansible has 3 main properties:
`1. Inventory : we know ansible can run multiple server at once. means from my computer I can setup (for ex. docker) in multiple computer. todo that we need server info (ip, username, password) . those infor provide by inventory. 
 2. Playbook : the file's is also called playbook. (learn more for details). we can write multiple module in single playbook.
 3. Module : small programme todo task (ex: crate file, start file etc)
 
 
 ssh without password:
 step 1: ssh-keygen 
 step 2: ssh-copy-id ip_remoteserver_whre_we_want_copy then provide password. done
 
 now if we ssh those ip then it will not ask password later
 