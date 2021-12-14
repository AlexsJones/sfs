
docker:
	docker buildx build --platform linux/amd64 -t tibbar/sfs:latest .
	docker push tibbar/sfs:latest
helm:
	cd chart && helm upgrade --install sfs . && cd ../
