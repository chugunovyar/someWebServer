create table ARTICLES
(
    id integer PRIMARY KEY GENERATED BY DEFAULT AS IDENTITY,
    headline varchar(256) not null,
    content text not null,
    pub_date TIMESTAMP
)