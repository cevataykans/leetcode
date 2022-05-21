# Write your MySQL query statement below
SELECT
    P.product_id,
    P.product_name
FROM
    Product P JOIN ( SELECT product_id, MIN(sale_date) as min_date, MAX(sale_date) as max_date FROM Sales GROUP BY product_id) as S ON P.product_id = S.product_id
WHERE
    max_date <= "2019-03-31" AND min_date >= "2019-01-01"