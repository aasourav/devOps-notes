
How Docker Works
```md
Human -> Docker client -> REST -> Docker Deamon -> ContainerD -> SystemD 
         =====================================     ==========    ======== 
                        Docker Engine               Runtime       Kernel
Note: -> means interact 
```

### Component of Docker
- Docker Image   ( RootFS + UserFS + AppFS)
- Docker Registry (A place where docker image stored)
- Docker Container ( A moment where you run image only copy of flatten image get One MNT (mount to) User Space with NET and PMAP)

Some tools: https://www.devopsschool.com/blog/list-of-top-container-runtime-interface-projects/
- Podman



### Docker Journey

Human  -> Docker run imagename

Human Interact with <b>Docker Client</b> Client sent to the <b>Daemon</b> , Deamon will check Local Repo If Not foud then It will go to the Remote repository , It Download the Image and keep it to the Local Repo. When you get the Image you get the <b>Container</b>  container is managed/created by ContainerD. ContainerD will talk to the <b>Kernel</b> 



### Lifecycle of VMS
Create -> Start -> Stop -> Start -> Restart -> Pause -> Unpause -> Kill -> Remove

What is Pause vs Unpause , Stop vs Kill

```sh
ps -eaf | grep docker
```


To run Docker Need Root or a Usern need to be a part of Group called "docker"

### Lifecycle of Container
- Create ( docker create )
- Start ( docker start )
- Stop  ( docker stop )
- Restart ( docker restart )
- Pause ( docker pause )
- Unpause ( docker unpause )
- Kill ( docker kill )
- Remove ( docker remove )







