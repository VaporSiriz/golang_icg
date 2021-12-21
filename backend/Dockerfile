FROM vaporsiriz/bloackchain_staticanalysis:base

# for real image file
COPY . /app
WORKDIR /app

RUN chmod 655 *.sh

# RUN go mod init go_icg
# RUN go get -v github.com/hyperledger/fabric@v1.4.2
# RUN go get -u github.com/gin-gonic/gin
# RUN go mod tidy
# RUN go build

CMD ["./run.sh"]