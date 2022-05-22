# Write your MySQL query statement below

    SELECT
        E.employee_id
    FROM
        Employees E
    WHERE
        E.employee_id NOT IN (
            SELECT
                S.employee_id
            FROM
                Salaries S
        )
    
    UNION
    
    SELECT
        S.employee_id
    FROM
        Salaries S
    WHERE
        S.employee_id NOT IN (
            SELECT
                E.employee_id
            FROM
                Employees E
        )
ORDER BY
    employee_id ASC