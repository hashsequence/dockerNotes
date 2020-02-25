
# Docker Notes

## 19 Checking our Docker install and config

* Nginx (web server) container
    * makes it easy to learn docker stuff


gives us info on our docker:
```console
$ docker info
```

this is the new way to run docker container (old way is $docker run)
```console
$ docker container run
```

### image vs Container

* an image is the app you want to run
* container is an instance of the image running as a process docker hub

* docker container run --publish 80:80 nginx
    * downloads nginx from docker hub
    * start a new container of that image
    * opens the local host

```console
$ sudo docker container run --publish 80:80 --detach nginx
$ sudo docker container ls
$ sudo docker container ls
CONTAINER ID        IMAGE               COMMAND                  CREATED              STATUS              PORTS                NAMES
ad1df56e92df        nginx               "nginx -g 'daemon of…"   About a minute ago   Up About a minute   0.0.0.0:80->80/tcp   suspicious_nightingale
$ sudo docker container ls -a
CONTAINER ID        IMAGE               COMMAND                  CREATED              STATUS                     PORTS                NAMES
22f59b3f689b        nginx               "nginx -g 'daemon of…"   About a minute ago   Created                                         webhost
ad1df56e92df        nginx               "nginx -g 'daemon of…"   2 minutes ago        Up 2 minutes               0.0.0.0:80->80/tcp   suspicious_nightingale
6369fb6bebf2        nginx               "nginx -g 'daemon of…"   3 minutes ago        Exited (0) 2 minutes ago                        determined_elion
4eab28f4ac02        nginx               "nginx -g 'daemon of…"   5 minutes ago        Exited (0) 3 minutes ago                        wonderful_varahamihira
3cf841f5841d        nginx               "nginx -g 'daemon of…"   9 minutes ago        Exited (0) 5 minutes ago                        brave_hugle
d1e085e5ccb9        hello-world         "/hello"                 18 hours ago         Exited (0) 18 hours ago                         suspicious_roentgen
$ docker container stop ad1
```

* if you want to stop the docker use docker container ls to see you process id
* then run the stop command with the processid (only need the first few digits for it to be unique)


```console
$ sudo docker container run --publish 80:80 --detach --name webhost2 nginx
950d1cf9403ba3eb440b29c7bc636f28ede2c356214259169e5faa2fd28d6226
$ sudo docker container logs webhost2
172.17.0.1 - - [01/Apr/2019:00:57:48 +0000] "HEAD / HTTP/1.1" 200 0 "-" "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/73.0.3683.75 Safari/537.36" "-"
172.17.0.1 - - [01/Apr/2019:00:57:48 +0000] "HEAD / HTTP/1.1" 200 0 "-" "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/73.0.3683.75 Safari/537.36" "-"
```


Display the running processes of a container:

```console
$ sudo docker container top
"docker container top" requires at least 1 argument.
See 'docker container top --help'.

Usage:  docker container top CONTAINER [ps OPTIONS]

```

### this is to show running docker containers

removing the docker container webhost2


```console 
$ sudo docker container ls
CONTAINER ID        IMAGE               COMMAND                  CREATED              STATUS              PORTS                NAMES
950d1cf9403b        nginx               "nginx -g 'daemon of…"   About a minute ago   Up About a minute   0.0.0.0:80->80/tcp   webhost2
$ sudo docker container rm -f 950
$ sudo docker container ls
CONTAINER ID        IMAGE               COMMAND             CREATED             STATUS     
```

## 21 Debrief: What Happens When We Run a Container

1. looks for the image locally in image cache
2. looks in remote image repository (dockerhub)
3. downloads latest version
4. creates a new container based on that image and prepares to start
5. gives it a vertual ip on a private network inside the docker engine
6. opens up port 80 on host and forwards to port 80 in containers
7. start container by using the CMD in the image Dockerfile

take this command:

```console
$ docker container run --publish 8080:80 --name webhost -d nginx:1.11 nginx -T
```
* 1.11 --> change version of IMAGE
* -T --> change CMD run on start
* 8080:80 --> change host listening port
    *  example difference between 8080:80 and 80:8080 :
    ```console
     docker run --restart always --name myjenkins -p 8080:8080
     ```
     8080:80 refers that in the container you are using port 80 and you are forwarding that port to host machine's 8080 port. So you are running Jenkins on port 80 inside your container wherever in scenario 2 you are running Jenkins on port 8080 inside the container and exposing it over the same port on host machine. For example if I am running mysql in container I may use 8080:3306 so mysql would be running on port 3306 but exposed on 8080 of host machine but if choose it to be 8080:80 for mysql it may not work because as per the code of mysql it binds itself on port 3306 not port 80. Same is the scenario in your case of Jenkins too.
     * When you say 8080:80, it means any request coming on port 8080 will be forwarded to service running on port 80 inside your docker container. Similarly 8080:8080 means any request coming for port 8080 will be forwarded to service running on port 8080 inside your container. tldr: Port for Outside World: Actual Port of service in container
     * https://stackoverflow.com/questions/52173352/what-is-the-difference-between-publishing-808080-and-80808080-in-a-docker-run
     



## 22 container vs. VM: Its the process

```console
$ sudo docker run --name mongo -d mongo
$ sudo docker top mongo
UID                 PID                 PPID                C                   STIME               TTY                 TIME                CMD
999                 31821               31800               1                   18:07               ?                   00:00:01            mongod --bind_ip_all
ps
  PID TTY          TIME CMD
26896 pts/4    00:00:00 console
32007 pts/4    00:00:00 ps
$ sudo docker stop mongo
sudo docker ps
CONTAINER ID        IMAGE               COMMAND             CREATED             STATUS     
$ ps aux | grep mongo
avwong13 32138  0.0  0.0  14224   968 pts/4    S+   18:11   0:00 grep --color=auto mongo
```

* this command looks for process with mongo
* lets now below start mongo again and use ps aux

```console
$ sudo docker ps
CONTAINER ID        IMAGE               COMMAND                  CREATED             STATUS              PORTS               NAMES
1733fc79806e        mongo               "docker-entrypoint.s…"   5 minutes ago       Up 3 seconds        27017/tcp           mongo
$ ps aux | grep mongo
999      32200 13.3  0.9 1095764 75832 ?       Ssl  18:12   0:01 mongod --bind_ip_all
avwong13 32308  0.0  0.0  14224   960 pts/4    S+   18:12   0:00 grep --color=auto mongo
```

## 24 Managing Multiple Containers

### assignment 

* docs.docker.com and --help
* run a nginx, mysql, httpd
* run all of them --detach(or -d), name them with --name
* nginx should listen on 80:80, httpd on 8080:80, mysql on 3306:3306
* when running mysql, use the --env(or -e) to pass in MYSQL_RANDOM_ROOT_PASSWORD=yes
* use docker container logs on mysql to find the random MYSQL_RANDOM_ROOT_PASSWORD
* clean it all up with docker container stop and docker container rm   (both can acceptmultiple names or ID's)
* use docker container ls to check


### solution

```console
$ sudo docker container run --publish 80:80 --detach --name c1 nginx
$ sudo docker container run --publish 8080:80 --detach --name c2 httpd
$ sudo docker container run -d -p 3306:3306 --name c3 -e MYSQL_RANDOM_ROOT_PASSWORD=yes mysql
$ sudo docker container logs c3 | grep PASSWORD

$ sudo docker container ls
CONTAINER ID        IMAGE               COMMAND                  CREATED             STATUS              PORTS                               NAMES
b11df1035d6d        mysql               "docker-entrypoint.s…"   4 minutes ago       Up 3 minutes        0.0.0.0:3306->3306/tcp, 33060/tcp   c3
eff5ef3783e3        httpd               "httpd-foreground"       15 minutes ago      Up 15 minutes       0.0.0.0:8080->80/tcp                c2
bf632191458d        nginx               "nginx -g 'daemon of…"   17 minutes ago      Up 17 minutes       0.0.0.0:80->80/tcp                  c1

$ curl localhost:8080
<html><body><h1>It works!</h1></body></html>

```

cleaning up 

```console
$ sudo docker container ls
CONTAINER ID        IMAGE               COMMAND             CREATED             STATUS     

$ sudo docker container ps
CONTAINER ID        IMAGE               COMMAND             CREATED             STATUS 

$ sudo docker container ps -a
CONTAINER ID        IMAGE               COMMAND                  CREATED             STATUS                         PORTS               NAMES
b11df1035d6d        mysql               "docker-entrypoint.s…"   6 minutes ago       Exited (0) 34 seconds ago                          c3
eff5ef3783e3        httpd               "httpd-foreground"       17 minutes ago      Exited (0) 36 seconds ago                          c2
bf632191458d        nginx               "nginx -g 'daemon of…"   19 minutes ago      Exited (0) 36 seconds ago                          c1
1733fc79806e        mongo               "docker-entrypoint.s…"   37 minutes ago      Exited (0) 31 minutes ago                          mongo
22f59b3f689b        nginx               "nginx -g 'daemon of…"   About an hour ago   Created                                            webhost
ad1df56e92df        nginx               "nginx -g 'daemon of…"   About an hour ago   Exited (0) About an hour ago                       suspicious_nightingale
6369fb6bebf2        nginx               "nginx -g 'daemon of…"   About an hour ago   Exited (0) About an hour ago                       determined_elion
4eab28f4ac02        nginx               "nginx -g 'daemon of…"   About an hour ago   Exited (0) About an hour ago                       wonderful_varahamihira
3cf841f5841d        nginx               "nginx -g 'daemon of…"   About an hour ago   Exited (0) About an hour ago                       brave_hugle
d1e085e5ccb9        hello-world         "/hello"                 19 hours ago        Exited (0) 19 hours ago                            suspicious_roentgen

$ sudo docker rm b11 eff bf6
b11
eff
bf6

$ sudo docker container ps -a
CONTAINER ID        IMAGE               COMMAND                  CREATED             STATUS                         PORTS               NAMES
1733fc79806e        mongo               "docker-entrypoint.s…"   37 minutes ago      Exited (0) 31 minutes ago                          mongo
22f59b3f689b        nginx               "nginx -g 'daemon of…"   About an hour ago   Created                                            webhost
ad1df56e92df        nginx               "nginx -g 'daemon of…"   About an hour ago   Exited (0) About an hour ago                       suspicious_nightingale
6369fb6bebf2        nginx               "nginx -g 'daemon of…"   About an hour ago   Exited (0) About an hour ago                       determined_elion
4eab28f4ac02        nginx               "nginx -g 'daemon of…"   About an hour ago   Exited (0) About an hour ago                       wonderful_varahamihira
3cf841f5841d        nginx               "nginx -g 'daemon of…"   About an hour ago   Exited (0) About an hour ago                       brave_hugle
d1e085e5ccb9        hello-world         "/hello"                 19 hours ago        Exited (0) 19 hours ago                            suspicious_roentgen
```

## 26 Whats Going On in Containers: CLI Process Monitering

* gives me a json array of details of this container:
```console
$ sudo docker container inspect mysql
```
*gives me details on running container

```console
$ sudo docker container stats mysql

```

## 27 Getting a Shell Inside Containers

* start new container interactively

```console
$ docker container run -it
```

* run additional comand in existing container

```console
$ docker container exec -it
```

* run ubuntu container

```console
$ sudo docker container run -it --name ubuntu ubuntu
```

## 28 Docker Networks: Concepts for Private and Public Comms

```console
docker container run -p
```

* for local dev/testing network usualy "just work"
* quick port check with docker container port <container>
* understand how dockers talk to each other
* each container connected to private virtual network "bridge"
* all containers on a virt net can talk to each other without -p
* best practice is to create a new virt net for each App
* make new virtual networks

unless I specify the -p no networks coming to my network is actually getting into my container.


lets say we have a mysql container without -p  and we have a httpd with -p 8080:80


as soon as traffic comes into 8080 it will route into httpd, the httpd is free to talk to the mysql container via listening port


you cant have more than one container listening on same port


## 30 Docker Networks: CLI Management of virtual Networks