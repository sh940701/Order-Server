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

* 이후 controller/buyercontroller.go 파일 상단의 daycountID 상수를 생성된 daycount Document의 ObjectId로 변경해준다.  
  <img width="578" alt="스크린샷 2023-01-03 15 45 17" src="https://user-images.githubusercontent.com/100397903/210310815-090f3766-eacb-477d-868b-832464c109f7.png">

### 코드 구성

코드 구성에 대한 정보는 우측 페이지에서 확인할 수 있다.
https://quilt-stoplight-9b6.notion.site/Docs-0c6140c809904bdcb46a41fbbea48369

# 피드백 반영 개선사항

## 피드백

<img width="607" alt="스크린샷 2023-01-03 16 14 45" src="https://user-images.githubusercontent.com/100397903/210313748-5e0e35d5-1057-4511-b809-7860839a8e3f.png">
<img width="610" alt="스크린샷 2023-01-03 15 57 37" src="https://user-images.githubusercontent.com/100397903/210312054-a3698710-97bf-4cb5-9134-edead7f68b24.png">
  
## buyer - 주문하기
  - 메뉴를 주문, 추가, 변경할 때 수량을 기입할 수 있도록 수정
  - 주문과 관련하여 메뉴 추가, 변경, 새 주문시 각 메뉴의 현재 주문 가능한 Limit을 확인하고 그 이상일 경우 abort하는 로직 전체적으로 추가
  - IsReviewed, Status 값을 각각 false, “주문접수” 로 default 설정을 해 주어 주문시 payload에서 제외
    <img width="744" alt="스크린샷 2023-01-03 15 04 37" src="https://user-images.githubusercontent.com/100397903/210307021-26603a23-af7a-4e03-8ddb-ca3afc14458a.png">
    <img width="723" alt="스크린샷 2023-01-03 15 07 45" src="https://user-images.githubusercontent.com/100397903/210307267-a8511d50-b27f-41d4-a5f5-1d5a84b948f3.png">

## buyer - 주문에 메뉴 추가시 중복 체크

- 주문에 메뉴를 추가할 경우, 해당 주문에 이미 같은 메뉴가 있는지 확인하는 로직을 추가
  <img width="735" alt="스크린샷 2023-01-03 15 16 59" src="https://user-images.githubusercontent.com/100397903/210308124-3c8799be-f047-4ec9-83a3-86af025ba826.png">
  <img width="725" alt="스크린샷 2023-01-03 15 17 20" src="https://user-images.githubusercontent.com/100397903/210308154-3f73cda5-ce80-44bd-9f8a-7bf7b602b052.png">
  <img width="735" alt="스크린샷 2023-01-03 15 18 23" src="https://user-images.githubusercontent.com/100397903/210308243-bc862bd1-7e22-47a2-a13a-94e326912b2f.png">
  <img width="724" alt="스크린샷 2023-01-03 15 23 32" src="https://user-images.githubusercontent.com/100397903/210308722-8a2c233c-45a1-41eb-8018-9e38f82bb5bd.png">

## buyer - 리뷰 작성시

- 평점이 5를 초과하면 abort하는 로직 추가
- 해당 주문의 각 메뉴에 대해 리뷰가 가능하고, 주문의 모든 메뉴에 대한 리뷰가 완료되었을 시, 주문의 전체 리뷰 여부를 true로 업데이트하는 로직 추가
  <img width="730" alt="스크린샷 2023-01-03 15 25 26" src="https://user-images.githubusercontent.com/100397903/210308867-28702b65-67c1-45ae-bd90-0fe198a6a95f.png">
  <img width="727" alt="스크린샷 2023-01-03 15 26 00" src="https://user-images.githubusercontent.com/100397903/210308923-7a6f0ecd-f1d4-4c50-884d-e5f68d649ecf.png">
  <img width="741" alt="스크린샷 2023-01-03 15 26 11" src="https://user-images.githubusercontent.com/100397903/210308946-f477dc4e-4119-4132-9469-14dc220caf3e.png">
  <img width="735" alt="스크린샷 2023-01-03 15 27 29" src="https://user-images.githubusercontent.com/100397903/210309091-3079789d-2c77-42c5-b17a-cc38d7c8e880.png">

## seller - 새로운 메뉴 데이터 추가시

- OrderedCount, Avg, Suggestion, IsVisible값을 각각 0, 0, false, true로 default 설정을 해 주어 새로운 메뉴 데이터 추가시 payload에서 제외
  <img width="719" alt="스크린샷 2023-01-03 15 30 02" src="https://user-images.githubusercontent.com/100397903/210309335-0717b78b-3885-4c48-969d-e49a2844dd79.png">

## seller - 메뉴 데이터 삭제시

- 메소드를 POST -> DELETE 로 변경
- 메뉴 데이터에 IsVisible 멤버변수를 추가하여 삭제시 플래그를 활용하여 visible을 false로 변경 + 데이터를 조회할 때 기본적으로 IsVisible이 true인 데이터만 가져오도록 설정
  <img width="966" alt="스크린샷 2023-01-03 15 39 25" src="https://user-images.githubusercontent.com/100397903/210310262-3a36b8a8-5090-48c2-a464-827c17da885e.png">
  <img width="988" alt="스크린샷 2023-01-03 15 40 36" src="https://user-images.githubusercontent.com/100397903/210310360-0af5c3c7-53ec-4b7f-b641-2e7e63886ff7.png">
