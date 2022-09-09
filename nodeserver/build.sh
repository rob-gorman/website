ver=$1
docker build -t ghcr.io/rob-gorman/nodeserver:${ver} .
docker push ghcr.io/rob-gorman/nodeserver:${ver}
# docker run -p3000:3000 --name nodeserver nodeserver:${ver}