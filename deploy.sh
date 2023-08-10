docker build -t snippetboxapp:1.0dev -f docker/Dockerfile \
    --build-arg MYSQL_ROOT_PASSWORD=$(grep "MYSQL_ROOT_PASSWORD" .env | cut -d '=' -f2) \
    --build-arg MYSQL_DATABASE=$(grep "MYSQL_DATABASE" .env | cut -d '=' -f2) \
    --build-arg MYSQL_USER=$(grep "MYSQL_USER" .env | cut -d '=' -f2) \
    --build-arg MYSQL_PASSWORD=$(grep "MYSQL_PASSWORD" .env | cut -d '=' -f2) .

docker run --name snippetbox -p 4000:4000 snippetboxapp:1.0dev
