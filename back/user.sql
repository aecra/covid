create table if not exists users(
    name varchar(64) not null primary key,
    email varchar(64),
    position varchar(8) default 'school',
    state varchar(8) default 'on',
    eaisess varchar(32),
    uukey varchar(32),
    home varchar(65533)
);

create table if not exists record(
  id int not null primary key auto_increment,
  time datetime default current_time,
  name varchar(64) not null,
  email varchar(64) not null,
  position varchar(8) default 'school',
  result varchar(255) default ''
);