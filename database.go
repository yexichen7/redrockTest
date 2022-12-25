mysql> desc users;
+----------+-------------+------+-----+---------+----------------+
| Field    | Type        | Null | Key | Default | Extra          |
+----------+-------------+------+-----+---------+----------------+
| id       | bigint      | NO   | PRI | NULL    | auto_increment |
| name     | varchar(20) | YES  |     |         |                |
| password | varchar(20) | YES  |     |         |                |
+----------+-------------+------+-----+---------+----------------+

mysql> desc borrowrecords;
+--------+-------------+------+-----+---------+-------+
| Field  | Type        | Null | Key | Default | Extra |
+--------+-------------+------+-----+---------+-------+
| name   | varchar(20) | YES  |     |         |       |
| states | varchar(20) | YES  |     |         |       |
+--------+-------------+------+-----+---------+-------+

mysql> desc  booklibrary;
+----------+-------------+------+-----+---------+----------------+
| Field    | Type        | Null | Key | Default | Extra          |
+----------+-------------+------+-----+---------+----------------+
| id       | bigint      | NO   | PRI | NULL    | auto_increment |
| name     | varchar(20) | YES  |     |         |                |
| states   | varchar(20) | YES  |     |         |                |
| quantity | int         | YES  |     | 0       |                |
+----------+-------------+------+-----+---------+----------------+

预处理语句减少对MySQL的频繁请求，使得服务器高效运行，可以多次使用，并且效率更加高效

