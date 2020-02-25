
# Docker Notes

## Checking our Docker install and config

* Nginx (web server) container
    * makes it easy to learn docker stuff


gives us info on our docker:
```bash
docker info
```

this is the new way to run docker container (old way is $docker run)
```bash
docker container run
```

### image vs Container

* an image is the app you want to run
* container is an instance of the image running as a process docker hub

* docker container run --publish 80:80 nginx
    * downloads nginx from docker hub
    * start a new container of that image
    * opens the local host

```bash
sudo docker container run --publish 80:80 --detach nginx
sudo docker container ls
sudo docker container ls
CONTAINER ID        IMAGE               COMMAND                  CREATED              STATUS              PORTS                NAMES
ad1df56e92df        nginx               "nginx -g 'daemon of…"   About a minute ago   Up About a minute   0.0.0.0:80->80/tcp   suspicious_nightingale
sudo docker container ls -a
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
