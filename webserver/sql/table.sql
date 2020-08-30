CREATE TABLE IF NOT EXISTS `user_tb`(
   `id` INT UNSIGNED AUTO_INCREMENT,
   `userno` VARCHAR(100) NOT NULL,
   `username` VARCHAR(40) NOT NULL,
   `userpassword` VARCHAR(40) NOT NULL,
   `createtime` DATETIME,
   PRIMARY KEY ( `id` )
)ENGINE=InnoDB DEFAULT CHARSET=utf8;

insert into user_tb(userno,username,userpassword,createtime)values('1001','yuwei','123456',now());
insert into user_tb(userno,username,userpassword,createtime)values('1002','gongwenming','123456',now());