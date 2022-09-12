ver=$1
docker stop server
docker rm server
docker build --tag ghcr.io/rob-gorman/webserver:${ver} .
# docker push ghcr.io/rob-gorman/webserver:${ver}
docker run -d -p3000:3000 --name server ghcr.io/rob-gorman/webserver:${ver}