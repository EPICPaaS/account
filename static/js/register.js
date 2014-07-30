$(document).ready(function(){
	handleSubmit();
});

    var handleSubmit = function() {
		
        $("#register_form").validate({
            rules : {
                UserName : {
                    required : true,
					minlength:5,
					maxlength:30
                },
                Email : {
                    required : true,
					email:true
                },
                CompanyName : {
                    required : true
                },
				QQ : {
                    required : true
                },
				TelNum : {
                    required : true
                },
				Password : {
                    required : true,
					minlength: 5
					
                },
				PasswordRe : {
                    required : true,
					minlength: 5,
                    equalTo: "#RegisterForm-Password"
                },
				Captcha : {
                    required : true
                },
				CompanyName : {
                    required : true
                }
            },
            messages : {
                UserName : {
                    required : "用户名不能为空",
					minlength:"最小长度为5",
					maxlength:"最大长度为30"
                },
                Email : {
                    required : "邮箱不能为空",
					email:"邮箱格式错误"
                },
                CompanyName : {
                    required : "公司名称不能为空"
                },
				QQ : {
                    required : "QQ不能为空"
                },
				TelNum : {
                    required : "电话不能为空"
                },
				Password : {
                    required : "密码不能为空",
					minlength: "密码长度不能小于5"
					
                },
				PasswordRe : {
                    required : "重复密码不能为空",
					minlength: "重复密码长度不能小于5",
                    equalTo: "两次输入密码不一致"
                },
				Captcha : {
                    required : "验证码不能为空"
                }
            },
			 highlight : function(element) {  
                $(element).closest('.form-group').addClass('has-error');  
            },  
  
            success : function(label) {  
                label.closest('.form-group').removeClass('has-error');  
                label.remove();  
            },  
  
            errorPlacement : function(error, element) {  
                element.parent('div').append(error);  
            },  
  
            submitHandler : function(form) {  
                form.submit();  
            }  
        });

        $("#register_form input").keypress(function(e) {
            if (e.which == 13) {
                if ($("#register_form").validate().form()) {
                    $("#register_form").submit();
                }
                return false;
            }
        });
      }

