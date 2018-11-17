<html>
<head>
    <title></title>
</head>
<body>
<form action="http://127.0.0.1:9090/login" method="post">
    用户名:<input type="text" name="username"><br>
    密码:<input type="password" name="password"><br>
    年龄:<input type="text" name = "age"><br>

    <select name = "login_method">
        <option vaule="passwd" selected>用户名+密码</option>
        <option value="digitalCert">数字证书</option>
        <option value="Ukey" >USB KEY</option>
    </select>
    <br>
    <input type="radio" name="gender" value="1">男 <br>
    <input type="radio" name="gender" value="2">女<br>

    <input type="checkbox" name="interest" value="basketball">篮球<br>
    <input type="checkbox" name="interest" value="football">足球<br>
    <input type="checkbox" name="interest" value="baseball">排球<br>

    <input type="hidden" name="token" value="{{.}}}">


    <input type="submit" value="登陆">



</form>
</body>
</html>