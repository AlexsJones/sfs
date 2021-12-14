
docker:
	docker build -t tibbar/sfs:latest .
	docker push tibbar/sfs:latest
helm:
	cd chart && helm upgrade --install sfs . && cd ../