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
			});