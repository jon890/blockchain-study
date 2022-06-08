# 도커 정리

- 도커는 개발하고, 선적하고(shipping), 애플리케이션을 구동하기 위한 오픈 플랫폼이다.
- 도커는 우리의 인프라환경과 애플리케이션을 분리할 수 있게 한다. 그리하여 소프트웨어 전달을 쉽게 할 수 있다.

## 컨테이너란 무엇인가?

- 간단하게 말하면, 컨테이서는 호스트 머신의 다른 프로세스와 격리된 샌드박스화 된 프로세스이다.
- 격리는 커널 네임스페이스와 c그룹의 이점이 있으며, 리눅스의 오래전부터 있던 기능이다.

## 컨테이너 이미지란 무엇인가?

- 컨테이너가 동작할때, 격리된 파일시스템을 사용한다.
- 이러한 커스텀 파일시스템은 컨테이너 이미지를 통해 제공된다.
- 이미지가 컨테이너의 파일시스템을 포함하기 때문에, 애플리케이션을 구동할 때 필요한것 전부가 포함되어야 한다.
- (모든 의존성, 설정, 스크립트, 바이너리 등)

- `chroot`에 친숙하다면, 컨테이너를 확장된 버전의 `chroot`로 생각해보자.

## 간단한 도커 예제 - node.js

1. Dockerfile 작성

   ```Docker
   # synctax=docker/dockerfile:1
   FROM node:12-alpine
   RUN apk add --no-cache python2 g++ male
   WORKDIR /app
   COPY . .
   RUN yarn install --production
   CMD ["node", "src/index.js"]
   EXPOSE 3000
   ```

2. 컨테이너 이미지 빌드

   - `docker build -t getting-started`
   - node:12-alpine 이미지로부터 시작 => 우리 컴퓨터에 이미지가 없으므로 dockerhub로 부터 받아옴
   - `CMD` 명령어는 이미지로부터 컨테이너를 시작할 떄 수행되는 명령어
   - 마지막으로 `-t` 플래그는 우리가 읽을 수 있도록 이미지에 이름을 주는 플래그
   - `.`은 `docker build` 명령어에게 `Dockerfile`를 현재 폴더에서 찾으라고 알려준다.

3. 애플리케이션 컨테이너 시작

   - `docker run -dp 3000:3000 getting-started`
   - `-d`와 `-p` 플래그를 기억하는가?
   - d : detached mode (background mode)
   - p : host's port 3000 to the container's port 3000
   - 포트 매핑을 하지 않으면, 우리는 애플리케이션에 접근할 수 없다.

4. 애플리케이션 업데이트

   - `docker ps` : 컨테이너 목록 출력
   - `docker stop <the-container-id>` (실행중인 컨테이너 중지)
   - `docker rm <the-container-id>` (컨테이너 제거)
   - 위 명령어는 `docker rm -f <the-container-id>`로 축약 가능 `-f` 플래그는 강제로 종료하고 삭제함
   - `docker build -t getting-started .`로 재빌드
   - `docker run -dp 3000:3000 getting-started` 컨테이너 시작
