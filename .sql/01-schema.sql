create DATABASE dev 
    with owner admin;

\connect dev; 

CREATE EXTENSION pgcrypto;

CREATE SCHEMA helloworld AUTHORIZATION admin;

CREATE table helloworld.person(
    Id UUID PRIMARY KEY DEFAULT gen_random_uuid(), 
    Age INTEGER, 
    Person VARCHAR(128), 
    FavoriteFood VARCHAR(128)
);

GRANT all privileges 
    on all tables
    in SCHEMA helloworld
    to admin;
