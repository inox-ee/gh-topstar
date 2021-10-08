.DEFAULT_GOAL := all

all:
		go build -v -o a.out
		./a.out
		rm ./a.out

build:
		go build -v -o GitHubTopStar
