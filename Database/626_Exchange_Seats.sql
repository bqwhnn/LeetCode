-- 小美是一所中学的信息科技老师，她有一张 seat 座位表，平时用来储存学生名字和与他们相对应的座位 id。

-- 其中纵列的 id 是连续递增的

-- 小美想改变相邻俩学生的座位。

-- 你能不能帮她写一个 SQL query 来输出小美想要的结果呢？

-- 示例：

-- +---------+---------+
-- |    id   | student |
-- +---------+---------+
-- |    1    | Abbot   |
-- |    2    | Doris   |
-- |    3    | Emerson |
-- |    4    | Green   |
-- |    5    | Jeames  |
-- +---------+---------+
-- 假如数据输入的是上表，则输出结果如下：

-- +---------+---------+
-- |    id   | student |
-- +---------+---------+
-- |    1    | Doris   |
-- |    2    | Abbot   |
-- |    3    | Green   |
-- |    4    | Emerson |
-- |    5    | Jeames  |
-- +---------+---------+
-- 注意：

-- 如果学生人数是奇数，则不需要改变最后一个同学的座位。

-- sql 架构
Create table seat(id int, student varchar(255))
Truncate table seat
insert into seat (id, student) values ('1', 'Abbot')
insert into seat (id, student) values ('2', 'Doris')
insert into seat (id, student) values ('3', 'Emerson')
insert into seat (id, student) values ('4', 'Green')
insert into seat (id, student) values ('5', 'Jeames')

-- 转换思路，交换 id 而不是交换 student
-- answer
SELECT (CASE
            WHEN MOD(id,2) = 1 AND id = (SELECT COUNT(*) FROM seat) THEN id
            WHEN MOD(id,2) = 1 THEN id+1
            ELSE id-1
       END) AS id, student
FROM seat
ORDER BY id;