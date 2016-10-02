# Implement a query to get a list of all teachers and how many students they each teach. If a teacher teaches
# the same student in two courses, you should double count the student. Sort the list in descending order of
# the number of students a teacher teaches.
# We can construct this query step by step. First, let's get a list of TeacherIDs and how many students are
# associated with each TeacherID. This is very similar to the earlier query.

SELECT
 Teachers.TeacherID,
 count(StudentCourses.StudentID) AS StudentCount
FROM
 Teachers
 LEFT JOIN Courses ON Teachers.TeacherID = Courses.TeacherID
 LEFT JOIN StudentCourses ON StudentCourses.CourseID = Courses.CourseID
GROUP BY
 Teachers.TeacherID;

# Note that this INNER JOIN will not select teachers who aren't teaching classes. We'll handle that in the
# below query when we join it with the list of all teachers.

SELECT
 Teachers.TeacherName,
 StudentSize.StudentCount
FROM Teachers
 LEFT JOIN (
            SELECT
             Teachers.TeacherID,
             count(StudentCourses.StudentID) AS StudentCount
            FROM
             Teachers
             LEFT JOIN Courses ON Teachers.TeacherID = Courses.TeacherID
             LEFT JOIN StudentCourses ON StudentCourses.CourseID = Courses.CourseID
            GROUP BY
             Teachers.TeacherID
           ) StudentSize
  ON Teachers.TeacherID = StudentSize.TeacherID
ORDER BY StudentSize.StudentCount DESC;

# Note how we handled the NULL values in the SELECT statement to convert the NULL values to zeros.