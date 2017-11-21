<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>detail</title>
    <link rel="stylesheet" type="text/css" href="assets/css/detail.css" />
</head>
<body>
    <h1>用户详情</h1>
    <table>
        <tr>
            <th>用户名</th>
            <th>电话</th>
            <th>邮箱</th>
        </tr>
        <tr>
            <td>{{.Username}}</td>
            <td>{{.Phone}}</td>
            <td>{{.Email}}</td>
        </tr>
    </table>
</body>
</html>
