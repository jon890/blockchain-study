# NestJS 에서 Provider란?

- 공식 문서를 번역해서 읽어보고 의미를 파악해본다.
- 스프링의 개념을 익히 알고 있어서 크게 어려운 부분은 없었다.
- 출처 : https://docs.nestjs.com/providers

## Provider

- Provider들은 Nest의 근간이 되는 개념이다.
- 많은 Nest 클래스들은 Provider로 취급된다. (서비스, 레퍼지토리, 팩토리, 헬퍼 등)
- Provider의 주요한 사상은 의존성으로 주입될수 있다는 것이다.
- 객체들이 서로 많은 관계를 맺을 수 있는 것을 의미한다.
- 객체의 인스턴스를 `연결(Wiring Up)`하는 기능은 대부분 Nest 런타임 시스템에 위임될 수 있다

- 이전 장에서, 간단한 CatsController를 작성해보았다.
- Controller는 Http 요청을 처리하고 복잡한 작업을 <b>Provider</b>에게 위임할 수 있다.
- Provider는 모듈에서 `provider`라고 정의되는 평범한 자바스크립트 클래스이다.

### HINT

- Nest는 보다 객체 지향적인 방법으로 의존관계를 설계하고 조직하는 것을 가는ㅇ하게 한다.
- 우리는 `SOLID` 원칙을 따르는 것을 적극 추천한다.

### Services

- 간단하게 CatsService를 만들어보자.
- 서비스는 데이터를 저장하고 가져오는 책임이 있다.
- 그리고 CatsController가 사용하도록 설계된다.
- 그리하여 Provider의 후보가 되는것은 적절하다

```Typescript
import { Injectable } from "@nestjs/common";
import { Cat } from "./interface/cat.interface";

@Injectable()
export class CatsService {
    private readonly cats: Cat[] = [];

    create(cat: Cat) {
        this.cats.push(cat);
    }

    findAll(): Cat[] {
        return this.cats;
    }
}
```

- CatService는 하나의 멤버변수와 두 개의 메소드를 가진 클래스이다.
- 새로운 특징은 `@Injectable()` 데코레이터를 사용했다는 것이다.
- `Injectable()` 데코레이터는 메타데이터에 붙일수 있고, `CatsService`가 Neest IoC Container에 의해 관리되는 것을 선언한다.
- 그나저나 `Cat` 인터페이스는 아래와 같이 생겼다

```Typescript
export interface Cat {
    name: string;
    age: number;
    breed: string;
}
```

- 이제 우리는 고양이를 가져오는 서비스 클래스가 있다. `CatsController`에서 사용해보자.

```Typescript
import { Controller, Get, Post, Body } from '@nestjs/common';
import { CreateCatDto } from './dto/create-cat.dto';
import { CatsService } from './cats.service';
import { Cat } from './interfaces/cat.interface';

@Controller('cats')
export class CatsController {
  constructor(private catsService: CatsService) {}

  @Post()
  async create(@Body() createCatDto: CreateCatDto) {
    this.catsService.create(createCatDto);
  }

  @Get()
  async findAll(): Promise<Cat[]> {
    return this.catsService.findAll();
  }
}
```

- `CatService`는 클래스 생성자를 통해 주입된다.
- `private` 문법을 사용한 것을 주의하자.
- 이 축약문법은 `catService` 변수를 즉시 선언하고 초기화 할 수 있다.

### 의존성 주입

- Nest는 익히 알려진 의존성 주입 패턴을 전반적으로 많이 사용한다.
- 우리는 이 개념에 대해 좋은 글을 추천해주고 싶다. (Angular 문서)
- https://angular.io/guide/dependency-injection

- Nest에서는 타입스크립트의 능력덕분에 타입으로 의존성을 관리하는 것이 매우 쉽다.
- 아래의 예처럼 Nest는 `catsService`를 `CatsService`의 인스턴스로 만들고 반환한다.
- 일반적으로 싱글톤 인스턴스 이며 다른곳에서 존재하던 인스턴스를 반환한다.
- 이 의존성은 우리의 컨트롤러 생성자를 통해 전달된다.

```Typescript
constructor(private catsService: CatsService) {}
```
