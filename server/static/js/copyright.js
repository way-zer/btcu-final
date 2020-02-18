$(document).ready(function () {
    //注册
    $("register-from").validate({
        rules: {
            username: {
                required: true,
            },
            password: {
                required: true
            },
            repassword: {
                required: true,
                equalTo: "#register-password"
            }
        },
        messages: {
            username: {
                required: "请输入用户名"
            },
            password: {
                required: "请输入密码"
            },
            repassword: {
                required: "请确认密码",
                equalTo: "两次输入的密码必须相等"
            }
        },
        submitHandler: function (form) {
            var urlStr = "/register";
            $(form).ajaxSubmit({
                url: urlStr,
                type: "post",
                dataType: "json",
                success: function (data, status) {
                    alert("data:" + data.message)
                    if (data.code == 1) {
                        setTimeout(function () {
                            window.location.href = "/login"
                        }, 1000)
                    }
                },
                err: function (data, status) {
                    alert("err:" + data.message + ":" + status)
                }
            })
        }
    });


    //登录
    $("#login-form").validate({
        rules: {
            username: {
                required: true
            },
            password: {
                required: true
            }
        },
        messages: {
            username: {
                required: "请输入用户名"
            },
            password: {
                required: "请输入密码"
            }
        },
        submitHandler: function (form) {
            var urlStr = "/login"
            $(form).ajaxSubmit({
                url: urlStr,
                type: "post",
                dataType: "json",
                success: function (data, status) {
                    alert("data:" + data.message + ":" + status)
                    if (data.code == 1) {
                        setTimeout(function () {
                            window.location.href = "/"
                        }, 1000)
                    }
                },
                error: function (data, status) {
                    alert("err:" + data.message + ":" + status)

                }
            });
        }
    });

    // 添加版权信息的表单
    $("add-copyright-form").validate({
        rules: {
            name: "required",
            author: "required",
            press: "required",
            privateKey: "required",
        },
        messages: {
            name: "请输入作品名",
            author: "请输入作者名",
            press: "请输入出版社名",
            privateKey:"请输入用户私钥",
        },
        submitHandler: function (form) {
            var urlStr = "/copyright/add";
            //判断版权id确定提交的表单的服务器地址
            //若id大于零，说明是修改版权
            var Id = $("#add-copyright-id").val();
            alert("copyrightId:" + Id);
            if (Id > 0) {
                urlStr = "/copyright/update"
            }
            alert("urlStr:" + urlStr);

            $(form).ajaxSubmit({
                url: urlStr,
                type: "post",
                dataType: "json",
                success: function (data, status) {
                    alert(":data:" + data.message);
                    setTimeout(function () {
                        window.location.href = "/"
                    }, 1000)
                },
                error: function (data, status) {
                    alert("err:" + data.message + ":" + status)
                }
            });
        }
    });

    // 上传文件
    $("#document-upload-button").click(function () {
        var filedata = $("#document-upload-file").val();
        if (filedata.length <= 0) {
            alert("请选择文件!");
            return
        }
        //文件上传通过Formdata去储存文件的数据
        var data = new FormData()
        data.append("upload", $("#document-upload-file")[0].files[0]);
        alert(data)
        var urlStr = "/upload"
        $.ajax({
            url: urlStr,
            type: "post",
            dataType: "json",
            contentType: false,
            data: data,
            processData: false,
            success: function (data, status) {
                alert(":data:" + data.message);
                if (data.code == 1) {
                    $("#document-hash").val(data.hash)
                } else {
                    $("#document-hash").val("请重新上传文件")
                }
            },
            error: function (data, status) {
                alert("err:" + data.message + ":" + status )
            }
        })
    })

});