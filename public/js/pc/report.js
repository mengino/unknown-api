$('#jubao').click(function () {
    $('.wid_wrap').fadeIn();
    var url = window.location.href;
    $("#qs_url").val(url);

})
$('.win_mengban').click(function () {
    $('.wid_wrap').fadeOut()
})

$('#qs_sub').click(function () {
    var qs_url = $('#qs_url').val();
    var qs_cont = $('#qs_cont').val();
    var qs_call = $('#qs_call').val();
    var strRegex = '(https?|ftp|file)://[-A-Za-z0-9+&@#/%?=~_|!:,.;]+[-A-Za-z0-9+&@#/%=~_|]';

    var re = new RegExp(strRegex);
    if (qs_url == '') {
        alert('请输入举报网址')
        return false;
    } else if (!re.test(qs_url)) {
        alert("请输入正确的url地址");
        return false;
    } else if (qs_cont == '') {
        alert('请输入举报内容')
        return false;
    } else if (qs_call == '') {
        alert('请输入联系方式')
        return false;
    }else{
        $.ajax({
            url:"/report.html",
            data:{url:qs_url,content:qs_cont,contact:qs_call},
            type:"post",
            dataType: 'json',
            success:function(data){
                alert(data['msg']);
                $('.wid_wrap').fadeOut()
            }
        });
    }
})