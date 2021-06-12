drop table if exists `likes`;
drop table if exists `posts`;
drop table if exists `users`;

create table `users` (
    `user_id` int unsigned auto_increment,
    `name` varchar(20) not null,
    `create_at` datetime not null,
    `update_at` datetime not null,
    primary key (`user_id`)
) engine=innodb default charset=utf8;

create table `posts` (
    `post_id` int unsigned auto_increment,
    `title` varchar(20) not null,
    `body` varchar(255) not null,
    `author_id` int unsigned not null,
    `create_at` datetime not null,
    `update_at` datetime not null,
    primary key (`post_id`),
    foreign key (`author_id`) references users(`user_id`) on update cascade

) engine=innodb default charset=utf8;

create table `likes` (
    `like_id` int unsigned auto_increment,
    `post_id` int unsigned not null,
    `user_id` int unsigned not null,
    `create_at` datetime not null,
    primary key (`like_id`),
    foreign key (`post_id`) references posts(`post_id`) on update cascade,
    foreign key (`user_id`) references users(`user_id`) on update cascade
) engine=innodb default charset=utf8;

insert into users (name, create_at, update_at) values ('a', '2021-06-12', '2021-06-12');
insert into users (name, create_at, update_at) values ('b', '2021-06-12', '2021-06-12');
insert into posts (title, body, author_id, create_at, update_at) values ('a title', 'a content', (select user_id from users where user_id = 1), '2021-06-12', '2021-06-12');
-- insert into likes (post_id, user_id, create_at) values ((select post_id from posts where post_id = 1),(select user_id from users where user_id = 1), '2021-06-12');
-- insert into likes (post_id, user_id, create_at) values ((select post_id from posts where post_id = 1),(select user_id from users where user_id = 2), '2021-06-12');
