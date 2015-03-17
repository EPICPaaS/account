$(document).ready(function(){
	handleSubmit();
});

    var handleSubmit = function() {
		
        $("#changePassword_form").validate({
            rules : {
				Password : {
                    required : true,
					minlength: 5,
					maxlength:30,
					
                },
				PasswordRe : {
                    required : true,
					minlength: 5,
					maxlength:30,
                    equalTo: "#PasswordForm-Password"
                }
            },
            messages : {
				Password : {
                    required : "密码不能为空",
					minlength: "密码长度不能小于5",
					maxlength:"最大长度为30"
					
                },
				PasswordRe : {
                    required : "重复密码不能为空",
					minlength: "重复密码长度不能小于5",
					maxlength:"最大长度为30",
                    equalTo: "两次输入密码不一致"
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

        $("#changePassword_form input").keypress(function(e) {
            if (e.which == 13) {
                if ($("#changePassword_form").validate().form()) {
                    $("#changePassword_form").submit();
                }
                return false;
            }
        });
	

      }
	
