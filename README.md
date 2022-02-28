# nft-trade-opensea

Pour builder:
1. > cd .../nft-trade-opensea
2. > GOOS=js GOARCH=wasm go build -o main.wasm
3. sudo docker build . -t nft-test

Pour executer:
1. > sudo docker run -it -p 8080:80 nft-test
2. Accéder à http://localhost:8080 sur un navigateur internet