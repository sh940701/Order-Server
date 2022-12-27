# (주문)생산자와 (주문)소비자 간 역할 수행 서비스를 제공하는 서버
gin 프레임워크를 이용한 WAS 서버를 통해 주문 생산과 주문 소비가 이루어진다.
생산자와 소비자는 url의 1차 path에서 buyer / seller로 구분된다.

### 시작하기

toml 파일을 자신의 환경에 맞게 바꿔준다.
*  config/config.toml
  ```toml
  port = "WAS 서버를 실행시킬 포트번호"
  host = "서버와 연동할 mongoDB실행 url"
  ```

실행을 위한 go package를 설치한다.
*  go 의존성 패키지 설치
  ```shell
  go mod init
  go mod tidy
  ```
  
swagger 파일을 만들어준다
*  swagger 패키지 설치
  ```shell
  swag init
  ```
  
*  main 파일 실행
  ```shell
  go run main.go
  ```
  
### 프로그램 실행방법
제공하는 기능의 api 리스트는 아래와 같다.
<img width="1167" alt="스크린샷 2022-12-25 23 42 41" src="https://user-images.githubusercontent.com/100397903/209472370-88fcbb61-aa57-4a87-acf4-131b40265b5f.png">
실행 테스트를 위해선 run go main.go 를 통해 프로그램을 실행한 뒤
http://localhost:8080/swagger/index.html#/ 페이지로 접속하면 swagger를 통해 실행할 수 있다.

### 코드 구성
코드 구성에 대한 정보는 우측 페이지에서 확인할 수 있다.
https://quilt-stoplight-9b6.notion.site/Docs-0c6140c809904bdcb46a41fbbea48369
