# budongsango
## 한국 아파트 실거래가 go lang
go 언어로 개발하는 한국 아파트 실거래가 라이브러리입니다.

[공공데이터포털](https://www.data.go.kr)에서 제공하는 [아파트매매 실거래 상세 자료](https://www.data.go.kr/data/15057511/openapi.do)를 사용하여 데이터를 처리합니다.

### 라이브러리
```
go get -u github.com/go-sql-driver/mysql

go get github.com/suapapa/go_hangul/encoding/cp949

```
### 빌드 방법
```
go build apt_real.go
```

### 실행 방법
```
./apt_real(linux) or ./apt_real.exe (window)
```

### 서비스키
```
servicekey 파일에 저장하고 사용하세요.
```
