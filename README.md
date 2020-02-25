
# Docker Notes

## 19 Checking our Docker install and config

* Nginx (web server) container
    * makes it easy to learn docker stuff


gives us info on our docker:
```console
docker info
```

this is the new way to run docker container (old way is $docker run)
```console
docker container run
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
docker container run --publish 8080:80 --name webhost -d nginx:1.11 nginx -T
```
* 1.11 --> change version of IMAGE
* -T --> change CMD run on start
* 8080:80 --> change host listening port


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

