build:
	docker build --tag container-from-scratch .
run:
	docker run --rm -it container-from-scratch sh