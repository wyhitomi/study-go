help:
	./cmd/help.sh

test:
	echo "Testing something..."

build:
	echo "building something..."

run:
	echo "running something..."

docker-build:
	echo "running docker buildx build"
	#docker buildx build --platform=darwin -t ${PWD##*/} .

docker-push:
	echo "running docker push"
	#docker push 
