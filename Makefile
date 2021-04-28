run:
	docker-compose build && docker-compose up -d

stop:
	docker-compose down

signupTest:
	cd server/users/ && go test -v -cover ./... && cd ../../

mysql:
	docker exec -it mysql_db /bin/bash