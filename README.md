Atlant GRPC
====================

#### Install

    sudo apt install -y protobuf-compiler make docker-ce python
    pip install -U --user docker-compose
    go get google.golang.org/protobuf/cmd/protoc-gen-go \
        google.golang.org/grpc/cmd/protoc-gen-go-grpc && \
     
#### Build and tests

    make test
    make start

#### Test with client

    ./client --server-uri 127.0.0.1:8080 fetch \
        --fetch-uri https://support.staffbase.com/hc/en-us/article_attachments/360009197011/username-password-recovery-code.csv \
        --sh \
        --separator ";"
        
    ./client --server-uri 127.0.0.1:8080

#### Stop

    make stop
