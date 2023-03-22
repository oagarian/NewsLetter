# NewsLetter

## How to run the server:
```ssh
cd server/src

go run ./
```

## MySQL code:

```ssh
CREATE DATABASE NEWSLETTER;
USE NEWSLETTER;

CREATE TABLE REGISTRED_USERS (
    ID INTEGER NOT NULL PRIMARY KEY,
    EMAIL VARCHAR(40) NOT NULL
);
```

## ToDo:

* Create email sender;
