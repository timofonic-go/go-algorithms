# 14.2 Open Requests
# Write a SQL query to get a list of all buildings and the number of open requests (Requests in which status equals 'Open').

SELECT
  BuildingName,
  IFNULL(ReqCounts.Count, 0) AS Count
FROM Buildings
  LEFT JOIN
    (SELECT Apartments.BuildingID, count(*) AS Count
    FROM Requests INNER JOIN Apartments
    ON Requests.AptID = Apartments.AptID
    WHERE Requests.Status = 'Open'
    GROUP BY Apartments.BuildingID) ReqCounts

  ON ReqCounts.BuildingID = Buildings.BuildingID;