# Write your MySQL query statement below
SELECT
    MAX(salary) AS SecondHighestSalary
FROM
    (
    SELECT
        salary,
        DENSE_RANK() OVER( ORDER BY salary DESC) as row_num
    FROM
        Employee
    ) AS temp
WHERE
    row_num = 2