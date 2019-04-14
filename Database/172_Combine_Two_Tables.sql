-- 表1: Person

-- +-------------+---------+
-- | 列名         | 类型     |
-- +-------------+---------+
-- | PersonId    | int     |
-- | FirstName   | varchar |
-- | LastName    | varchar |
-- +-------------+---------+
-- PersonId 是上表主键
-- 表2: Address

-- +-------------+---------+
-- | 列名         | 类型    |
-- +-------------+---------+
-- | AddressId   | int     |
-- | PersonId    | int     |
-- | City        | varchar |
-- | State       | varchar |
-- +-------------+---------+
-- AddressId 是上表主键
 
-- 编写一个 SQL 查询，满足条件：无论 person 是否有地址信息，都需要基于上述两表提供 person 的以下信息：

-- FirstName, LastName, City, State

-- sql 架构
Create table Person (PersonId int, FirstName varchar(255), LastName varchar(255))
Create table Address (AddressId int, PersonId int, City varchar(255), State varchar(255))
Truncate table Person
insert into Person (PersonId, LastName, FirstName) values ('1', 'Wang', 'Allen')
Truncate table Address
insert into Address (AddressId, PersonId, City, State) values ('1', '2', 'New York City', 'New York')

-- 数据库在通过连接两张或多张表来返回记录时，都会生成一张中间的临时表，然后再将这张临时表返回给用户。
-- 在使用 left join 时，on 和 where 条件的区别：
-- on 是在生成临时表时使用的条件，不管 on 中的条件是否为真，都会返回左边表中的记录。
-- where 时在临时表生成好后，再对临时表进行过滤的条件。这时候不会必定返回左边表的记录，凡是条件不为真的都被过滤了。
-- answer
SELECT a.FirstName, a.LastName, b.City, b.State
FROM Person a LEFT JOIN Address b
ON a.PersonId = b.PersonId;