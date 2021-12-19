# backend
<모듈화 진행>  
$ go mod init go_icg  
$ go get -v github.com/hyperledger/fabric@v1.4.2  
$ go get -u github.com/gin-gonic/gin  
$ go mod tidy  

# dockerization
1. go-z3 모듈을 같은 경로에 다운로드.      
2. dockerfile을 이용하여 이미지 생성 후 container 생성 및 run    
3. go-z3/vendor/z3/build/libz.so 를 /lib/x86_64-linux-gnu/ 으로 copy
4. z3.go 파일에 libz3.a -> libz3.so로 변경  
5. go build 하여 실행이 잘되는지 확인.  
   
위의 것은 base image 생성 완료하였음.  
vaporsiriz/bloackchain_staticanalysis:base를 바탕으로 하면 됨.  

# 실행
$ cd backend

$ docker build -t bloackchain_staticanalysis .

$ docker run --rm -p 8080:8080 bloackchain_staticanalysis