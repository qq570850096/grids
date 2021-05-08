$('#Info').on('click', function() {
	layer.open({
		type: 2,
		area: ['1000px', '600px'],
		title: '个人中心',
		shadeClose: true, //点击弹层外区域关闭
		content: 'userinfo/info.html',
	});
});

$('#RePass').on('click', function() {
	layer.open({
		type: 2,
		area: ['430px', '320px'],
		title: '修改密码',
		shadeClose: true, //点击弹层外区域关闭
		content: 'userinfo/RePass.html',
	});
});

$("dd>a:not(.unTab)").click(function(e) {
	e.preventDefault();
	var id = $(this).attr("id");
	var length = $("li[lay-id=" + id + "]").length;
	
	var height = $(window).height() - 185;
	var url = $(this).attr("href");
	
	var title = $(this).text();
	
	layui.use('element', function() {
		var element = layui.element;
		var layer = layui.layer;
		var content = '<iframe style="width: 100%; height: ' + height + 'px" src="' + url +'" frameborder="0" scrolling="no"></iframe>'
		if (length == 0) {
			element.tabAdd('tab', {
				title: title,
				content: content,
				id: id
			});
		}
		element.tabChange("tab", id);
	});
});
