



```bash
cd cmd/ 

GOFLAGS=-mod=vendor go run main.go  # surpise 

GOFLAGS=-mod=mod go run main.go -h  # The original result 


```