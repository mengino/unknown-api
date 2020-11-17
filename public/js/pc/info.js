;(function(){
    //点击导航
    var windowWindth = $(window).width();
    var right = (windowWindth - 1200) / 2
    var objHeight = $(".g-azd-crank").offset().top;
    $(window).scroll(function () {
        var scrollHeight = $(window).scrollTop();
        if (scrollHeight >= 200 + 20) {
            $(".g-info-tab").addClass("m-info-float");
        } else {
            $(".g-info-tab").removeClass("m-info-float");
        }
        if (scrollHeight > objHeight) {
            $(".g-azd-crank").css({
                'width': '286px',
                'position': 'fixed',
                'top': '-12px',
                'right': right + 'px',
                'zIndex': '9999',
                'margin': '12px 0px 0px'
            })
        } else {
            $(".g-azd-crank").css({
                'position': 'static',
                'width':'286px'
            })
        }
    });
})();