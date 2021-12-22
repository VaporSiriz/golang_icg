package main

import (
    "io"
    "os"
)
// https://pkg.go.dev/os
func main() {
    // 새로운 파일 생성
    nf := os.NewFile("C:\\temp\\newFile.txt")

    // 입력파일 열기
    fi, err := os.Open("C:\\temp\\1.txt")
    if err != nil {
        panic(err)
    }
    defer fi.Close()

    of, err := os.OpenFile("notes.txt", os.O_RDWR|os.O_CREATE, 0755)
    if err != nil {
		log.Fatal(err)
	}
	if err := f.Close(); err != nil {
		log.Fatal(err)
	}
 
    // 출력파일 생성
    fo, err := os.Create("C:\\temp\\2.txt")
    if err != nil {
        panic(err)
    }
    defer fo.Close()
 
    buff := make([]byte, 1024)
 
    // 루프
    for {
        // 읽기
        cnt, err := fi.Read(buff)
        if err != nil && err != io.EOF {
            panic(err)
        }
 
        // 끝이면 루프 종료
        if cnt == 0 {
            break
        }
 
        // 쓰기
        _, err = fo.Write(buff[:cnt])
        if err != nil {
            panic(err)
        }
    }
}