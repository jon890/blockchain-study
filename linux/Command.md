# 리눅스 커맨드 및 기본 개념 정리

## Termial

- 컴퓨터와 소통하기 위해서 키보드로 명령어를 입력하어 사용하는 것을 말함
- shell을 기본적으로 사용하고 있다
- shell은 기본적으로 bash, tcsh 등이 사용되며 Windows에서는 cmd를 기반으로 사용
- shell: 리눅스의 핵심인 커널과 사용자를 연결해주는 인터페이스

### ls [옵션] [파일명]

- 현재 디렉토리 내의 파일과 디렉터리 정보를 출력
- 옵션

  1. a: 디렉터리에 있는 모든 파일을 (.으로 시작하는 파일 포함)을 출력
  2. i: 파일의 아이노드(inode, 색인번호) 번호를 출력
  3. h: 파일 크기를 사람이 보기 쉬운 단위로 출력 (k: 킬로바이트, m: 메가바이트)
  4. l: 파일의 상세정보를 함께 출력 (소유자, 권한, 크기, 날짜)
  5. m: 파일을 쉼표로 구분하여 가로로 출력
  6. s: kb 단위의 파일 크기를 출력
  7. t: 최근 생성된 시간 순으로 파일 출력
  8. F: 파일 종류 별로 끝에 특수 문자 표시 (실행파일: \*, 디렉터리: /, 심볼링크: @, FIFOvkdlf : | 소켓파일 =)
  9. R: 지정한 디렉터리 아래에 있는 하위 디렉터리와 파일을 포함하여 출력
  10. S: 파일 크기가 큰 순서대로 출력

### rm [옵션] [파일이름]

- 파일이나 디렉터리를 삭제할 때 사용
- 옵션
  1. i: 파일이나 디렉터리가 삭제될 때마다 확인
  2. f: 사용자에게 확인하지 않고 삭제
  3. v: 각각의 파일 지우는 정보를 자세하게 모두 보여줌
  4. r: 해당 디렉터리의 하위 디렉터리까지 모두 삭제 (폴더를 삭제할 떄 많이 사용)

### cat [옵션] [파일이름]

- 파일의 내용을 볼 수 있는 명령어
- cat > [파일이름] 형태로 명령어를 작성했을 경우, 파일 생성 및 데이터 입력도 가능
  - 파일 저장 : Ctrl + d, 파일 종료 : Ctrl + z
- cat [파일명] | more: 엔터키를 입력할 때마다 한 줄씩 내려가면서 확인이 가능
- cat [파일명] | less: 화살표 위, 아래키로 페이지 올림, 내림이 가능

  - more과 less 상태에서 q를 누르면 종료

- 옵션
  1. n: 파일을 출력할 떄, 라인에 번호를 붙여 출력
  2. b: 공백 외의 글자가 있는 라인에 번호를 붙여 출력
