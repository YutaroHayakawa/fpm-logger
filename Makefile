build-image:
	docker build -t yutarohayakawa/fpm-logger:latest .

push-image:
	docker push yutarohayakawa/fpm-logger:latest
