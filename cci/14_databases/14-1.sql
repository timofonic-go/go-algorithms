# 14.1 Multiple Apartments: Write a SQL query to get a list of tenants who are renting more than one apartment.

# HINT: Whenever you write a GROUP BY clause in an interview (or in real life), make sure that anything in the SELECT clause is either an aggregate function or contained within the GROUP BY clause.

SELECT t.TenantName
FROM
  Tenants t INNER JOIN (SELECT TenantID
                        FROM AptTenants
                        GROUP BY TenantID
                        HAVING count(*) > 1) c
    ON t.TenantID = c.TenantID;

# SELECT StudentID
# FROM StudentCourses
# GROUP BY StudentID
# HAVING count(*) > 1;