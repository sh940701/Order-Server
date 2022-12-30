# (주문)생산자와 (주문)소비자 간 역할 수행 서비스를 제공하는 서버

gin 프레임워크를 이용한 WAS 서버를 통해 주문 생산과 주문 소비가 이루어진다.
생산자와 소비자는 url의 1차 path에서 buyer / seller로 구분된다.

### 시작하기

toml 파일을 자신의 환경에 맞게 바꿔준다.

- config/config.toml

```toml
port = "WAS 서버를 실행시킬 포트번호"
host = "서버와 연동할 mongoDB실행 url"
```

실행을 위한 go package를 설치한다.

- go 의존성 패키지 설치

```shell
go mod init
go mod tidy
```

swagger 파일을 만들어준다

- swagger 패키지 설치

```shell
swag init
```

#

- main 파일 실행

```shell
go run main.go
```

- 일별 주문번호의 정상적인 count를 위해선 orderedlist 컬렉션에 아래와 같이 {daycount : 0} 의 데이터를 먼저 넣어주어야 한다.
  <img width="233" alt="스크린샷 2022-12-28 13 28 09" src="https://user-images.githubusercontent.com/100397903/209757411-95b2d6be-b243-43b6-9817-aa1241f958f1.png">
  <img width="992" alt="스크린샷 2022-12-28 13 26 10" src="https://user-images.githubusercontent.com/100397903/209757348-eee6a96a-af9c-4291-98d2-aa07c86fbb8b.png">

## 프로그램 실행방법

제공하는 기능의 api 리스트는 아래와 같다.
<img width="1097" alt="스크린샷 2022-12-28 16 16 12" src="https://user-images.githubusercontent.com/100397903/209773726-91f8a9ce-2e89-4a86-9af8-8f127fa90f53.png">

테스트를 위해선 프로그램 실행 후 http://localhost:8080/swagger/index.html#/ 페이지로 접속하면 swagger를 통해 실행할 수 있다.

### seller

<hr />

- seller - 메뉴 신규 등록
  <img width="721" alt="스크린샷 2022-12-28 13 40 55" src="https://user-images.githubusercontent.com/100397903/209758419-ab05e02f-63be-4c1c-8982-d1efdcee6dc9.png">

- seller - 메뉴 업데이트
  <img width="755" alt="스크린샷 2022-12-28 15 22 15" src="https://user-images.githubusercontent.com/100397903/209767648-b9548157-dfb6-4509-a49c-6b57183a011f.png">

- seller - 메뉴 삭제
  <img width="720" alt="스크린샷 2022-12-28 15 24 16" src="https://user-images.githubusercontent.com/100397903/209767851-bbde25bf-c648-4947-a589-b8e436d8741e.png">

- seller - 주문 상태 업데이트
  <img width="734" alt="스크린샷 2022-12-28 15 43 16" src="https://user-images.githubusercontent.com/100397903/209770013-3d306c54-bcdb-40d4-99b7-f5e882c8b38a.png">
  <img width="725" alt="스크린샷 2022-12-28 15 43 27" src="https://user-images.githubusercontent.com/100397903/209770018-570facd3-5bf0-43dc-a363-aacfe199e31e.png">
  <img width="718" alt="스크린샷 2022-12-28 15 43 37" src="https://user-images.githubusercontent.com/100397903/209770110-86cf149a-afbc-4972-8521-6e0fc48738a7.png">
  <img width="725" alt="스크린샷 2022-12-28 15 43 47" src="https://user-images.githubusercontent.com/100397903/209770129-351fd625-5f19-4f7e-bb31-7df7ba546e71.png">

- seller - 주문 리스트 가져오기(page parameter 별 5개씩 반환)
  <img width="720" alt="스크린샷 2022-12-28 15 49 05" src="https://user-images.githubusercontent.com/100397903/209770547-9b7fa894-dc67-47e9-a0fa-bb97ba3654e7.png">

- seller - 추천 메뉴 업데이트
  <img width="721" alt="스크린샷 2022-12-28 15 54 28" src="https://user-images.githubusercontent.com/100397903/209771204-6f33b476-3b23-4065-b551-7e31a97f2479.png">

### buyer

<hr />

- buyer - 모든 메뉴 가져오기(page parameter 별 5개씩 반환)
  <img width="725" alt="스크린샷 2022-12-28 16 26 56" src="https://user-images.githubusercontent.com/100397903/209774950-8317b95b-4b5f-4111-8f4d-c58a4d88a9b2.png">

- buyer - 카테고리별 정렬된 메뉴 조회(평점, 주문 수, 추천메뉴)
  <img width="724" alt="스크린샷 2022-12-28 16 28 39" src="https://user-images.githubusercontent.com/100397903/209775155-adbe8c45-52d1-439f-bd2d-8ba7df8de7a2.png">

- buyer - 주문하기(주문이 완료되면 해당 메뉴의 orderedcount + 1, limit - 1이 이루어지고, limit이 0인 상태로 주문을 하게 되면 abort된다.)
  <img width="716" alt="스크린샷 2022-12-28 16 30 37" src="https://user-images.githubusercontent.com/100397903/209775380-ef9052b3-e071-4806-a6c2-525ae555af18.png">

- buyer - 주문 추가하기(주문의 상태가 배달중/배달완료일 경우 새 주문으로 추가됨)
  <img width="718" alt="스크린샷 2022-12-28 17 27 55" src="https://user-images.githubusercontent.com/100397903/209782481-4ae3e50f-1a2b-4fec-a80d-4aefcce9e9fa.png">
  <img width="726" alt="스크린샷 2022-12-28 17 30 04" src="https://user-images.githubusercontent.com/100397903/209782682-65019cf7-a9ab-4262-9a19-e88b1691175c.png">

- buyer - 주문의 현재 상태 가져오기
  <img width="978" alt="스크린샷 2022-12-28 17 33 33" src="https://user-images.githubusercontent.com/100397903/209783145-411e12c7-c5a7-4777-8414-028cb7e23134.png">

- buyer - 주문에서 메뉴 변경하기(주문의 상태가 배달중/배달완료일 경우 변경 불가 메시지 반환)
  <img width="719" alt="스크린샷 2022-12-28 17 36 36" src="https://user-images.githubusercontent.com/100397903/209783518-59aac0f5-b547-44bc-b388-a3d70d8fb04f.png">
  <img width="952" alt="스크린샷 2022-12-28 17 35 59" src="https://user-images.githubusercontent.com/100397903/209783551-efdd6de8-5fb4-48b6-9c31-cf1bc0e51b23.png">

- buyer - 리뷰 작성하기(이미 후기가 작성되었을 경우 에러 반환), body에 포함된 score는 메뉴의 평점에 반영된다.
  <img width="725" alt="스크린샷 2022-12-28 17 39 26" src="https://user-images.githubusercontent.com/100397903/209783879-155e78ac-f6cc-447c-ba83-5067fa8d2549.png">
  <img width="725" alt="스크린샷 2022-12-28 17 40 13" src="https://user-images.githubusercontent.com/100397903/209783961-fadc7bc9-728e-46ad-9a95-dc0d0fee2d21.png">

- buyer - 메뉴별 리뷰 확인하기(평점도 함께 반환, 리뷰가 없다면 에러반환)
  <img width="958" alt="스크린샷 2022-12-28 18 28 26" src="https://user-images.githubusercontent.com/100397903/209790149-a8da3382-00e7-4db2-8e43-6c49c5098391.png">
  <img width="1087" alt="스크린샷 2022-12-28 18 30 23" src="https://user-images.githubusercontent.com/100397903/209790432-fa194221-2b54-485c-966d-e8cfb936ed34.png">

### 코드 구성

코드 구성에 대한 정보는 우측 페이지에서 확인할 수 있다.
https://quilt-stoplight-9b6.notion.site/Docs-0c6140c809904bdcb46a41fbbea48369
