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
    primary key (`like_id`),
    foreign key (`post_id`) references posts(`post_id`) on update cascade,
    foreign key (`user_id`) references users(`user_id`) on update cascade
) engine=innodb default charset=utf8;

-- SELECT Count(Recipes.recipe_id)
-- FROM Recipes LEFT JOIN Reports
-- ON Recipes.recipe_id = Reports.recipe_id
-- WHERE Reports.recipe_id = NULL
