
function viewcount(type, id,aid){
    if ( type == "" || id <= 0 ){
        return false;
    }
    var referer = document.referrer;
    var sendData={'chal':type, 'referer':referer,'id':id,'aid':aid};
    $.ajax({
        url:"/views.html",
        type:"get",
        dataType:"json",
        data:sendData,
        success:function(data){}
    })
}