---- databse ----
USE fx_datastore;

---- drop ----
DROP TABLE IF EXISTS `users`;
DROP TABLE IF EXISTS `transactions`;

---- create ----
create table IF not exists `users`
(
  `id` int not null auto_increment comment 'ユーザーID',
  `login_id` varchar(255) not null UNIQUE comment 'ログインID',
  `last_name` varchar(255) comment '姓',
  `first_name` varchar(255) comment '名',
  `password` varchar(50) not null default 'Pa55w0rd',
  `mail_address` varchar(255) UNIQUE comment 'メールアドレス',
  `create_at` timestamp not null default current_timestamp comment '登録日時',
  `update_at` timestamp not null default current_timestamp on update current_timestamp comment '更新日時',
  PRIMARY KEY (`id`)
) DEFAULT CHARSET=utf8 COLLATE=utf8_bin;

create table IF not exists `transactions`
(
  `id` int not null auto_increment comment '取引ID',
  `currency_pair` char(7) not null comment '通貨ペア',
  `entry_point` decimal(9, 5) not null comment 'エントリー価格',
  `entry_time` char(19) not null comment 'エントリー時間',
  `exit_point` decimal(9, 5) not null comment 'クローズ価格',
  `exit_time` char(19) not null comment 'クローズ時間',
  `lot` decimal(5, 2) not null default 0 comment 'ロット',
  `spread` decimal(3, 2) comment 'スプレッド',
  `commission` int comment '手数料',
  `swap` int default 0 comment 'スワップ',
  `create_at` timestamp not null default current_timestamp comment '登録日時',
  `update_at` timestamp not null default current_timestamp on update current_timestamp comment '更新日時',
  PRIMARY KEY (`id`)
) DEFAULT CHARSET=utf8 COLLATE=utf8_bin;