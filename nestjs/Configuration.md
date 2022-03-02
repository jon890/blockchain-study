# NestJS Techniques Configuration

- 네스트 공식문서 설정을 관리하는 효율적인 방법을 해석해보며 공부해보았다.
- 스프링에서 환경을 설정하는 것과 크게 차이는 없는듯하지만 Node.js에서 환경을 관리하는 방식은 공부를 해야할 듯
- https://docs.nestjs.com/techniques/configuration

## Configuration (설정)

- 애플리케이션은 종종 다른 환경에서 실행된다. (로컬, 개발, 스테이징, 운영)
- 환경에따라, 다른 설정 구성을 사용해야한다.
- 예를 들어, 로컬 환경은 로컬 DB 인스턴스에만 유효한 특정 데이터베이스 자격 증명에 의존한다.
- 운영 환경은 다른 DB 자격 증명이 사용되어야 한다.
- 설정이 다양하게 바뀜에 따라, 제일 좋은 방법은 environment에 설정을 저장하는 것이다.

- 외부에 정의되는 환경 변수는 Node.js에서는 proess.env를 통해 전역적으로 보인다. (process.env.XXX)
- 우리는 각각 환경에에서 환경 변수를 별도로 설정하여 다중 환경 문제를 해결하려고 할 수 있다.
- 이는 특히 이러한 값을 쉽게 변경해야 하는 개발 및 테스트 환경에서 빠르게 다루기 어려워질수 있다.

- Node.js 애플리케이션에서는 흔히 특정값을 키-값 쌍으로 가지고 있고, 각각의 환경을 나타내는 `.env` 파일이 사용된다.
- 다른 환경에서 애플리케이션을 구동할 때에는 단지 적절한 `.env` 파일로 변경하면 그만이다.

- 네스트에서 이러한 테크닉을 사용하는 좋은 접근법은 `ConfigModule`을 생성하고 적절한 `.env` 파일을 불러오는 `ConfigService`를 노출하면 된다.
- `@nestjs/config` 패키지의 도움을 받거나, 직접 작성할 수 있다.
- 이 챕터에서는 이 패키지를 다루는 방법에 대해 다룬다.

### 설치

```TypeScript
$ npm i --save @nestjs/config
```

- HINT : @nestjs/config 패키지는 내부적으로 `dotenv`를 사용한다.
- NOTE : @nestjs/config는 TypeScript 4.1 이상 버전이 필요하다.

### 시작하기

- 설치가 완료되었다면 `ConfigModule`을 import 할 수 있다.
- 일반적으로 `AppModule`에 import하여 행위를 정적 메서드 `.forRoot()`를 사용하여 제어한다.
- 이 단계를 거치면, 환경 변수의 키-값 쌍이 해석되고 resolve(결정?)된다.
- 이후에 우리는 다른 기능적인 모듈에서 `ConfigModule`의 클래스인 `ConfigService`를 접근하는 여러 방법을 알아볼 것이다.

```TypeScript
app.module.ts

import { module } from '@nestjs/common';
import { ConfigMoudle } from '@nestjs/config';

@Module({
    imports: [ConfigModule.forRoot()],
})
export class AppModule {}
```

- 위 코드는 기본 위치로부터(프로젝트 루트 폴더) `.env` 파일을 불러오고 해석한다, `env` 파일의 키-값 쌍은 환경변수와 함께 process.env에 합쳐져 할당된다, 그리고 `ConfigService`를 통해 접근할 수 있는 private한 구조에 결과를 저장한다.
- `forRoot()` 메서드는 `get()` 메서드를 이용하여 설정 변수를 읽고/합칠 수 있는 `ConfigService` provider를 등록한다.
- `@nestjs/config`가 `dotenv`에 의존하기 떄문에 환경 변수 이름의 중복을 해결할 떄는 `dotenv`의 룰을 따른다.
- 런타임 환경과 환경 변수에서 키가 존재한다면 (OS Shell 을 통하여 export DATABASE_USER=test, .env 파일에도 선언), 런타임 환경 변수가 우선순위를 갖는다.

### 사용자 지정 env 파일 경로

- `@nestjs/config` 패키지는 기본적으로 `.env` 파일을 애플리케이션 루트 디렉터리부터 찾는다.
- `.env`파일의 경로를 다른 곳으로 지정하려면 `envFilePath` 속성을 `forRoot()`의 속성 객체에 다음과 같이 넘겨준다.

```TypeScript
ConfigModule.forRoot({
    envFilePath: '.development.env',
})
```

- 또한 우리는 `.env` 파일의 경로를 여러 곳으로 설정할 수 있다.

```TypeScript
ConfigModule.forRoot({
    envFilePath: ['.development.env', '.local.env']
})
```

- 여러 개의 파일이 발견되면, 첫 번째것이 우선순위를 갖는다.

### env 변수 로딩 끄기

- `.env` 파일을 로드하길 원하지 않는다면, 대신에 간단히 런타임 환경으로부터 환경 변수를 넘기든, 옵션 객체에 `ignoreEnvFile` 속성을 다음과 같이 `true`로 설정한다.

```TypeScript
ConfigModule.forRoot({
    ignoreEnvFile: true,
})
```

### 전역으로 모듈 사용하기

- `ConfigModule`을 다른 모듈에서도 사용하길 원한다면, 네스트 기본 모듈방식에 따라 import할 필요가 있다.
- 대안적으로 옵션 객체에 `isGlobal` 속성을 `true`로 설정할 수 있다.
- 이 경우에 다른 모듈에서 `ConfigModule`을 import할 필요가 없다.
- AppModule 같은 루트 모듈에서 한번 로드되어진다.

```TypeScript
ConfigModule.forRoot({
    isGlobal: true,
})
```

### 사용자 지정 환경 파일 목록

- 보다 복잡한 프로젝트를 위해, 우리는 중첩된 설정 객체를 반환하는 커스텀 설정파일을 활용할 수 있다.
- 이 것은 함수를 사용함으로써 설정파일을 그룹화 할 수 있고, 관련된 설정을 개별파일에 저장함으로써 독립적으로 관리할 수 있게 한다.

- 커스텀 설정 파일은 설정 객체를 반환하는 팩토리 함수를 export 한다.
- 설정 객체는 아무 일반 자바스크립트 객체가 될 수 있다.
- `process.env` 객체는 완전히 읽어진 키-값 환경 변수쌍을 포함할 것이다 (`.env` 파일과 함께 외부에 정의된 변수들이 아래와 같이 resolve and merge 된다(한글로 해석하기가 힘듬..))
- 반환된 설정 객체를 통하여, 적절한 타입으로 변환하던지, 기본 값을 설정하든지, 필요한 로직을 추가할 수 있다.

```TypeScript
export default () => ({
    port: parseInt(process.env.PORT, 10) || 3000,
    databaser: {
        host: process.env.DATABASE_HOST,
        port: parseInt(process.env.DATABASE_PORT, 10) || 5432
    }
});
```

- 우리는 옵션 객체의 `load` 속성을 이용하여 `ConfigModule.forRoot()` 메서드로 넘길 수 있다.

```TypeScript
import configuration from './config/configuration';

@Module({
    imports: [
        ConfigModule.forRoot({
            load: [configuration],
        }),
    ],
});
export class AppModule {}
```

- NOTICE : `load` 속성에 할당되는 값은 array이다, 다양한 설정 파일을 읽는 것이 허용된다.

- 커스텀 설정 파일들과, 우리는 파일들을 YAML 파일들로 관리할 수 있다.
- 아래의 예시는 YAML 포맷을 사용하여 설정파일을 구성한 예이다.

```YAML
http:
  host: 'localhost'
  port: 8080

db:
  postgres:
    url: 'localhost'
    port: 5432
    database: 'yaml-db'

  sqlite:
    database: 'sqlite.db'
```

- YAML 파일들을 읽고 해석하기 위해서는, `js-yaml` 패키지를 활용할 수 있다.

```TypeScript
$ npm i js-yaml
$ npm i -D @types/js-yaml
```

- 패키지가 설치되면, `yaml#load` 함수를 통해 위에서 만든 YAML 파일을 읽을 수 있다.

```TypeScript
config/configuration.ts

import { readFileSync } from 'fs';
import * as yaml from 'js-yaml';
import { join } from 'path';

const YAML_CONFIG_FILENAME = 'config.yaml';

export default () => {
    return yaml.load(
        readFileSync(join(__dirname, YAML_CONFIG_FILENAME), 'utf8'),
    ) as Record<string, any>;
};
```

- NOTE : 네스트 CLI는 우리의 "assets"을 빌드 과정에서 자동으로 "dist" 폴더에 옮겨 주지 않는다.
  YAML 파일들이 복사된 것을 확실히 하기 위해, `nest-cli.json` 파일에 `compilerOptions#assets` 옵션을 지정해라.
- 예를 들어, `config` 폴더가 `src` 폴더와 같은 레벨에 있다면, `compilerOption#assets`에 `"assets": [{"include": "../config/*.yaml", "outDir": "./dist/config"}"]` 추가해라.
- 더 자세히 알아보려면 https://docs.nestjs.com/cli/monorepo#assets 를 참조하자.

### ConfigService 사용하기

- `ConfigService`의 설정 값에 접근하기 위해서, 먼저 `ConfigService`를 inject(주입) 해야 한다.
- 다른 Provider와 마찬 가지로, 우리는 `ConfigSerivce`를 포함하는 모듈을 import 해야한다. (`ConfigModule`)
- `ConfigModule.forRoot()`에서 `isGlobal` 속성을 `true`로 지정했다면 안해도 된다.
- 아래와 같이 feature 모듈에서 import한다.

```TypeScript
feature.module.ts

@Moudle({
    imports: [ConfigModule],
});
```

- 그러면 우리는 표준 생성자 주입을 사용하여 주입 받을 수 있다.

```TypeScript
constructor(private configService: ConfigService) {}
```

- HINT : `ConfigService`는 `@nestjs/config` 패키지로 부터 import 된다.
- 그리고 클래스에서는 다음과 같이 사용 가능하다.

```TypeScript
// get and environment variable
const dbUser = this.configService.get<string>('DATABASE_USER');

// get a custom configuration value
const dbHost = this.configService.get<string>('database.host');
```

- 위와 같이 `configService.get()` 메서드를 사용하여 간단한 환경 변수를 이름으로 가져올 수 있다.
- 타입을 넘겨줌으로써 타입스크립트의 타입 힌트를 사용할 수 있다. (예 `get<string>(...)`)
- 또한 `get()` 메서드는 연결된 커스텀 설정 객체를 순회할수 있다.
- 또한 인터페이스를 타입 힌트로 사용함으로써 커스텀 설정 객체 전체를 얻어 올 수 있다.

```TypeScript
interface DatabaseConfig {
    host: string;
    port: number;
}

const dbConfig = this.configService.get<DatabaseConfig>('database');

// you can now use `dbConfig.port` and `dbConfig.host`
const port = dbConfig.port;
```

- 또한 `get()` 메서드는 2번째 매개변수로 기본값을 받을 수 있다, 키가 존재하지 않는다면 기본값을 반환한다. 아래의 예를 참조하자.

```TypeScript
// use "localhost" when "database.host" is not defined
const dbHost = this.configService.get<string>('datbase.host', 'localhost');
```

- `ConfigService`는 두개의 선택적 제너릭 (타입 변수)가 있다.
- 첫 번째는 존재하지 않는 설정 속성을 접근하는 것을 방지한다. 아래의 예를 보자

```TypeScript
interface EnvironmentVariables {
    PORT: number;
    TIMEOUT: string;
}

// somewhere in the code
constructor(private configService: ConfigService<EnvironmentVariables>) {
    const port = this.configService.get('PORT', { infer: true });

    // TypeScript Error: this is invalid as the URL property is not defined in EnvironmentVariables
    const url = this.configService.get('URL', { infer: true });
}
```

- `infer` 속성을 `true`로 사용하면, `ConfigService#get` 메서드는 자동으로 인터페이스의 타입을 추론한다. 예를 들어, `typeof port === "number"`, `EnvironmentVariables` 인터페이스의 `PORT`는 `number` 이기 떄문이다.

- 또한 `infer` 기능을 이용하여, 다음과 같이 심지어 . 표기법을 사용하지 않아도 커스텀 설정 객체의 속성에 연관된 타입을 추론할 수 있다.

```TypeScript
constructor(private configService: ConfigService<{ database: { host: string } }>) {
    const dbHost = this.configService.get('database.host', { infer: true })!;
    // typeof dbHost === "string"
    // ! --> non-null assertion operator
}
```

- 두번째 제네릭은 첫번째에 의존한다. `ConfigService`의 메서드가 `strictNulChecks`를 반환하면 모든 `undefined` 타입 선언을 제거 한다.(??)
- 예를들어

```TypeScript
// ...
constructor(private configService: ConfigService<{ PORT: number } , true>) {
    //
    const port = this.configService.get('PORT', { infer: true });
    //    ^^^ The type of port will be 'number' thus you don't need TS type assertions anymore
}
```

### 설정 네임스페이스

- 위와 같이 `ConfigModule`d은 여러 설정 파일을 읽고 정의하는 것을 허용한다.
- 이 챕터에서는 우리는 중첩된 여러 설정 객체의 계층구조를 관리하는 법을 보여준다.
- 대안으로, 우리는 `registerAs()` 함수를 사용하여 네임스페이스된 설정 객체를 반환할 수 있다.

```TypeScript
config/database.config.ts

export default registerAs('database', () => ({
    host: process.env.DATABSE_HOST,
    port: process.env.DATABASE_PORT || 5432
}));
```

- HINT : `registerAs` 함수는 `@nestjs/config` 패키지에서 import 할 수 있다.

- 다른 커스텀 설정 파일과 동일한 방법으로, 네임스페이스된 설정은 옵션 객체의 `forRoot()` 메소드의 `load` 속성을 통해 읽어 진다.

```TypeScript
import databaseConfig from './config/database.config';

@Module({
    imports: [
        ConfigModule.forRoot({
            load: [databaseConfig],
        }),
    ],
})
export class AppModule {}
```

- 이제. 표기법을 이용해 `database` 네임스페이스에서 `host` 값을 가져올 수 있다.
- 네임스페이스의 이름과 대응하는 `'database'` 속성 이름을 전치사로 사용하자. (`registerAs()` 함수의 첫 번째 매개변수로 넘겼던 값)

```TypeScript
const dbHost = this.configService.get(<string>('database.host'));
```

- 합리적인 대안은 `database` 네임스페이스를 직접 주입하는 것 이다.
- 강한 타입의 이점을 가져다 준다.

```TypeScript
constructor(
    @Inject(databaseConfig.KEY)
    private dbConfig: ConfigType<typeof databaseConfig>,
) {}
```

- HINT : `ConfigType`은 `@nestjs/config` 패키지에서 import 할 수 있다.

### 환경 변수 캐시

- `process.env`에 접근하는 것이 느리다면, `ConfigService#get` 메서드의 `process.env`에 저장할 때 속도를 향상시키기 위해 `ConfigModule.forRoot()` 의 옵션 객체에 `cache` 속성을 설정할 수 있다.

```TypeScript
ConfigModule.forRoot({
    cache: true,
});
```

### 부분 등록 (Partial Registration)

- 지금까지, `forRoot()` 메서드를 사용하여 루트 모듈에서 설정파일을 처리하는 법을 봤다.
- 여러 다른 폴더에 각 기능에 대한 설정을 가진 더 복잡한 프로젝트 구조일 수도 있을 것이다.
- 이러한 모든 설정 파일을 루트 모듈에서 읽는 것보다 `@nestjs/config` 패키지는 부분 등록 이라는 기능을 제공한다.
- 부분 등록 : 각 기능 모듈과 연관된 설정 파일만 참조하는 것
- `forFeature()` static 메서드를 사용하여 각 기능에 대한 부분 등록을 사용해 보자.

```TypeScript
import databaseConfig from './config/database.config';

@Module({
    imports: [ConfigModule.forFeature(databaseConfig)],
})
export class DatabaseModule {}
```

- WARNING : 특정 환경에서, 우리는 생성자보다 `onModuleInit()` 훅을 통해 부분 등록된 속성에 접근할 수 있다.
- 이것은 모듈 초기화 순서 떄문에, `forFeature()` 메서드는 모듈 초기화시에 동작하기 때문이다.
- 이러한 방식으로 로드된 모듈의 값에 생성자에서 접근할 때에, 이러한 방식으로 의존된 모듈은 아직 초기화 되지 않았을 수 있다.
- `onModuleInit()` 메서드는 모든 모듈이 초기화된 후 동작하기 떄문에, 이 방법은 안전하다.

### Validation (TODO)

### `main.ts`에서 사용하기

- 우리의 설정은 service에 저장되지만, `main.ts`에서도 사용할 수 있다.
- 이 방법은 애플리케이션 포트또는 CORS host에 저장된 변수를 사용할 수 있다.
- 서비스 객체에 접근하기 위해 우리는 꼭 `app.get()` 메소드를 사용해야 한다.

```TypeScript
const configService = app.get(ConfigService);
```

- 그러면 일반적으로 `get` 메서드와 설정 키를 이용하여 호출할 수 있다.

```TypeScript
const port = configService.get('PORT');
```
