swag

    swag init -g internal/casbin/delivery/http/http.go
    
grpc

     protoc --go_out=plugins=grpc:. *.proto
    
grpc 

    sudo mkdir -p /mongodata
    
    sudo docker run -it -v mongodata:/data/db --name mongodb -d mongo
    
    docker run -p 3306:3306 --name mysql -e MYSQL_ROOT_PASSWORD=1234 -e MYSQL_DATABASE=db -d mysql
    
    docker network create rabbits
    
    docker run -d --rm --net rabbits --hostname rabbit-1 --name rabbit-1 rabbitmq
    
    config run server reload
    
    sudo apt-get install inotify-tools
    
    git clone https://github.com/alexedwards/go-reload.git
    
    cd go-reload
    
    chmod +x go-reload
    
    sudo mv go-reload /usr/local/bin/

casbin

example

    p, alice, /alice_data/*, GET
    
    p, alice, /alice_data/resource1, POST
    
    p, bob, /alice_data/resource2, GET
    
    p, bob, /bob_data/*, POST
    
    p, cathy, /cathy_data, (GET)|(POST)
    
docker run \
    --name kibana \
    --link elasticsearch:elasticsearch \
    -p 5601:5601 \
   -d \
    docker.elastic.co/kibana/kibana:7.3.0 


docker run \
    --name elasticsearch \
    -p 9200:9200 \
    -p 9300:9300 \
    -e "discovery.type=single-node" \
    -d  \
    docker.elastic.co/elasticsearch/elasticsearch:7.3.0 
    
docker run --name some-redis -d -p 6379:6379  redis
   
  docker run -d --name kong-database \
                 -p 5434:5432 \
                 -e "POSTGRES_USER=kong" \
                 -e "POSTGRES_DB=kong" \
                 -e "POSTGRES_PASSWORD=kong" \
                 postgres
                  
   docker run -d --name kong \
       -e “KONG_LOG_LEVEL=debug” \
       -e “KONG_DATABASE=postgres” \
       -e “KONG_PG_HOST=kong-database” \
       -e “KONG_PROXY_ACCESS_LOG=/dev/stdout” \
       -e “KONG_ADMIN_ACCESS_LOG=/dev/stdout” \
       -e “KONG_PROXY_ERROR_LOG=/dev/stderr” \
       -e “KONG_ADMIN_ERROR_LOG=/dev/stderr” \
       -e “KONG_ADMIN_LISTEN=0.0.0.0:5000 ssl” \
       -p 9000:8000 \
       -p 9443:8443 \
       -p 9001:5000 \
       kong:latest
    
docker start elasticsearch
docker start kibana
docker start my-postgres
docker start mysql
docker start some-rabbit
docker start mongodb
docker start some-redis