layui.use(['form', 'layedit', 'laydate'], function() {
				var form = layui.form,
					layer = layui.layer,
					layedit = layui.layedit,
					laydate = layui.laydate;

				form.verify({
					pass: [
						/^[\S]{6,12}$/, '密码必须为6到12位，且不能出现空格'
					]
				});

				//监听提交
				form.on('submit(login)', function(data) {
					layer.alert(JSON.stringify(data.field), {
						title: '最终的提交信息'
					})
					//return false; // 阻止表单跳转
				});
			});