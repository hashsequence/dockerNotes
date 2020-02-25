
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


## 29 Docker Networks: CLI Management of virtual Networks

* show networks

```console
$ docker network ls
```

* docker network inspect

```console
$ inspect a network
```

* create network

```console
$ docker network create --driver
```

* attach a network to container

```console
$ docker network connect
```

* detach a network from container

```console
$ docker network disconnect
```

example:

```console

$ sudo docker network create my_app_net
43ff195809b0a76ca326c7581b91f76d0fff4d0a4611caf309adc1ebb4f98b13
$ sudo docker network ls
NETWORK ID          NAME                DRIVER              SCOPE
1b7f0aad50a8        bridge              bridge              local
e0c703b821cd        host                host                local
43ff195809b0        my_app_net          bridge              local
bd6a5a434f32        none                null                local

```

run this below to start container in your new network:


```console
$ sudo docker container run -d --name new_nginx --network my_app_net nginx
2109f3c42a36673d6ea2dc58bc1e484e485e42bab0213a7b88a5fad461351cb4
```

* connect my_app_net to the webhost:


```console
$ sudo docker container ls
CONTAINER ID        IMAGE               COMMAND                  CREATED              STATUS              PORTS                 NAMES
2109f3c42a36        nginx               "nginx -g 'daemon of…"   About a minute ago   Up About a minute   80/tcp                new_nginx
e28a17ea86a1        nginx               "nginx -g 'daemon of…"   17 minutes ago       Up 17 minutes       0.0.0.0:80->80/tcp    webhost
5e56ba83a0a4        mysql               "docker-entrypoint.s…"   6 hours ago          Up 6 hours          3306/tcp, 33060/tcp   mysql
a089cae98cce        nginx               "nginx -g 'daemon of…"   6 hours ago          Up 6 hours          80/tcp                nginx
$ sudo docker network ls
NETWORK ID          NAME                DRIVER              SCOPE
1b7f0aad50a8        bridge              bridge              local
e0c703b821cd        host                host                local
43ff195809b0        my_app_net          bridge              local
bd6a5a434f32        none                null                local
$  docker network connect 43ff195 e28a17

```

* now inspect webhost and go to Networks properties to see what its connected to

```console
$ sudo docker container inspect e28a17
```

*disconnection container from network

```console
$ sudo docker network disconnect 43ff195 e28a17

```

## 31 Docker Networks: DNS and how containers find each other

* forget static ip's as a way for containers to see each other, because its too dynamic
* docker has dns naming
* detached mode, shown by the option --detach or -d, means that a docker container runs in the background of your terminal.
* Docker has DNS naming
* Detached mode, shown by the option --detach or -d, means that a Docker container
runs in the background of your terminal.

lets have my_nginx pinging new_nginx:

```console
$ sudo docker container run -it -d --name new_nginx --network my_app_net nginx:alpine
3fc886d83c50199cb53a4397016317e55bf16e797e1157f24051dcac36f89851
$ sudo docker container ls
CONTAINER ID        IMAGE               COMMAND                  CREATED              STATUS              PORTS                NAMES
3fc886d83c50        nginx:alpine        "nginx -g 'daemon of…"   About a minute ago   Up About a minute   80/tcp               new_nginx
b80070fdd01b        nginx:alpine        "nginx -g 'daemon of…"   7 minutes ago        Up 7 minutes        80/tcp               my_nginx
4d460300cfa1        nginx               "nginx -g 'daemon of…"   20 minutes ago       Up 20 minutes       0.0.0.0:80->80/tcp   webhost
$ sudo docker container exec -it new_nginx ping my_nginx
PING my_nginx (172.18.0.2): 56 data bytes
64 bytes from 172.18.0.2: seq=0 ttl=64 time=0.394 ms
64 bytes from 172.18.0.2: seq=1 ttl=64 time=0.220 ms
64 bytes from 172.18.0.2: seq=2 ttl=64 time=0.217 ms
64 bytes from 172.18.0.2: seq=3 ttl=64 time=0.185 ms
64 bytes from 172.18.0.2: seq=4 ttl=64 time=0.243 ms
64 bytes from 172.18.0.2: seq=5 ttl=64 time=0.235 ms
64 bytes from 172.18.0.2: seq=6 ttl=64 time=0.216 ms
```

--- my_nginx ping statistics ---
7 packets transmitted, 7 packets received, 0% packet loss
round-trip min/avg/max = 0.185/0.244/0.394 ms


* host name will always be the same
* link defines link between containers
* compose will automatically spin up virtual network

## 32 Assignment: Using Containers for CLI Testing

* use different linux distro containers to  check curl cli tool version

* use two different terminal windows to start bash in centos:7 and ubuntu 14.04 using -it 

* ensure curl is up to date on both distro

* ensure curl is up to date on both distro

* centos : yum update curls 

* check curl --version


solution:

on ubuntu: 

```console
$ sudo docker container run -it -d --name ubuntu ubuntu:14.04
$ sudo docker container exec -it ubuntu bash

```


centos:

```console
$ sudo docker container run -it -d --name centos centos:7
$ sudo docker container exec -it centos bash
[root@7fae1027116f /]# yum update curls
```

## 34 Assignment: DNS Round Robin Test

* know how to use -it to get shell in container
* know how to run a container
* understand basics of dns
* dns round robin test is two different hosts with dns alias that responds to the same dns name, for example google.com have multiple dns records.
* create a virtual network (default bridege driver)
* create two containes from elasticsearch:2
* use --network-alias search when creating them to give them an additional dns name to respond to
* run alpine nslookup search with --net to see the two containers list for the same DNS name
* run centos curl -s search:9200 with --net multiple times until you see both "name" fields show
*  If instead you’d like Docker to automatically clean up the container and remove the file system when the container exits, you can add the --rm flag


solution:


```console
 $ sudo docker container run -it -d --name el2 --network my_app_net --alias-net search elasticsearch:2
 $ sudo docker container run -it -d --name el2 --network my_app_net --alias-net search elasticsearch:2
 $ sudo docker container run --rm --net my_app_net  alpine nslookup search alpine nslookup search
 $ sudo docker container run --net my_app_net centos curl -s search:9200
 
```

## 35 What's In An Image (and What Isn't)

whats an image:

* app binaries and dependencies
* metadata about image data nd how to run image
* not a complete OS, no kernel (eg no drivers)
* small as one file like a golang static binary
* big as ubuntu distro with apt, and apache, php installed
 
## 36 the mighty hub, using docker hub registry images

hub.docker.com

official vs good images vs bad images
download images and image tags

```console
docker pull nginx:1.11  (the tag here is the version)
```

## 37 images and their layers: Discover the Image Cache

* history and inspect commands
* copy on write concept
* image layers
* union file system

history and layers

```console
$ docker history nginx:latest
WARNING: Error loading config file: /home/avwong13/.docker/config.json: stat /home/avwong13/.docker/config.json: permission denied
IMAGE               CREATED             CREATED BY                                      SIZE                COMMENT
2bcb04bdb83f        10 days ago         /bin/sh -c #(nop)  CMD ["nginx" "-g" "daemon…   0B
<missing>           10 days ago         /bin/sh -c #(nop)  STOPSIGNAL SIGTERM           0B
<missing>           10 days ago         /bin/sh -c #(nop)  EXPOSE 80                    0B
```
* container layer
    * lets say you have an apache image and you run container 1 and container 2 of that image and you make a change in container 2 then in terms of file space will on show the differences between the image base and whats happening in the containers

* copy on write : copy and write is when the we make a change in a file in a container then the file system takes a copy of that file  from the base and writes a copy into that container


* inspect gives you back the metadata

* ExposedPorts tells you what ports to open to access it

* each layer is uniquely identified and only stored once on a host

## 38 image tagging and pushing to docker hub

tagging:

```console
$docker image tag
```

tag is not exactly version, but its more like a pointer to that commit


how to tag images:


```console
$ sudo docker image tag --help

Usage:	docker image tag SOURCE_IMAGE[:TAG] TARGET_IMAGE[:TAG]

Create a tag TARGET_IMAGE that refers to SOURCE_IMAGE
```

how to push to dockerhub:


```console
$ sudo docker image ls
REPOSITORY           TAG                 IMAGE ID            CREATED             SIZE
api                  latest              e9b9008933c3        4 days ago          794MB
$ sudo docker login
$ sudo docker image push <repository>
```

## 39 Building Images: The Dockerfile Basics

Dockerfile syntax
```dockerfile
From <another image>
#all images must have a FROM
#usually from a minimal linux distro like debian or alpine

ENV NGINX_VERSION 1.11.10-1-jessie
# optional environment variable thats used in later lines and set as envvar when
#container is running

RUN <commands>
RUN <more commands and so on>
RUN ln -sf /dev/stdout /var/log/nginx/access.log \
    && ln -sf dev/stderr /var/log/nginx/error.log
#foward request and error logs to docker log collector

EXPOSE
#expose these ports on the docker virtual network
#you still need to use -p or -P to open/forward these ports on host

CMD ["nginx", "-g", "daemon off;"]
#required: run this command when container is launged
#only one cmd allowed, so if there are multiple, last one wins

```

## 40 Building Images: Running Docker Builds

have a Dockerfile
build docker file in your current directory:


```console
$docker image build -t <tag name> .
```

## 41 Building Images: Extending Official Images

dockerfile
```dockerfile
  FROM nginx:latest
  #highly reccommend you always pin version for anything beyond dev/learn
  WORKDIR /usr/share/nginx/html
  #change working directory to root of nginx webhost
  #using WORKDIR is prefered to using RUN cd /some/path

  COPY index.html index.html
  #missing CMD because there is already a CMD in the FROM image

```

```console
$ docker image build -t nginx-custom
```