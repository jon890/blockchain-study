## DB 영속화

- 눈치 못챗겠지만, 우리의 해야할 일 목록은 우리가 컨테이너를 시작할 떄마다 깨끗이 지워지고 있다.
- 왜 이럴까?
- 컨테이너가 어떻게 동작하는지 더 깊이 이해해보자.

### 컨테이너의 파일 시스템

- 컨테이너가 동작할 때, 파일 시스템을 위하여 이미지의 여러 계층을 사용한다.
- 각 컨테이너는 개개의 scratch space에 접근하여 파일을 생성, 수정, 삭제 할 수 있다.
- 같은 이미지를 사용하더라도, 다른 컨테이너에서는 그러한 변경들을 볼 수 없다.

#### 연습으로 확인해보자

- 두 개의 컨테이너를 시작하고 각각 파일을 생성해보자.
- 한 컨테이너에서 생성한 파일을 다른 컨테이너에서 볼 수 없다는 것을 확인해보자.

1. 우분투 컨테이너를 시작하고 `/data.txt`라는 1 ~ 10000 사이의 무작위 숫자를 생성하는 파일을 만들자.

   - `docker run -d ubuntu bash -c "shuf -i 1-10000 -n 1 -o /data.txt && tail -f /dev/null"`
   - 위 명령어는, bash 셸을 시작하고 두 개의 명령어를 실행한다.
   - 첫 번째 명령어는 한 개의 무작위 숫자를 뽑아서 `/data.txt`에 저장한다.
   - 두 번째 명령어는 컨테이너가 실행될 떄 파일을 계속 지켜볻다.

2. 컨테이너의 `exec`로 들어가 결과물을 확인함으로써 검증해보자.

   - `docker exec <container-id> cat /data.txt`
   - 무작위 수를 확인할 수 있다!

3. 다른 우분투 컨테이너를 시작했지만 우리는 같은 파일이 없다는 것을 확인할 수 있다.

   - `docker run -it ubuntu ls /`
   - `data.txt` 파일이 없다! 첫 번째 컨테이너의 scratch 영역에 쓰였기 떄문이다

4. 첫 번째 컨테이너를 `docker rm -f <container-id>` 명령어를 통해 지우자

## 컨테이너 볼륨

- 앞선 실험과 같이, 각 컨테이너는 시작 될때 이미지 정의로 부터 시작된다는 것을 보았다.
- 컨테이너를 생성하고, 수정하고, 파일을 지우는 동안에 그 변화들은 컨테이너가 제거되면 잃어버리게 되고 모든 변화들은 컨테이너로부터 독립적이다.
- 볼륨을 사용하면 변화시킬 수 있다.

- 볼륨은 컨테이너 뒤에 있는 호스트 머신의 특정한 파일 시스템 주소로 연결할 수 있게 한다.
- 컨테이너의 폴더가 마운트 되면, 변화들을 호스트 머신의 폴더에서도 보여진다.
- 같은 디렉토리를 마운트하고 컨테이너를 재시작해도, 우리는 같은 파일을 볼 수 있다.

- 두 가지 종류의 볼륨이 있다. 우리는 결국 두 가지를 모두 사용하게 될 것이다, 하지만 먼저 이름 있는 볼륨부터 시작해보자.

## 해야할 일 영속화

- 기본적으로, 해야할 일 앱은 SQLite Database의 `/etc/todos/todo.db`에 데이터를 저장한다.
- SQLite에 친숙하지 않아도 걱정하지 말아라!
- 간단한 RDB이며 하나의 파일에 데이터를 저장한다.
- 해야할 일 앱이 큰 규모의 앱이 아니므로, 작은 데모로써는 적합하다!
- 다른 데이터베이스 엔진으로 바꾸는것은 후에 얘기하자.

- 데이터베이스가 하나의 파일을 가짐으로, 호스트에 파일을 영속화하고 다음 컨테이너에게 사용하게 할 수 있다.
