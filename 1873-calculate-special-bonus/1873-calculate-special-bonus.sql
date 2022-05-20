# Write your MySQL query statement below
SELECT
    employee_id,
    IF ( employee_id % 2 != 0, IF ( name IS NOT NULL AND LEFT( name, 1) != 'M', salary, 0), 0 ) AS bonus
FROM
    Employees