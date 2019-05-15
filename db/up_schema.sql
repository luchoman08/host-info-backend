
CREATE TABLE domain_info  (
    id bigserial primary key,
    servers_chaged boolean,
    ssl_grade varchar(5),
    previous_ssl_grade varchar(5),
    logo varchar(200),
    title varchar(50),
    is_down boolean,
    logic_deleted boolean
);

CREATE TABLE server_info
(
    id        bigserial primary key,
    address   varchar(15),
    ssl_grade varchar(3),
    country   varchar(10),
    owner     varchar(50),
    created timestamp default now(),
    edited timestamp,
    domain bigserial references domain_info(id),
    logic_deleted boolean
);