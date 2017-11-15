<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>注册-{{config "String" "globaltitle" ""}}</title>
    {{template "inc/meta.tpl" .}}
</head>
<body>
<div class="container">
    <form class="form-signin" id="login-form">
        <h2 class="form-signin-heading">请填入你的注册信息</h2>
        <label for="inputEmail" class="sr-only">Email</label>
        <input type="tel" name="email" class="form-control" placeholder="email" required autofocus>
        <label for="inputPassword" class="sr-only">密码</label>
        <input type="password" name="password" class="form-control" placeholder="密码" required>
        <label for="inputPasswordConform" class="sr-only">再次输入密码</label>
        <input type="password" name="password_conform" class="form-control" placeholder="再次输入密码" required>
        <button class="btn btn-lg btn-primary btn-block" type="submit">提交</button>
    </form>
</div>
{{template "inc/foot.tpl" .}}
</body>
</html>

