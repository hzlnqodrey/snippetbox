# FOR THIS TO WORK ( U NEED 2 CONTAINER: APP CONTAINER and DB CONTAINER. CONNECT THEM)

# CREATE MYSQL SERVER CONTAINER FIRST
#  Additionally, since you're running your application in another container and you want them to communicate, 
#  you should also ensure that both containers are connected to the same Docker network.
#  This will allow them to communicate using their service names.
docker network create my-network
docker run -d --name mysql-container -e MYSQL_ROOT_PASSWORD={password} --network my-network mysql:8.0.15

# INITIAL BUILD
docker build -t snippetboxapp:1.0dev \
  --build-arg MYSQL_ROOT_PASSWORD=$(grep "MYSQL_ROOT_PASSWORD" .env | cut -d '=' -f2) \
  --build-arg MYSQL_DATABASE=$(grep "MYSQL_DATABASE" .env | cut -d '=' -f2) \
  --build-arg MYSQL_USER=$(grep "MYSQL_USER" .env | cut -d '=' -f2) \
  --build-arg MYSQL_PASSWORD=$(grep "MYSQL_PASSWORD" .env | cut -d '=' -f2) .

# BUILD NEW VERSION
docker build -t snippetboxapp:{new_tag} .

# then try to run application env
docker run --name snippetboxapp -p 4000:4000 --network my-network snippetboxapp:latest

# run in background
docker run -d --name snippetboxapp -p 4000:4000 --network my-network snippetboxapp:latest


