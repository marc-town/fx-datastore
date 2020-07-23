---- databse ----
USE fx_datastore;

---- drop ----
DROP TABLE IF EXISTS `users`;
DROP TABLE IF EXISTS `transactions`;

---- create ----
create table IF not exists `users`
(
  `id` int not null auto_increment 'ユーザーID',
  `login_id` varchar(255) not null UNIQUE comment 'ログイン用ID',
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
  `currency_pair` not null comment '通貨ペア',
  `entry_point` not null decimal(9, 5) comment 'エントリー価格',
  `entry_time` not null char(19) comment 'エントリー時間',
  `exit_point` not null decimal(9, 5) comment 'クローズ価格',
  `exit_time` not null char(19) comment 'クローズ時間',
  `lot` not null default 0 decimal(5, 2) comment 'ロット',
  `spread` decimal(3, 2) comment 'スプレッド',
  `commission` int comment '手数料',
  `swap` int default 0 comment 'スワップ',
  `create_at` timestamp not null default current_timestamp comment '登録日時',
  `update_at` timestamp not null default current_timestamp on update current_timestamp comment '更新日時',
  PRIMARY KEY (`id`)
) DEFAULT CHARSET=utf8 COLLATE=utf8_bin;