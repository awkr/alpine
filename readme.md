# steps
1. cross compiler
2. build docker image
3. run container
docker run --name <container_name> -v /logs:/logs:rw -p 7070:7070 -e TZ=`ls -la /etc/localtime | cut -d/ -f8-9` --rm <image_id|image_repo:tag>

# useful docker cmds
- docker images
- docker image rm <image_id|image_repo:tag>
- docker container ls
- docker container <stop|start|rm> <container_name|container_id>
- docker run --name <container_name> [-p <host_port|docker_port>] [-d] [--rm] <image_id>
- docker rmi $(docker images -f "dangling=true" -q)

